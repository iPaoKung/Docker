package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type message struct {
	User    string `json:"user"`
	Success string `json:"msg_success"`
	Fail    string `json:"msg_fail"`
}

// Handler
// ddmmyy.hhmmss|TransactionID|SUCCESS/Fail|Message
func logger(c echo.Context) error {
	var msg message
	var s string
	c.Bind(&msg)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ltime|log.Ldate|log.Lmicroseconds)

	if msg.Fail != "" {
		s = fmt.Sprintf("Fail|%s|%s", msg.User, msg.Fail)
	} else {
		s = fmt.Sprintf("Success|%s|%s", msg.User, msg.Success)
	}

	logger.Print(s)
	fmt.Print(&buf)

	return c.String(http.StatusOK, s)
}
func paolog(c echo.Context) error {
	var msg message
	var s string
	c.Bind(&msg)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ltime|log.Ldate|log.Lmicroseconds)

	if msg.Fail != "" {
		s = fmt.Sprintf("Pao-Fail|%s|%s", msg.User, msg.Fail)
	} else {
		s = fmt.Sprintf("Pao-Success|%s|%s", msg.User, msg.Success)
	}

	logger.Print(s)
	fmt.Print(&buf)

	return c.String(http.StatusOK, s)
}
func helloworld(c echo.Context) error {
	var msg message
	var s string
	c.Bind(&msg)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ltime|log.Ldate|log.Lmicroseconds)

	if msg.Fail != "" {
		s = fmt.Sprintf("Hello-Fail|%s|%s", msg.User, msg.Fail)
	} else {
		s = fmt.Sprintf("Hello-Success|%s|%s", msg.User, msg.Success)
	}

	logger.Print(s)
	fmt.Print(&buf)

	return c.String(http.StatusOK, s)
}
func main() {
	// Echo instance
	e := echo.New()

	// Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", logger)
	e.POST("/pao", paolog)
	e.POST("/helloworld", helloworld)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
