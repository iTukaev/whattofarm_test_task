package dbservise

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Payload is a structure for MongoDB document
type Payload struct {
	Total int `json:"total"`
	Actions map[string]*SubCountries `json:"actions"`
	Countries map[string]*SubActions `json:"countries"`
}

type Timestamp struct {
	timestamp primitive.Timestamp `bson:"timestamp"`
}

// GetData return MongoDB's document as a JSON string
// and <nil> if all OK.
// Return error, if search or marshalling are incorrect
func (s *service) GetData(timeBegin, timeEnd string) ([]byte, error) {
	collection := s.client.Database(s.database).Collection(s.collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	timeMin, err := time.Parse("2006-01-02_15_-07", timeBegin)
	if err != nil {
		fmt.Println("begin error")
		return nil, errors.New("invalid time")
	}
	timeMax, err := time.Parse("2006-01-02_15_-07", timeEnd)
	if err != nil {
		fmt.Println("end error")
		return nil, errors.New("invalid time")
	}

	// limiting options
	optsMin := primitive.Timestamp{T: uint32(timeMin.UTC().Unix())}
	optsMax := primitive.Timestamp{T: uint32(timeMax.UTC().Unix())}

	cursor, err := collection.Find(ctx, bson.M{"timestamp":bson.M{"$gte":optsMin, "$lte":optsMax}})
	if err != nil {
		return nil, fmt.Errorf("MongoDB document getiing error: %w", err)
	}
	payload, err := aggregate(ctx, cursor)
	if err != nil {
		return nil, fmt.Errorf("MongoDB document %w", err)
	}

	result, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("MongoDB document marshalling error: %w", err)
	}

	return result, nil
}

// aggregate results' structures
func aggregate(ctx context.Context, cursor *mongo.Cursor) (*Payload, error) {
	payload := &Payload{
		Actions: make(map[string]*SubCountries),
		Countries: make(map[string]*SubActions),
	}

	for cursor.Next(ctx) {
		buf := &Payload{
			Actions: make(map[string]*SubCountries),
			Countries: make(map[string]*SubActions),
		}
		if err := cursor.Decode(buf); err != nil {
			return nil, fmt.Errorf("decoding error: %w", err)
		}
		payload.Total += buf.Total
		// подсчёт суммы структуры действий
		// если в результирующей структуре такого значения ещё нет - создаёт новое
		for keyAct, valAct := range buf.Actions {

			if _, ok := payload.Actions[keyAct]; !ok {
				payload.Actions[keyAct] = new(SubCountries)
				payload.Actions[keyAct].Countries = make(map[string]*TotalCounter)
			}
			payload.Actions[keyAct].Total += valAct.Total
			// подсчёт суммы вложенной структуры стран
			// если в результирующей структуре такого значения ещё нет - создаёт новое
			for keyCnt, valCnt := range valAct.Countries {
				if _, ok := payload.Actions[keyAct].Countries[keyCnt]; !ok {
					payload.Actions[keyAct].Countries[keyCnt] = new(TotalCounter)
				}
				payload.Actions[keyAct].Countries[keyCnt].Total += valCnt.Total
			}
		}
		// подсчёт суммы структуры стран
		// если в результирующей структуре такого значения ещё нет - создаёт новое
		for keyCnt, valCnt := range buf.Countries {
			if _, ok := payload.Countries[keyCnt]; !ok {
				payload.Countries[keyCnt] = new(SubActions)
				payload.Countries[keyCnt].Actions = make(map[string]*TotalCounter)
			}
			// подсчёт суммы вложенной структуры действий
			// если в результирующей структуре такого значения ещё нет - создаёт новое
			for keyAct, valAct := range valCnt.Actions {
				if _, ok := payload.Countries[keyCnt].Actions[keyAct]; !ok {
					payload.Countries[keyCnt].Actions[keyAct] = new(TotalCounter)
				}
				payload.Countries[keyCnt].Actions[keyAct].Total += valAct.Total
			}
			payload.Countries[keyCnt].Total += valCnt.Total
		}
	}
	return payload, nil
}