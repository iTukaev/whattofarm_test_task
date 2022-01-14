package dbservise

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
	"whattofarm/internal/db/dbclient"
)


type DBStruct struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Total int `bson:"total"`
	sync.Mutex `bson:"-"`
	Actions map[string]*TotalCounter `bson:"actions"`
	Countries map[string]*TotalCounter `bson:"countries"`
}

type TotalCounter struct {
	Total int `bson:"total"`
}

func NewDBStruct() *DBStruct {
	return &DBStruct{
		Total: 0,
		Actions:   make(map[string]*TotalCounter),
		Countries: make(map[string]*TotalCounter),
	}
}

type service struct {
	client *mongo.Client
	data   *DBStruct
	database string
	collection string
}


type Service interface {
	Update(action, country string) error
	Disconnect(timeout time.Duration) error
	GetDocumentID() error
	GetData() (string, error)
}

func NewService(user, password, host, database, collection string) (Service, error) {
	client, err := dbclient.Connect(user, password, host)
	if err != nil {
		return nil, err
	}
	
	return &service{
		client: client,
		data:   NewDBStruct(),
		database: database,
		collection: collection,
	}, nil
}