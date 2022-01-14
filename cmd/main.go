package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"time"
	"whattofarm/internal/db/dbservise"
	"whattofarm/internal/handlers/getdata"
	"whattofarm/internal/handlers/update"
)


var (
	USER = "whattofarm"
	PASSWORD = "what"
	HOST = "localhost:27017"
	ADDR = "localhost:8000"
	DATABASE = "whattofarm"
	COLLECTION = "testtask"
)

func init() {
	flag.StringVar(&USER, "u", USER, "MongoDB user name")
	flag.StringVar(&PASSWORD, "p", PASSWORD, "MongoDB user password")
	flag.StringVar(&HOST, "h", HOST, "MongoDB host \"hostname:port\"")
	flag.StringVar(&ADDR, "a", ADDR, "Server host \"hostname:port\"")
	flag.StringVar(&DATABASE, "d", DATABASE, "MongoDB database name")
	flag.StringVar(&COLLECTION, "c", COLLECTION, "MongoDB collection name")
}


func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	logger := e.Logger

	service, err := dbservise.NewService(USER, PASSWORD, HOST, DATABASE,COLLECTION)
	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		if err = service.Disconnect(5*time.Second); err != nil {
			logger.Error(err)
		}
	}()

	if err = service.GetDocumentID(); err != nil {
		logger.Fatal(err)
	}

	e.GET("/counter.gif", update.NewHandler(service))
	e.GET("/", getdata.NewHandler(service))

	logger.Fatal(e.Start(ADDR))
}