package user

type CreateUserForm struct {
	Username   string `json:"username" validate:"required,min=5,max=32"`
	Email      string `json:"email" validate:"required,email"`
	Credential string `json:"credential" validate:"required,min=9"`
}

type BasicAuthForm struct {
	Username   string `json:"username" validate:"required_without=Email,omitempty,min=5,max=32"`
	Email      string `json:"email" validate:"required_without=Username,omitempty,email"`
	Credential string `json:"credential" validate:"required,min=9"`
}
