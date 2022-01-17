package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
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

	service, err := dbservise.NewService(USER, PASSWORD, HOST, DATABASE, COLLECTION, timeHours())
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		if err = service.Disconnect(5*time.Second); errors.Unwrap(err) != nil {
			logger.Error(err)
		} else {
			logger.Infof("MongoDB disconnected:", "host: "+ HOST)
		}
	}()
	logger.Infof("MongoDB connected:", "host: "+ HOST, "user: "+ USER)

	e.GET("/counter.gif", update.NewHandler(service))
	e.GET("/", getdata.NewHandler(service))

	go func() {
		logger.Fatal(e.Start(ADDR))
	}()

	c := cron.New()
	//_, err = c.AddFunc("@hourly", func() {
	//	if err := service.NewBin(timeHours()); err != nil {
	//		logger.Error(err)
	//	}
	//})
	//if err != nil {
	//	logger.Fatalf("Cron AddFunc error:", err)
	//}
	_, err = c.AddFunc("*/1 * * * *", func() {
		if err := service.NewBin(timeHours()); err != nil {
			logger.Error(err)
		}
	})
	if err != nil {
		fmt.Println("cron error")
	}
	c.Start()

	<-ctx.Done()
	c.Stop()
	if err = e.Close(); err != nil {
		logger.Errorf("Server stopped with error:", err)
	} else {
		logger.Infof("Server stopped:", "host: "+ HOST)
	}

	for i := 0; i < 3; i++ {
		if err := service.NewBin(timeHours()); err != nil {
			logger.Error(err)
			time.Sleep(3*time.Second)
			continue
		}
		return
	}
}


// stop application by ^C
func stop(cancel context.CancelFunc) {
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
	<- exitCh
	cancel()
}

// timeHours resets minutes and seconds and returns integer time in Unix format
func timeHours() int {
	timeNow := time.Now()
	return int(timeNow.UTC().Unix()) - timeNow.Minute() * 60 - timeNow.Second()
}