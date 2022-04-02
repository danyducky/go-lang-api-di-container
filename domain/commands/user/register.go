package user

// Represent the fields that allows the user to register.
type RegisterCommand struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}
