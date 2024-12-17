package service

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) GetId() int64 {
	return 123
}
