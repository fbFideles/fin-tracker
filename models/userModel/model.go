package userModel

// ReqUserSingUp de
type ReqUserSingUp struct {
	FirstName *string `json:"first_name" validator:"required"`
	LastName  *string
	Email     *string
	Password  *string
}
