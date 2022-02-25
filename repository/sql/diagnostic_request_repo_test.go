package sql_test

import (
	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosticRequestRepo", func() {

	var (
		diagnosticRequestRepo service.DiagnosticRequestRepo
		diagnosticRequestOne  model.DiagnosticRequest
		diagnosticRequestTwo  model.DiagnosticRequest
		diagnosticRequestID   uint
		userID                uint
		vetOrgID              uint
		customerID            uint
	)

	BeforeEach(func() {
		diagnosticRequestRepoImpl := sql.DiagnosticRequestRepo{Database: database}
		diagnosticRequestRepo = &diagnosticRequestRepoImpl

		diagnosticRequestID = 12345
		userID = 12345
		vetOrgID = 12345
		customerID = 12345

		diagnosticRequestOne = model.DiagnosticRequest{}
		diagnosticRequestTwo = model.DiagnosticRequest{}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.DiagnosticRequest{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

})
