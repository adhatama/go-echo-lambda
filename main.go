package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	server := Server{}
	server.Mount(e)

	isLambda := os.Getenv("LAMBDA")

	if isLambda == "TRUE" {
		lambdaAdapter := &LambdaAdapter{Echo: e}
		lambda.Start(lambdaAdapter.Handler)
	} else {
		e.Logger.Fatal(e.Start(":1234"))
	}
}

type Server struct {
}

func (s *Server) Mount(e *echo.Echo) {
	e.GET("/hello", s.GetHello)
	e.POST("/hello", s.PostHello)
	e.GET("/greeting", s.GetGreeting)
	e.POST("/address", s.PostAddress)
}

func (s *Server) GetHello(c echo.Context) error {
	log.Print(c.Request())

	return c.JSON(http.StatusOK, "Hello World")
}

func (s *Server) PostHello(c echo.Context) error {
	log.Print("MASUK HELLO NIIIIH")
	log.Print(c.Request())
	log.Print(c.ParamValues())
	log.Print(c.FormParams())

	name := c.FormValue("name")
	hello := "Hello, " + name + "!"

	return c.JSON(http.StatusOK, hello)
}

func (s *Server) GetGreeting(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to test")
}

func (s *Server) PostAddress(c echo.Context) error {
	address := c.FormValue("address")
	txt := "Your address is " + address

	return c.JSON(http.StatusOK, txt)
}
