package sql_test

import (
	"flag"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jmelchio/vetlab/model"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var (
	database *gorm.DB
	err      error
	dbHost   string
)

// init makes it possible to set the database host name on the command line
// i.e. 'ginkgo repository/... -- -dbHost=<hostname>
func init() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "Database hostname")
}

var _ = BeforeSuite(func() {
	dsn := fmt.Sprintf("host=%s port=5432 user=postgres password=password dbname=vetlab sslmode=disable", dbHost)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	err = database.Migrator().AutoMigrate(&model.DiagnosticReport{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.DiagnosticRequest{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.Customer{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.VetOrg{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.User{})
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	_ = database.Migrator().DropTable(&model.DiagnosticReport{})
	_ = database.Migrator().DropTable(&model.DiagnosticRequest{})
	_ = database.Migrator().DropTable(&model.Customer{})
	_ = database.Migrator().DropTable(&model.VetOrg{})
	_ = database.Migrator().DropTable(&model.User{})

})
