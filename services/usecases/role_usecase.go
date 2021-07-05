package usecases

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
)

type (
	roleUseCase struct {
		roleRepo domains.RoleRepository
	}
)

func NewRoleUseCase(repo domains.RoleRepository) domains.RoleUseCase {
	return &roleUseCase{
		roleRepo: repo,
	}
}
func (r *roleUseCase) CreateRole(role *models.RoleItems) error {
	err := r.roleRepo.CreateRole(role)
	return err
}
func (r *roleUseCase) UpdateRole(role *models.RoleItems, id string) error {
	var chk models.RoleItems
	err := r.roleRepo.GetRole(&chk, id)
	if err != nil {
		return err
	}
	err = r.roleRepo.UpdateRole(role, id)
	return err
}
func (r *roleUseCase) GetRole(role *models.RoleItems, id string) error {
	err := r.roleRepo.GetRole(role, id)
	return err
}
func (r *roleUseCase) GetAllRoles(role *[]models.RoleItems) error {
	err := r.roleRepo.GetAllRoles(role)
	return err
}
