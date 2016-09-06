package mongo

import (
	"github.com/soyking/douban-rent-tools/group"
	"gopkg.in/mgo.v2/bson"
)

const (
	TOPIC_COLLECTION = "topic"
)

type MongoDBStorage struct {
	mongoDBHandler *MongoDBHandler
}

func NewMongoDBStorage(addr, username, password, database string) (*MongoDBStorage, error) {
	m, err := NewMongoDBHandler(addr, username, password, database, TOPIC_COLLECTION)
	if err != nil {
		return nil, err
	}

	// TODO: INDEX
	return &MongoDBStorage{m}, nil
}

func (m *MongoDBStorage) Save(topics []*group.Topic) error {
	for _, topic := range topics {
		err := m.mongoDBHandler.Upsert(
			bson.M{"_id": topic.URL},
			topic,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MongoDBStorage) Query(q interface{}) ([]group.Topic, error) {
	return []group.Topic{}, nil
}
