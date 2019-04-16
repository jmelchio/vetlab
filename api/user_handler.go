package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jmelchio/vetlab/model"
	"github.com/tedsuo/rata"
)

// UserServer struct allows the UserService injection into the REST handler
type UserServer struct {
	UserService UserService
}

// Constants for error messages and requests
const (
	CreateUser         = "create_user"
	UpdateUser         = "update_user"
	DeleteUser         = "delete_user"
	Login              = "login"
	FindUser           = "find_user"
	FindUserByUserName = "find_user_by_user_name"

	UnableToCreateUser = "Unable to create a user"
	UnableToUpdateUser = "Unable to update a user"
	UnableToDeleteUser = "Unable to delete a user"
	UnableToLoginUser  = "Unable to login the user"
	UnableToFindUser   = "Unable to find the user(s)"
)

// UserRoutes are the REST endpoint routes for the user REST interface
var UserRoutes = rata.Routes{
	{Path: "/user", Method: rata.POST, Name: CreateUser},
	{Path: "/user", Method: rata.PUT, Name: UpdateUser},
	{Path: "/user", Method: rata.DELETE, Name: DeleteUser},
	{Path: "/user/login", Method: rata.POST, Name: Login},
	{Path: "/user/:user_id", Method: rata.GET, Name: FindUser},
	{Path: "/user/username/:user_name", Method: rata.GET, Name: FindUserByUserName},
}

// NewUserHandler provides the factory function to create the REST interface for user actions
func NewUserHandler(userService UserService) (http.Handler, error) {
	userServer := &UserServer{UserService: userService}

	handlers := rata.Handlers{
		CreateUser:         openCors(http.HandlerFunc(userServer.CreateUser), "*"),
		UpdateUser:         openCors(http.HandlerFunc(userServer.UpdateUser), "*"),
		DeleteUser:         openCors(http.HandlerFunc(userServer.DeleteUser), "*"),
		Login:              openCors(http.HandlerFunc(userServer.Login), "*"),
		FindUser:           openCors(http.HandlerFunc(userServer.FindUser), "*"),
		FindUserByUserName: openCors(http.HandlerFunc(userServer.FindUserByUserName), "*"),
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

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(newUser); err != nil {
		log.Printf("Problem encoding new user: %s", err.Error())
	}
}

// UpdateUser takes care of the api request to update a User in the system
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

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(newUser); err != nil {
		log.Printf("Problem encoding new user: %s", err.Error())
	}
}

// DeleteUser takes care of hanlding the api request for deleting a User
func (userServer *UserServer) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var deleteUser model.User
	err = json.Unmarshal(requestBody, &deleteUser)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	err = userServer.UserService.DeleteUser(context.TODO(), deleteUser)
	if err != nil {
		http.Error(writer, UnableToDeleteUser, http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

// Login handles the api request to login to the system
func (userServer *UserServer) Login(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var loginRequest model.LoginRequest
	err = json.Unmarshal(requestBody, &loginRequest)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	loginUser, err := userServer.UserService.Login(context.TODO(), loginRequest.UserName, loginRequest.Password)
	if err != nil {
		http.Error(writer, UnableToLoginUser, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(loginUser); err != nil {
		log.Printf("Problem encoding login user: %s", err.Error())
	}
}

// FindUser handles the api request to find a user by their user id
func (userServer *UserServer) FindUser(writer http.ResponseWriter, request *http.Request) {
	userID := rata.Param(request, "user_id")

	if uintValue, err := strconv.ParseUint(userID, 10, 32); err == nil {
		foundUser, err := userServer.UserService.FindUserByID(context.TODO(), uint(uintValue))
		if err != nil {
			http.Error(writer, UnableToFindUser, http.StatusNotFound)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(foundUser); err != nil {
			log.Printf("Problem encoding found user: %s", err.Error())
		}
		return
	}
	http.Error(writer, NoParamsFound, http.StatusBadRequest)
	return
}

// FindUserByUserName handles the request to find a user by their user name
func (userServer *UserServer) FindUserByUserName(writer http.ResponseWriter, request *http.Request) {
	userName := rata.Param(request, "user_name")

	foundUser, err := userServer.UserService.FindUserByUserName(context.TODO(), userName)
	if err != nil {
		http.Error(writer, UnableToFindUser, http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(foundUser); err != nil {
		log.Printf("Problem encoding found user: %s", err.Error())
	}
	return
}
