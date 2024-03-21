package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/cmd/config"
	"github.com/postech-fiap/employee-registration/cmd/repository"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http/middlewares"
)

func main() {
	// config
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// repository
	_, err = repository.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer repository.CloseConnection()

	// amqp
	//AMQPChannel, err := amqp.OpenConnection(configuration)
	//if err != nil {
	//	panic(err)
	//}
	//defer amqp.CloneConnection()

	// queue publisher
	// ...

	// usecase
	// ...

	// service
	pingService := http.NewPingService()

	// queue consumer
	// ...

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)

	address := fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)
	router.Run(address)
}
