package api

import (
	"net/http"

	"github.com/tedsuo/rata"
)

// DiagnosticRequestServer struct allows the DiagnosticReportService injection into the REST handler
type DiagnosticRequestServer struct {
	DiagnosticRequestService DiagnosticRequestService
}

const (
	SubmitDiagnosticRequest = "submit_diagnostic_request"
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
}
