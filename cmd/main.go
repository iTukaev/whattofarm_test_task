package main

import (
	"context"
	"errors"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
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

	ctx, cancel := context.WithCancel(context.Background())
	go stop(cancel)
	start(ctx, e)
}


// start MongoDB services and listener
func start(ctx context.Context, e *echo.Echo) {
	logger := e.Logger

	service, err := dbservise.NewService(USER, PASSWORD, HOST, DATABASE,COLLECTION)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("MongoDB connected:", "host: "+ HOST, "user: "+ USER)

	if err = service.GetDocumentID(); err != nil {
		logger.Fatal(err)
	}

	e.GET("/counter.gif", update.NewHandler(service))
	e.GET("/", getdata.NewHandler(service))

	go func() {
		logger.Fatal(e.Start(ADDR))
	}()

	<-ctx.Done()
	if err = service.Disconnect(5*time.Second); errors.Unwrap(err) != nil {
		logger.Error(err)
	} else {
		logger.Infof("MongoDB disconnected:", "host: "+ HOST)
	}

	if err = e.Close(); err != nil {
		logger.Errorf("Server stopped with error:", err)
	} else {
		logger.Infof("Server stopped:", "host: "+ HOST)
	}
}


// stop application by ^C
func stop(cancel context.CancelFunc) {
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
	<- exitCh
	cancel()
}