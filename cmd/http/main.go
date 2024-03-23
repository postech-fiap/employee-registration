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

	emailRepository := repository.NewEmailRepository(configuration)
	mirrorRepository := repository.NewMirrorRepository(conn)
	findRegisterDayByUserIdRepository := repository.NewFindRegisterDayByUserIdRepository(conn)
	registerRepository := repository.NewRegisterRepository(conn)

	//amqp
	AMQPChannel, err := amqp.OpenConnection(configuration)
	if err != nil {
		panic(err)
	}
	defer amqp.CloneConnection()

	// queue publisher
	registerQueuePublisher := publisher.NewRegisterQueuePublisher(AMQPChannel)

	// usecase
	mirrorUseCase := usecase.NewMirrorUseCase(mirrorRepository)
	reportUseCase := usecase.NewReportUseCase(mirrorUseCase, emailRepository)
	registerUseCase := usecase.NewRegisterUseCase(registerRepository, registerQueuePublisher)
	findRegisterDayByUserIdUseCase := usecase.FindAllRegisterDayByUserIdUseCase(findRegisterDayByUserIdRepository)

	// service
	pingService := http.NewPingService()
	reportHandler := http.NewReportHandler(reportUseCase)
	registerService := http.NewRegisterService(registerUseCase)
	dailyRegistryHandler := http.NewFindAllDailyRegisterHandler(findRegisterDayByUserIdUseCase)

	// queue consumer
	orderQueueConsumer := consumer.NewRegisterQueueConsumer(AMQPChannel, registerUseCase)
	orderQueueConsumer.Listen()

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)
	router.GET("/report", reportHandler.Handle)
	router.GET("/register", dailyRegistryHandler.Handle)
	router.POST("/register", registerService.Register)

	address := fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)
	router.Run(address)
}
