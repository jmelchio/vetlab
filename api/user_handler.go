package api

import (
	"net/http"

	"github.com/concourse/src/github.com/tedsuo/rata"
)

type UserServer struct {
	UserService UserService
}

const (
	CreateUser = "CreateUser"
	UpdateUser = "UpdateUser"
	DeleteUser = "DeleteUser"
	Login      = "Login"
	FindUser   = "FindUser"
)

var UserRoutes = rata.Routes{
	{Path: "/user/create", Method: "POST", Name: CreateUser},
	{Path: "/user/update", Method: "PUT", Name: UpdateUser},
	{Path: "/user/delete", Method: "DELETE", Name: DeleteUser},
	{Path: "/user/login", Method: "POST", Name: Login},
	{Path: "/user/find", Method: "GET", Name: FindUser},
}

func NewUserHandler(userService UserService) (http.Handler, error) {
	userServer := &UserServer{UserService: userService}

	handlers := rata.Handlers{
		CreateUser: http.HandlerFunc(userServer.CreateUser),
		UpdateUser: http.HandlerFunc(userServer.UpdateUser),
		DeleteUser: http.HandlerFunc(userServer.DeleteUser),
		Login:      http.HandlerFunc(userServer.Login),
		FindUser:   http.HandlerFunc(userServer.Login),
	}

	return rata.NewRouter(UserRoutes, handlers)
}

func (userServer *UserServer) CreateUser(writer http.ResponseWriter, request *http.Request) {
}

func (userServer *UserServer) UpdateUser(writer http.ResponseWriter, request *http.Request) {
}

func (userServer *UserServer) DeleteUser(writer http.ResponseWriter, request *http.Request) {
}

func (userServer *UserServer) Login(writer http.ResponseWriter, request *http.Request) {
}

func (userServer *UserServer) FindUser(writer http.ResponseWriter, request *http.Request) {
}
