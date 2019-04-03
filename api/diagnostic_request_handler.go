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

// DiagnosticRequestServer struct allows the DiagnosticReportService injection into the REST handler
type DiagnosticRequestServer struct {
	DiagnosticRequestService DiagnosticRequestService
	VetOrgService            VetOrgService
	UserService              UserService
	CustomerService          CustomerService
}

const (
	SubmitDiagnosticRequest                  = "submit_diagnostic_request"
	UnableToSubmitDiagnosticRequest          = "Unable to submit diagnostic request"
	DiagnosticRequestByID                    = "diagnostic_request_by_id"
	DiagnosticRequestsByVetOrgID             = "diagnotistic_requests_by_vetorg_id"
	DiagnosticRequestsByUserID               = "diagnotistic_requests_by_user_id"
	DiagnosticRequestsByCustomerID           = "diagnotistic_requests_by_customer_id"
	DiagnosticRequestsByVetOrgIDAndDateRange = "diagnostic_requests_by_vetorg_id_and_date_range"
	UnableToParseParams                      = "Unable to parse request parameters(s)"
	ErrorFetchingDiagnosticRequests          = "Error occurred attempting to retrieve diagnostic request(s)"
	ErrorFetchingVetOrg                      = "Error occurred attempting to retrieve vetOrg"
	ErrorFetchingUser                        = "Error occurred attempting to retrieve user"
)

// DiagnosticRequestRoutes are the REST endpoint routes for the diagnostic requests REST interface
var DiagnosticRequestRoutes = rata.Routes{
	{Path: "/diagnosticrequest", Method: rata.POST, Name: SubmitDiagnosticRequest},
	{Path: "/diagnosticrequest/:request_id", Method: rata.GET, Name: DiagnosticRequestByID},
	{Path: "/diagnosticrequest/vetorg/:vetorg_id", Method: rata.GET, Name: DiagnosticRequestsByVetOrgID},
	{Path: "/diagnosticrequest/vetorg/:vetorg_id/start/:start_date/end/:end_date", Method: rata.GET, Name: DiagnosticRequestsByVetOrgIDAndDateRange},
	{Path: "/diagnosticrequest/user/:user_id", Method: rata.GET, Name: DiagnosticRequestsByUserID},
	{Path: "/diagnosticrequest/customer/:customer_id", Method: rata.GET, Name: DiagnosticRequestsByCustomerID},
}

// NewDiagnosticRequestHandler provides the factory function to create the REST interface for report requests
func NewDiagnosticRequestHandler(
	diagnosticRequestService DiagnosticRequestService,
	vetOrgService VetOrgService,
	userService UserService,
	customerService CustomerService,
) (http.Handler, error) {
	diagnosticRequestServer := &DiagnosticRequestServer{
		DiagnosticRequestService: diagnosticRequestService,
		VetOrgService:            vetOrgService,
		UserService:              userService,
		CustomerService:          customerService,
	}

	handlers := rata.Handlers{
		SubmitDiagnosticRequest:                  http.HandlerFunc(diagnosticRequestServer.SubmitDiagnosticRequest),
		DiagnosticRequestByID:                    http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequest),
		DiagnosticRequestsByVetOrgID:             http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByVetOrg),
		DiagnosticRequestsByUserID:               http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByUser),
		DiagnosticRequestsByCustomerID:           http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByCustomer),
		DiagnosticRequestsByVetOrgIDAndDateRange: http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByDateRange),
	}

	return rata.NewRouter(DiagnosticRequestRoutes, handlers)
}

// SubmitDiagnosticRequest is the REST function that allows for the creation of a report request
func (diagnosticRequestServer *DiagnosticRequestServer) SubmitDiagnosticRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var diagnosticRequest model.DiagnosticRequest
	err = json.Unmarshal(requestBody, &diagnosticRequest)
	if err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}

	newDiagnosticRequest, err := diagnosticRequestServer.DiagnosticRequestService.SubmitDiagnosticRequest(context.TODO(), diagnosticRequest)
	if err != nil {
		http.Error(writer, UnableToSubmitDiagnosticRequest, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(newDiagnosticRequest); err != nil {
		log.Printf("Problem encoding new diagnostic request: %s", err.Error())
	}
}

// FindDiagnotisticRequest is a handler that handles searches for diagstic requests by ID
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequest(writer http.ResponseWriter, request *http.Request) {
	requestID, err := strconv.ParseUint(rata.Param(request, "request_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	diagnosticRequest, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByID(context.TODO(), uint(requestID))
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(diagnosticRequest); err != nil {
		log.Printf("Problem encoding returned diagnostic request(s): %s", err.Error())
		return
	}
}

// FindDiagnotisticRequestByVetOrg is a handler that handles searches for diagnostic requests by VetOrg
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByVetOrg(writer http.ResponseWriter, request *http.Request) {
	vetOrgID, err := strconv.ParseUint(rata.Param(request, "vetorg_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	vetOrg, err := diagnosticRequestServer.VetOrgService.FindVetOrgByID(context.TODO(), uint(vetOrgID))
	if err != nil {
		http.Error(writer, ErrorFetchingVetOrg, http.StatusNotFound)
		return
	}

	diagnosticRequestList, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByVetOrg(context.TODO(), *vetOrg)
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(diagnosticRequestList); err != nil {
		log.Printf("Problem encoding returned diagnostic request(s): %s", err.Error())
		return
	}
}

// FindDiagnotisticRequestByUser is a handler that handles searches for diagnostic requests by User
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByUser(writer http.ResponseWriter, request *http.Request) {
	userID, err := strconv.ParseUint(rata.Param(request, "user_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	user, err := diagnosticRequestServer.UserService.FindUserByID(context.TODO(), uint(userID))
	if err != nil {
		http.Error(writer, ErrorFetchingUser, http.StatusNotFound)
		return
	}

	diagnosticRequestList, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByUser(context.TODO(), *user)
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(diagnosticRequestList); err != nil {
		log.Printf("Problem encoding returned diagnostic request(s): %s", err.Error())
		return
	}
}

// FindDiagnotisticRequestByCustomer is a handler that handles searches for diagnostic requests by Customer
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByCustomer(writer http.ResponseWriter, request *http.Request) {
}

// FindDiagnotisticRequestByDateRange is a handler that handles searches for diagnostic requests by VetOrg and date range
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByDateRange(writer http.ResponseWriter, request *http.Request) {
}
