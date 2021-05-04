package userModel

// ReqUserSingUp de
type ReqUserSingUp struct {
	FirstName *string `json:"first_name" validator:"required"`
	LastName  *string `json:"last_name" validator:"required"`
	Email     *string `json:"email" validator:"required"`
	Password  *string `json:"password" validator:"required"`
}
