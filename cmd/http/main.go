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
	findRegisterDayByUserIdRepository := repository.NewFindRegisterDayByUserIdRepository(conn)

	// amqp
	//AMQPChannel, err := amqp.OpenConnection(configuration)
	//if err != nil {
	//	panic(err)
	//}
	//defer amqp.CloneConnection()

	// queue publisher
	// ...

	// usecase
	findRegisterDayByUserIdUseCase := usecase.FindAllRegisterDayByUserIdUseCase(findRegisterDayByUserIdRepository)

	// service
	pingService := http.NewPingService()
	registerDayHandler := http.NewFindRegisterDayByUserIdHandler(findRegisterDayByUserIdUseCase)

	// queue consumer
	// ...

	router := gin.New()
	router.Use(middlewares.ErrorService)
	router.GET("/ping", pingService.Ping)
	router.GET("/user/:id/register", registerDayHandler.Handle)

	address := fmt.Sprintf("%s:%s", configuration.Server.Host, configuration.Server.Port)
	router.Run(address)
}
