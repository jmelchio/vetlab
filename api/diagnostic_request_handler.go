package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/tedsuo/rata"

	"github.com/jmelchio/vetlab/model"
)

// DiagnosticRequestServer struct allows the DiagnosticReportService injection into the REST handler
type DiagnosticRequestServer struct {
	DiagnosticRequestService DiagnosticRequestService
	VetOrgService            VetOrgService
	UserService              UserService
	CustomerService          CustomerService
}

// Constants for error messages and request naming
const (
	SubmitDiagnosticRequest                  = "submit_diagnostic_request"
	UnableToSubmitDiagnosticRequest          = "unable to submit diagnostic request"
	DiagnosticRequestByID                    = "diagnostic_request_by_id"
	DiagnosticRequestsByVetOrgID             = "diagnotistic_requests_by_vetorg_id"
	DiagnosticRequestsByUserID               = "diagnotistic_requests_by_user_id"
	DiagnosticRequestsByCustomerID           = "diagnotistic_requests_by_customer_id"
	DiagnosticRequestsByVetOrgIDAndDateRange = "diagnostic_requests_by_vetorg_id_and_date_range"
	UnableToParseParams                      = "unable to parse request parameters(s)"
	ErrorFetchingDiagnosticRequests          = "error occurred attempting to retrieve diagnostic request(s)"
	ErrorFetchingVetOrg                      = "error occurred attempting to retrieve vetOrg"
	ErrorFetchingUser                        = "error occurred attempting to retrieve user"
	ErrorFetchingCustomer                    = "error occurred attempting to retrieve customer"
)

// DiagnosticRequestRoutes are the REST endpoint routes for the diagnostic requests REST interface
var DiagnosticRequestRoutes = rata.Routes{
	{Path: "/diagnosticrequests", Method: rata.POST, Name: SubmitDiagnosticRequest},
	{Path: "/diagnosticrequests/:request_id", Method: rata.GET, Name: DiagnosticRequestByID},
	{Path: "/diagnosticrequests/vetorg/:vetorg_id", Method: rata.GET, Name: DiagnosticRequestsByVetOrgID},
	{Path: "/diagnosticrequests/vetorg/:vetorg_id/start/:start_date/end/:end_date", Method: rata.GET, Name: DiagnosticRequestsByVetOrgIDAndDateRange},
	{Path: "/diagnosticrequests/user/:user_id", Method: rata.GET, Name: DiagnosticRequestsByUserID},
	{Path: "/diagnosticrequests/customer/:customer_id", Method: rata.GET, Name: DiagnosticRequestsByCustomerID},
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
		SubmitDiagnosticRequest:                  openCors(http.HandlerFunc(diagnosticRequestServer.SubmitDiagnosticRequest), "*"),
		DiagnosticRequestByID:                    openCors(http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequest), "*"),
		DiagnosticRequestsByVetOrgID:             openCors(http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByVetOrg), "*"),
		DiagnosticRequestsByUserID:               openCors(http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByUser), "*"),
		DiagnosticRequestsByCustomerID:           openCors(http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByCustomer), "*"),
		DiagnosticRequestsByVetOrgIDAndDateRange: openCors(http.HandlerFunc(diagnosticRequestServer.FindDiagnotisticRequestByDateRange), "*"),
	}

	return rata.NewRouter(DiagnosticRequestRoutes, handlers)
}

// SubmitDiagnosticRequest is the REST function that allows for the creation of a report request
func (diagnosticRequestServer *DiagnosticRequestServer) SubmitDiagnosticRequest(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, EmptyBody, http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	var diagnosticRequest model.DiagnosticRequest
	dec := json.NewDecoder(request.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&diagnosticRequest); err != nil {
		http.Error(writer, InvalidBody, http.StatusBadRequest)
		return
	}

	newDiagnosticRequest, err := diagnosticRequestServer.DiagnosticRequestService.SubmitDiagnosticRequest(request.Context(), diagnosticRequest)
	if err != nil {
		http.Error(writer, UnableToSubmitDiagnosticRequest, http.StatusInternalServerError)
		return
	}

	writeJSONResponse(writer, http.StatusCreated, newDiagnosticRequest)
}

