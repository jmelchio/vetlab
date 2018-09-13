package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jmelchio/vetlab/model"
	"github.com/tedsuo/rata"
)

// UserServer struct allows the UserService injection into the REST handler
type UserServer struct {
	UserService UserService
}

const (
	CreateUser = "create_user"
	UpdateUser = "update_user"
	DeleteUser = "delete_user"
	Login      = "login"
	FindUser   = "find_user"

	EmptyBody          = "Body of the request is empty"
	InvalidBody        = "Body of the request is invalid"
	UnableToCreateUser = "Unable to create a user"
	UnableToUpdateUser = "Unable to update a user"
	UnableToDeleteUser = "Unable to delete a user"
	UnableToParseBody  = "Unable to parse request body"
	UnableToLoginUser  = "Unable to login the user"
	UnableToFindUser   = "Unable to find the user(s)"
	NoParamsFound      = "No parameters found on request"
)

// UserRoutes are the REST endpoint routes for the user REST interface
var UserRoutes = rata.Routes{
	{Path: "/user/create", Method: rata.POST, Name: CreateUser},
	{Path: "/user/update", Method: rata.PUT, Name: UpdateUser},
	{Path: "/user/delete", Method: rata.DELETE, Name: DeleteUser},
	{Path: "/user/login", Method: rata.POST, Name: Login},
	{Path: "/user/find", Method: rata.GET, Name: FindUser},
}

// NewUserHandler provides the factory function to create the REST interface for user actions
func NewUserHandler(userService UserService) (http.Handler, error) {
	userServer := &UserServer{UserService: userService}

	handlers := rata.Handlers{
		CreateUser: http.HandlerFunc(userServer.CreateUser),
		UpdateUser: http.HandlerFunc(userServer.UpdateUser),
		DeleteUser: http.HandlerFunc(userServer.DeleteUser),
		Login:      http.HandlerFunc(userServer.Login),
		FindUser:   http.HandlerFunc(userServer.FindUser),
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

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(loginUser)
}

func (userServer *UserServer) FindUser(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(writer, NoParamsFound, http.StatusBadRequest)
		return
	}

	if len(request.Form) != 1 {
		http.Error(writer, NoParamsFound, http.StatusBadRequest)
		return
	}

	possibleParams := []string{"user_name", "user_id", "vet_org_id"}
	for _, param := range possibleParams {
		valueFound := request.Form.Get(param)
		if len(valueFound) > 0 {
			switch param {
			case "user_name":
				foundUser, err := userServer.UserService.FindUserByUserName(context.TODO(), valueFound)
				if err != nil {
					http.Error(writer, UnableToFindUser, http.StatusNotFound)
					return
				}
				writer.WriteHeader(http.StatusOK)
				json.NewEncoder(writer).Encode(foundUser)
				return
			case "user_id":
				if uintValue, err := strconv.ParseUint(valueFound, 10, 32); err == nil {
					// Extra conversion to uint seems needed to bug in strconv.ParseUint
					foundUser, err := userServer.UserService.FindUserByID(context.TODO(), uint(uintValue))
					if err != nil {
						http.Error(writer, UnableToFindUser, http.StatusNotFound)
						return
					}
					writer.WriteHeader(http.StatusOK)
					json.NewEncoder(writer).Encode(foundUser)
					return
				}
				http.Error(writer, NoParamsFound, http.StatusBadRequest)
				return
			case "vet_org_id":
				if uintValue, err := strconv.ParseUint(valueFound, 10, 32); err == nil {
					// Extra conversion to uint seems needed to bug in strconv.ParseUint
					foundUsers, err := userServer.UserService.FindUsersByVetOrgID(context.TODO(), uint(uintValue))
					if err != nil {
						http.Error(writer, UnableToFindUser, http.StatusNotFound)
						return
					}
					writer.WriteHeader(http.StatusOK)
					json.NewEncoder(writer).Encode(foundUsers)
					return
				}
				http.Error(writer, NoParamsFound, http.StatusBadRequest)
				return
			}
		}
	}
	http.Error(writer, NoParamsFound, http.StatusBadRequest)
	return
}
