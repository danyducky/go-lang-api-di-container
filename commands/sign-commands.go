package sign

// Represent the fields that allows the user to log in.
type LoginCommand struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
