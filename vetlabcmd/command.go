package vetlabcmd

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/repository"
	"github.com/jmelchio/vetlab/service"
)

var (
	err         error
	database    *gorm.DB
	userHandler http.Handler
)

func Run() {
	database, err = gorm.Open("postgres", "host=localhost user=postgres sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %s", err.Error()))
	}

	userRepo := repository.SQLUserRepo{Database: database}
	userService := service.User{UserRepo: userRepo}
	userHandler, err = api.NewUserHandler(userService)
	if err != nil {
		panic(fmt.Sprintf("Unable to create the user handler: %s", err.Error()))
	}
	http.Handle("/", userHandler)

	http.ListenAndServe(":8080", nil)
}
