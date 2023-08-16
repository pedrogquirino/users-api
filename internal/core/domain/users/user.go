package users

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
}
type UserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func ConvertUserRequestToUser(userRequest UserRequest) *User {
	return &User{
		Name:     userRequest.Name,
		Age:      userRequest.Age,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}
