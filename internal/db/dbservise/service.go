package dbservise

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)


type DBStruct struct {
	ID primitive.ObjectID `bson:"_id"`
	Total int `bson:"total"`
	sync.Mutex
	Actions map[string]*TotalCounter `bson:"actions"`
	Countries map[string]*TotalCounter `bson:"countries"`
}

type TotalCounter struct {
	Total int `bson:"total"`
}

func NewDBStruct() *DBStruct {
	return &DBStruct{
		Actions: make(map[string]*TotalCounter),
		Countries: make(map[string]*TotalCounter),
	}
}


type service struct {
	coll *mongo.Collection
	data *DBStruct
	database string
	collection string
}


type Service interface {
	Update(action, country string) error
	Disconnect(timeout time.Duration) error
}

func NewService(collection *mongo.Collection) Service {
	return &service{
		coll: collection,
		data: NewDBStruct(),
	}
}