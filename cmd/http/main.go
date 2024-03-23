package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/postech-fiap/employee-registration/cmd/amqp"
	"github.com/postech-fiap/employee-registration/cmd/config"
	repositoryAdapter "github.com/postech-fiap/employee-registration/cmd/repository"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http"
	"github.com/postech-fiap/employee-registration/internal/adapter/handler/http/middlewares"
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/consumer"
	"github.com/postech-fiap/employee-registration/internal/adapter/queue/publisher"
	"github.com/postech-fiap/employee-registration/internal/adapter/repository"
	"github.com/postech-fiap/employee-registration/internal/core/usecase"
)

func main() {
	// config
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// repository
	conn, err := repositoryAdapter.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer repositoryAdapter.CloseConnection()
	reportRepository := repository.NewReportRepository(conn)

	//amqp
	AMQPChannel, err := amqp.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer amqp.CloneConnection()

	// queue publisher
	registerQueuePublisher := publisher.NewRegisterQueuePublisher(AMQPChannel)

	// usecase

	registerUseCase := usecase.NewRegisterUseCase(reportRepository, registerQueuePublisher)

	// service
	pingService := http.NewPingService()
	registerService := http.NewRegisterService(registerUseCase)

	// queue consumer
	orderQueueConsumer := consumer.NewRegisterQueueConsumer(AMQPChannel, registerUseCase)
	orderQueueConsumer.Listen()

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)
	router.POST("/register", registerService.Register)

	address := fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)
	router.Run(address)
}
