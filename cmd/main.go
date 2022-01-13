package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"time"
	"whattofarm/internal/db/dbclient"
	"whattofarm/internal/db/dbcollection"
	"whattofarm/internal/db/dbservise"
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

	client, err := dbclient.Connect(USER, PASSWORD, HOST)
	if err != nil {
		logger.Fatalf("MongoDB client", err)
	}

	collection, err := dbcollection.Connect(DATABASE, COLLECTION, client)
	if err != nil {
		logger.Fatalf("MongoDB collection", err)
	}

	service := dbservise.NewService(collection)
	defer func() {
		if err = service.Disconnect(10 * time.Second); err != nil {
			logger.Errorf("MongoDB client disconnection error: ", err)
		}
	}()

	e.GET("/counter.gif", update.NewHandler(service))

	logger.Fatal(e.Start(ADDR))
}