// FindDiagnotisticRequest is a handler that handles searches for diagstic requests by ID
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequest(writer http.ResponseWriter, request *http.Request) {
	requestID, err := strconv.ParseUint(rata.Param(request, "request_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	diagnosticRequest, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByID(request.Context(), uint(requestID))
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writeJSONResponse(writer, http.StatusOK, diagnosticRequest)
}

// FindDiagnotisticRequestByVetOrg is a handler that handles searches for diagnostic requests by VetOrg
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByVetOrg(writer http.ResponseWriter, request *http.Request) {
	vetOrgID, err := strconv.ParseUint(rata.Param(request, "vetorg_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	vetOrg, err := diagnosticRequestServer.VetOrgService.FindVetOrgByID(request.Context(), uint(vetOrgID))
	if err != nil {
		http.Error(writer, ErrorFetchingVetOrg, http.StatusNotFound)
		return
	}

	diagnosticRequestList, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByVetOrg(request.Context(), *vetOrg)
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writeJSONResponse(writer, http.StatusOK, diagnosticRequestList)
}

// FindDiagnotisticRequestByUser is a handler that handles searches for diagnostic requests by User
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByUser(writer http.ResponseWriter, request *http.Request) {
	userID, err := strconv.ParseUint(rata.Param(request, "user_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	user, err := diagnosticRequestServer.UserService.FindUserByID(request.Context(), uint(userID))
	if err != nil {
		http.Error(writer, ErrorFetchingUser, http.StatusNotFound)
		return
	}

	diagnosticRequestList, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByUser(request.Context(), *user)
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writeJSONResponse(writer, http.StatusOK, diagnosticRequestList)
}

// FindDiagnotisticRequestByCustomer is a handler that handles searches for diagnostic requests by Customer
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByCustomer(writer http.ResponseWriter, request *http.Request) {
	customerID, err := strconv.ParseUint(rata.Param(request, "customer_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	customer, err := diagnosticRequestServer.CustomerService.FindCustomerByID(request.Context(), uint(customerID))
	if err != nil {
		http.Error(writer, ErrorFetchingCustomer, http.StatusNotFound)
		return
	}

	diagnosticRequestList, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByCustomer(request.Context(), *customer)
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writeJSONResponse(writer, http.StatusOK, diagnosticRequestList)
}

// FindDiagnotisticRequestByDateRange is a handler that handles searches for diagnostic requests by VetOrg and date range
func (diagnosticRequestServer *DiagnosticRequestServer) FindDiagnotisticRequestByDateRange(writer http.ResponseWriter, request *http.Request) {
	vetOrgID, err := strconv.ParseUint(rata.Param(request, "vetorg_id"), 10, 32)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	startDateAsString := rata.Param(request, "start_date")
	endDateAsString := rata.Param(request, "end_date")

	if len(startDateAsString) == 0 || len(endDateAsString) == 0 {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	const shortForm = "20060102"

	startDate, err := time.Parse(shortForm, startDateAsString)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(shortForm, endDateAsString)
	if err != nil {
		http.Error(writer, UnableToParseParams, http.StatusBadRequest)
		return
	}

	vetOrg, err := diagnosticRequestServer.VetOrgService.FindVetOrgByID(request.Context(), uint(vetOrgID))
	if err != nil {
		http.Error(writer, ErrorFetchingVetOrg, http.StatusNotFound)
		return
	}

	diagnosticRequestList, err := diagnosticRequestServer.DiagnosticRequestService.FindRequestByDateRange(request.Context(), startDate, endDate, *vetOrg)
	if err != nil {
		http.Error(writer, ErrorFetchingDiagnosticRequests, http.StatusNotFound)
		return
	}

	writeJSONResponse(writer, http.StatusOK, diagnosticRequestList)
}
