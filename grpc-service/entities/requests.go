package entities

type Status struct {
	Code    int
	Message string
}

type CreateUserRequest struct {
	User User
}

type CreateUserResponse struct {
	Status
}

type GetUserRequest struct {
	UserID string
}
type GetUserResponse struct {
	User
	Status
}
