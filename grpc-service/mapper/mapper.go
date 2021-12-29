package mapper

import "github.com/timoteoBone/final-project-microservice/grpc-service/entities"

func CreateUserRequestToUser(userReq entities.CreateUserRequest) entities.User {

	user := entities.User{
		userReq.User.Id,
		userReq.User.Name,
		userReq.User.Pass,
		userReq.User.Age,
	}
	return user
}
