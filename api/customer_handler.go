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

// CustomerServer struct allows the CustomerService injection into the REST handler
type CustomerServer struct {
	CustomerService CustomerService
}

const (
	CreateCustomer = "create_customer"
	UpdateCustomer = "update_customer"
	DeleteCustomer = "delete_customer"
	CustomerLogin  = "customer_login"
	FindCustomer   = "find_customer"

	UnableToCreateCustomer = "Unable to create a customer"
	UnableToUpdateCustomer = "Unable to update a customer"
	UnableToDeleteCustomer = "Unable to delete a customer"
	UnableToLoginCustomer  = "Unable to login the customer"
	UnableToFindCustomer   = "Unable to find the customer(s)"
)

// CustomerRoutes are the REST endpoint routes for the customer REST interface
var CustomerRoutes = rata.Routes{
	{Path: "/customer", Method: rata.POST, Name: CreateCustomer},
	{Path: "/customer", Method: rata.PUT, Name: UpdateCustomer},
	{Path: "/customer", Method: rata.DELETE, Name: DeleteCustomer},
	{Path: "/customer/login", Method: rata.POST, Name: CustomerLogin},
	{Path: "/customer", Method: rata.GET, Name: FindCustomer},
}

// NewCustomerHandler provides the factory function to create the REST interface for customer actions
func NewCustomerHandler(customerService CustomerService) (http.Handler, error) {
	customerServer := &CustomerServer{CustomerService: customerService}

	handlers := rata.Handlers{
		CreateCustomer: http.HandlerFunc(customerServer.CreateCustomer),
		UpdateCustomer: http.HandlerFunc(customerServer.UpdateCustomer),
		DeleteCustomer: http.HandlerFunc(customerServer.DeleteCustomer),
		CustomerLogin:  http.HandlerFunc(customerServer.CustomerLogin),
		FindCustomer:   http.HandlerFunc(customerServer.FindCustomer),
	}

	return rata.NewRouter(CustomerRoutes, handlers)
}

// CreateCustomer is the REST endpoint function that allows for the creation of customers in the system
func (customerServer *CustomerServer) CreateCustomer(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, UnableToParseBody, http.StatusBadRequest)
		return
	}

	var createCustomer model.Customer
	err = json.Unmarshal(requestBody, &createCustomer)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	newCustomer, err := customerServer.CustomerService.CreateCustomer(context.TODO(), createCustomer)
	if err != nil {
		http.Error(writer, UnableToCreateCustomer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(newCustomer); err != nil {
		log.Printf("Problem encoding new customer: %s", err.Error())
	}
}

func (customerServer *CustomerServer) UpdateCustomer(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var updateCustomer model.Customer
	err = json.Unmarshal(requestBody, &updateCustomer)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	newCustomer, err := customerServer.CustomerService.UpdateCustomer(context.TODO(), updateCustomer)
	if err != nil {
		http.Error(writer, UnableToUpdateCustomer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(newCustomer); err != nil {
		log.Printf("Problem encoding new customer: %s", err.Error())
	}
}

func (customerServer *CustomerServer) DeleteCustomer(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var deleteCustomer model.Customer
	err = json.Unmarshal(requestBody, &deleteCustomer)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}
	err = customerServer.CustomerService.DeleteCustomer(context.TODO(), deleteCustomer)
	if err != nil {
		http.Error(writer, UnableToDeleteCustomer, http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func (customerServer *CustomerServer) CustomerLogin(writer http.ResponseWriter, request *http.Request) {
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
	loginCustomer, err := customerServer.CustomerService.Login(context.TODO(), loginRequest.UserName, loginRequest.Password)
	if err != nil {
		http.Error(writer, UnableToLoginCustomer, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(loginCustomer); err != nil {
		log.Printf("Problem encoding login customer: %s", err.Error())
	}
}

func (customerServer *CustomerServer) FindCustomer(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(writer, NoParamsFound, http.StatusBadRequest)
		return
	}

	if len(request.Form) != 1 {
		http.Error(writer, NoParamsFound, http.StatusBadRequest)
		return
	}

	possibleParams := []string{"user_name", "customer_id"}
	for _, param := range possibleParams {
		valueFound := request.Form.Get(param)
		if len(valueFound) > 0 {
			switch param {
			case "user_name":
				foundCustomer, err := customerServer.CustomerService.FindCustomerByUserName(context.TODO(), valueFound)
				if err != nil {
					http.Error(writer, UnableToFindCustomer, http.StatusNotFound)
					return
				}
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(writer).Encode(foundCustomer); err != nil {
					log.Printf("Problem encoding found customer: %s", err.Error())
				}
				return
			case "customer_id":
				if uintValue, err := strconv.ParseUint(valueFound, 10, 32); err == nil {
					foundCustomer, err := customerServer.CustomerService.FindCustomerByID(context.TODO(), uint(uintValue))
					if err != nil {
						http.Error(writer, UnableToFindCustomer, http.StatusNotFound)
						return
					}
					writer.Header().Set("Content-Type", "application/json")
					writer.WriteHeader(http.StatusOK)
					if err := json.NewEncoder(writer).Encode(foundCustomer); err != nil {
						log.Printf("Problem encoding found customer: %s", err.Error())
					}
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
