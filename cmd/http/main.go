package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/cmd/config"
	"github.com/postech-fiap/employee-registration/cmd/repository"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http/middlewares"
	"github.com/postech-fiap/employee-registration/internal/core/usecase"
)

func main() {
	// config
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// repository
	conn, err := repository.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer repository.CloseConnection()
	mirrorRepository := repository.NewMirrorRepository(conn)
	emailRepository := repository.NewEmailRepository(configuration)

	// amqp
	//AMQPChannel, err := amqp.OpenConnection(configuration)
	//if err != nil {
	//	panic(err)
	//}
	//defer amqp.CloneConnection()

	// queue publisher
	// ...

	// usecase
	mirrorUseCase := usecase.NewMirrorUseCase(mirrorRepository)
	reportUseCase := usecase.NewReportUseCase(mirrorUseCase, emailRepository)

	// service
	pingService := http.NewPingService()
	reportHandler := http.NewReportHandler(reportUseCase)

	// queue consumer
	// ...

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)
	router.GET("/report", reportHandler.Handle)

	address := fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)
	router.Run(address)
}
