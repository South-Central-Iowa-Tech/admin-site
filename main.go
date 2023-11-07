package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "redtowernews/admin-site/controllers"
)

func main() {
  e := echo.New()

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.GET("/", controllers.Hello)

  e.Logger.Fatal(e.Start(":3500"))
}
