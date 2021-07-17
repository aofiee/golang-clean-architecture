package constants

const (
	ERR_DNS_CONNECTION_EMPTY = "dns connection is empty"

	ERR_PARSE_JSON_FAIL          = "cannot parse json"
	ERR_INPUT_ERROR              = "input error"
	ERR_CREATE_ROLE_SUCCESSFUL   = "create role successful"
	ERR_CANNOT_CREATE_ROLE       = "cannot create role"
	ERR_CANNOT_PARSE_PARAMS      = "cannot parse params"
	ERR_CANNOT_GET_ALL_ROLES     = "cannot get all roles"
	ERR_GET_ALL_ROLE_SUCCESSFULE = "get all role successful"
	ERR_CANNOT_DELETE_ROLE       = "cannot delete roles"
	ERR_CANNOT_RESTORE_ROLES     = "cannot restore roles"
	ERR_CANNOT_GET_ROLE_ID       = "cannot get role id"

	ERR_CREATE_USER_SUCCESSFUL      = "create user successful"
	ERR_LOGIN_SUCCESSFUL            = "login successful"
	ERR_USERNAME_PASSWORD_INCORRECT = "username and password is incorrect"

	ERR_CANNOT_GET_ROLE_NAME         = "cannot get role name from uuid"
	ERR_TOKEN_CANNOT_SIGNED_KEY      = "token cannot signed with a key"
	ERR_CANNOT_SAVE_TOKEN_TO_REDIS   = "token cannot save in queue service"
	ERR_CANNOT_DELETE_TOKEN_TO_REDIS = "token cannot delete from queue service"
	ERR_LOGOUT_COMPLETED             = "logout successful"
	ERR_TOKEN_SIGNED_NOT_MATCH       = "token signing error"
	ERR_REFRESH_TOKEN_EXPIRE         = "Refresh token expired"
	ERR_GET_USER_BY_UUID_NOT_FOUND   = "user not found"
	ERR_REFRESH_TOKEN_SUCCESSFUL     = "refresh token successful"
)