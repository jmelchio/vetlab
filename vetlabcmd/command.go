package vetlabcmd

import (
	"log"
	"net/http"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
)

var (
	err             error
	database        *gorm.DB
	userHandler     http.Handler
	customerHandler http.Handler
)

func Run() {
	database, err = gorm.Open("postgres", "host=localhost user=postgres dbname=vetlab sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err.Error())
	}

	autoMigrateDB(&model.DiagnosticReport{})
	autoMigrateDB(&model.DiagnosticRequest{})
	autoMigrateDB(&model.Customer{})
	autoMigrateDB(&model.VetOrg{})
	autoMigrateDB(&model.User{})

	userRepo := sql.UserRepo{Database: database}
	userService := service.User{UserRepo: &userRepo}
	userHandler, err = api.NewUserHandler(userService)
	if err != nil {
		log.Fatalf("Unable to create the user handler: %s", err.Error())
	}
	http.Handle("/users", userHandler)
	http.Handle("/users/", userHandler)

	customerRepo := sql.CustomerRepo{Database: database}
	customerService := service.Customer{CustomerRepo: &customerRepo}
	customerHandler, err = api.NewCustomerHandler(customerService)
	if err != nil {
		log.Fatalf("Unable to create the customer handler: %s", err.Error())
	}
	http.Handle("/customers", customerHandler)
	http.Handle("/customers/", customerHandler)

	log.Println("Starting listner on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func autoMigrateDB(table interface{}) {
	if err := database.AutoMigrate(table).Error; err != nil {
		dbCreateFatal(table, err)
	}
	if !database.HasTable(table) {
		dbMissing(reflect.TypeOf(table).String())
	} else {
		dbMigrated(reflect.TypeOf(table).String())
	}
}

func dbCreateFatal(table interface{}, err error) {
	log.Fatalf("Unable to create table: %s, %s", reflect.TypeOf(table).String(), err.Error())
}

func dbMissing(table string) {
	log.Fatalf("DB table missing: %s", table)
}

func dbMigrated(table string) {
	log.Printf("DB table migrated: %s", table)
}
