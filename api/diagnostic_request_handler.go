package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jmelchio/vetlab/model"
	"github.com/tedsuo/rata"
)

// DiagnosticRequestServer struct allows the DiagnosticReportService injection into the REST handler
type DiagnosticRequestServer struct {
	DiagnosticRequestService DiagnosticRequestService
}

const (
	SubmitDiagnosticRequest         = "submit_diagnostic_request"
	UnableToSubmitDiagnosticRequest = "Unable to submit diagnostic request"
)

// DiagnosticRequestRoutes are the REST endpoint routes for the diagnostic requests REST interface
var DiagnosticRequestRoutes = rata.Routes{
	{Path: "/diagnosticrequest", Method: rata.POST, Name: SubmitDiagnosticRequest},
}

// NewDiagnosticRequestHandler provides the factory function to create the REST interface for report requests
func NewDiagnosticRequestHandler(diagnosticRequestService DiagnosticRequestService) (http.Handler, error) {
	diagnosticRequestServer := &DiagnosticRequestServer{DiagnosticRequestService: diagnosticRequestService}

	handlers := rata.Handlers{
		SubmitDiagnosticRequest: http.HandlerFunc(diagnosticRequestServer.SubmitDiagnosticRequest),
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
