package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jmelchio/vetlab/model"
	"github.com/tedsuo/rata"
)

// UserServer struct allows the UserService injection into the REST handler
type UserServer struct {
	UserService UserService
}

const (
	CreateUser = "CreateUser"
	UpdateUser = "UpdateUser"
	DeleteUser = "DeleteUser"
	Login      = "Login"
	FindUser   = "FindUser"

	EmptyBody          = "Body of the request is empty"
	InvalidBody        = "Body of the request is invalid"
	UnableToCreateUser = "Unable to create a user"
	UnableToUpdateUser = "Unable to update a user"
	UnableToParseBody  = "Unable to parse request body"
)

// UserRoutes are the REST endpoint routes for the user REST interface
var UserRoutes = rata.Routes{
	{Path: "/user/create", Method: "POST", Name: CreateUser},
	{Path: "/user/update", Method: "PUT", Name: UpdateUser},
	{Path: "/user/delete", Method: "DELETE", Name: DeleteUser},
	{Path: "/user/login", Method: "POST", Name: Login},
	{Path: "/user/find", Method: "GET", Name: FindUser},
}

// NewUserHandler provides the factory function to create the REST interface for user actions
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

// CreateUser is the REST endpoint function that allows for the creation of users in the system
func (userServer *UserServer) CreateUser(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, UnableToParseBody, http.StatusBadRequest)
		return
	}

	var createUser model.User
	err = json.Unmarshal(requestBody, &createUser)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	newUser, err := userServer.UserService.CreateUser(context.TODO(), createUser)
	if err != nil {
		http.Error(writer, UnableToCreateUser, http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(newUser)
}

func (userServer *UserServer) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var updateUser model.User
	err = json.Unmarshal(requestBody, &updateUser)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	newUser, err := userServer.UserService.UpdateUser(context.TODO(), updateUser)
	if err != nil {
		http.Error(writer, UnableToUpdateUser, http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(newUser)
}

func (userServer *UserServer) DeleteUser(writer http.ResponseWriter, request *http.Request) {
}

func (userServer *UserServer) Login(writer http.ResponseWriter, request *http.Request) {
}

func (userServer *UserServer) FindUser(writer http.ResponseWriter, request *http.Request) {
}