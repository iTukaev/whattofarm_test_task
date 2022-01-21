package dbservise

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"whattofarm/internal/db/dbclient"
)

// DBStruct is used as structure MongoDB document
type DBStruct struct {
	//ID primitive.ObjectID `bson:"_id,omitempty"`
	TimeStamp primitive.Timestamp `bson:"timestamp,omitempty"`
	Total int `bson:"total,omitempty"`
	Actions map[string]*SubCountries `bson:"actions,omitempty"`
	Countries map[string]*SubActions `bson:"countries,omitempty"`
}

type SubCountries struct {
	Total int `bson:"total,omitempty"`
	Countries map[string]*TotalCounter `bson:"countries,omitempty"`
}

type SubActions struct {
	Total int `bson:"total,omitempty"`
	Actions map[string]*TotalCounter `bson:"actions,omitempty"`
}

type TotalCounter struct {
	Total int `bson:"total,omitempty"`
}

// NewDBStruct create new empty DBStruct
func NewDBStruct(timestamp int) *DBStruct {
	return &DBStruct{
		TimeStamp: primitive.Timestamp{
			T: uint32(timestamp),
		},
		Actions:   make(map[string]*SubCountries),
		Countries: make(map[string]*SubActions),
	}
}

type Service struct {
	client *mongo.Client
	sync.Mutex
	data   *DBStruct
	database string
	collection string
}

//// Service implement:
//// func (s *Service) Update(action, country string) error
//// func (s *Service) Disconnect(timeout time.Duration) error
//// func (s *Service) GetDocumentID() error
//// func (s *Service) GetData() (string, error)
//type Service interface {
//	Update(action, country string)
//	Disconnect(timeout time.Duration) error
//	GetData(timeBegin, timeEnd string) ([]byte, error)
//	NewBean(timestamp int) error
//}

// NewService return new instance of Service as a Service interface
// and <nil> if all OK.
// Return error, if connecting to MongoDB return error.
func NewService(user, password, host, database, collection string, timestamp int) (*Service, error) {
	client, err := dbclient.Connect(user, password, host)
	if err != nil {
		return nil, err
	}
	
	return &Service{
		client: client,
		data:   NewDBStruct(timestamp),
		database: database,
		collection: collection,
	}, nil
}