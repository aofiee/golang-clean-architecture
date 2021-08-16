package deliveries

import (
	"bytes"
	"context"
	"errors"
	"text/template"
	"time"

	"github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/mailgun/mailgun-go/v4"
)

type (
	forgotPasswordHandler struct {
		forgotPasswordUseCase domains.ForgorPasswordUseCase
	}
	paramEmail struct {
		Email string `json:"email" validate:"required,email,min=6,max=255"`
	}
	templateResetPassword struct {
		AppName  string
		LinkHash string
		SiteURL  string
		AppPort  string
	}
	paramsResetPassword struct {
		Password   string `json:"password" form:"password" validate:"required,min=6,max=64"`
		RePassword string `json:"re_password" form:"re_password" validate:"required,min=6,max=64"`
	}
)

func NewForgotPasswordHandler(usecase domains.ForgorPasswordUseCase, u *[]models.ModuleMethodSlug) *forgotPasswordHandler {
	moduleRepo := repositories.NewModuleRepository(databases.DB)
	moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
	for _, value := range *u {
		newModule := models.Modules{
			Name:        value.Name,
			Description: value.Description,
			ModuleSlug:  value.Slug,
			Method:      value.Method,
		}
		err := moduleUseCase.GetModuleBySlug(&newModule, value.Method, value.Slug)
		if err != nil {
			moduleUseCase.CreateModule(&newModule)
		}
	}
	return &forgotPasswordHandler{
		forgotPasswordUseCase: usecase,
	}
}
func (f *forgotPasswordHandler) ResetPassword(c *fiber.Ctx) error {
	var param paramEmail
	err := c.BodyParser(&param)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	linkHash, err := f.forgotPasswordUseCase.CreateForgotPasswordHash(param.Email)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_HASHLINK_IN_REDIS, fiber.StatusBadRequest)
	}
	mg := mailgun.NewMailgun(config.ENV.MyDomain, config.ENV.MailGunApiKey)
	sender := config.ENV.EmailAdministrator
	subject := config.ENV.AppName + " reset password"
	recipient := param.Email
	message := mg.NewMessage(sender, subject, "", recipient)
	t, _ := template.ParseFiles("../views/change_password/mail_change_password.html")
	tplData := templateResetPassword{
		AppName:  config.ENV.AppName,
		SiteURL:  config.ENV.SiteURL,
		LinkHash: linkHash,
		AppPort:  config.ENV.AppPort,
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, tplData); err != nil {
		return helpers.FailOnError(c, err, constants.ERR_TEMPLATE_PARSE_DATA, fiber.StatusNotAcceptable)
	}
	body := tpl.String()

	message.SetHtml(body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_MAIL_GUN_FORBIDDEN, fiber.StatusUnauthorized)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_SEND_LINK_RESET_PASSWORD_TO_EMAIL_SUCCESSFUL,
		"error": nil,
		"id":    id,
		"data":  resp,
	})
}

func (f *forgotPasswordHandler) ResetPasswordForm(c *fiber.Ctx) error {
	ok := f.forgotPasswordUseCase.CheckForgotPasswordHashIsExpire(c.Params("id"))
	if !ok {
		return c.Render("../views/change_password/result.html", fiber.Map{
			"error": errors.New("Error 404 - Page not found"),
			"msg":   nil,
		})
	}
	tplData := templateResetPassword{
		AppName:  config.ENV.AppName,
		SiteURL:  config.ENV.SiteURL,
		LinkHash: c.Params("id"),
		AppPort:  config.ENV.AppPort,
	}
	return c.Render("../views/change_password/reset_password.html", fiber.Map{
		"tplData": tplData,
	})
}
func (f *forgotPasswordHandler) ResetPasswordFormExec(c *fiber.Ctx) error {
	var params paramsResetPassword
	err := c.BodyParser(&params)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&params)
	if errorResponse != nil {
		return c.Render("../views/change_password/result.html", fiber.Map{
			"error": errors.New(constants.ERR_PASSWORD_VALIDATE),
			"msg":   nil,
		})
	}
	err = f.forgotPasswordUseCase.ResetPassword(c.Params("id"), params.Password, params.RePassword)
	if err != nil {
		return c.Render("../views/change_password/result.html", fiber.Map{
			"error": err,
			"msg":   nil,
		})
	}
	return c.Render("../views/change_password/result.html", fiber.Map{
		"error": nil,
		"msg":   constants.ERR_RESET_PASSWORD_COMPLETED,
	})
}
