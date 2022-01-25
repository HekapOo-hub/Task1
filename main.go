package main

import (
	"Task1/internal/handlers"
	"Task1/internal/repository"
	"Task1/internal/service"
	"context"
	"fmt"
	echo2 "github.com/labstack/echo/v4"
)

func main() {
	repo, err := repository.NewRepository(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	echo := &handlers.EchoHandler{echo2.New(), service.Service{repo}}
	echo.Register()
	err = echo.Start(":1323")
	if err != nil {
		fmt.Println(err)
	}
}
