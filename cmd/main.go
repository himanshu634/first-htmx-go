package main

import (
	"first_project/handler"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler()

	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":30000")
	fmt.Println("Working Projects")
}
