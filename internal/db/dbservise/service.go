package dbservise

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"whattofarm/internal/db/dbclient"
)

// DBStruct is used as structure MongoDB document
type DBStruct struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Total int `bson:"total,omitempty"`
	sync.Mutex `bson:"-"`
	Actions map[string]*TotalCounter `bson:"actions,omitempty"`
	Countries map[string]*TotalCounter `bson:"countries,omitempty"`
}

// TotalCounter is a structure for counting parameters
type TotalCounter struct {
	Total int `bson:"total"`
}

// NewDBStruct create new empty DBStruct
func NewDBStruct() *DBStruct {
	return &DBStruct{
		Total: 0,
		Actions:   make(map[string]*TotalCounter),
		Countries: make(map[string]*TotalCounter),
	}
}

type Service struct {
	client *mongo.Client
	data   *DBStruct
	database string
	collection string
}

// NewService return new instance of Service as a Service interface
// and <nil> if all OK.
// Return error, if connecting to MongoDB return error.
func NewService(user, password, host, database, collection string) (*Service, error) {
	client, err := dbclient.Connect(user, password, host)
	if err != nil {
		return nil, err
	}
	
	return &Service{
		client: client,
		data:   NewDBStruct(),
		database: database,
		collection: collection,
	}, nil
}