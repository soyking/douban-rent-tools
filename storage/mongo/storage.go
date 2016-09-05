package mongo

import "github.com/soyking/douban-rent-tools/group"

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

	return &MongoDBStorage{m}, nil
}

func (m *MongoDBStorage) Save(topics []*group.Topic) error {
	ts := []interface{}{}
	for i := range topics {
		ts = append(ts, topics[i])
	}
	return m.mongoDBHandler.Insert(ts...)
}

func (m *MongoDBStorage) Query(q interface{}) ([]group.Topic, error) {
	return []group.Topic{}, nil
}
