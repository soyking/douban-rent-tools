package mongo

import "github.com/soyking/douban-rent-tools/group"

type MongoDBStorage struct {
	mongoDBHandler *MongoDBHandler
}

func NewMongoDBStorage(addr, username, password, database, collection string) (*MongoDBStorage, error) {
	m, err := NewMongoDBHandler(addr, username, password, database, collection)
	if err != nil {
		return nil, err
	}

	return &MongoDBStorage{m}, nil
}

func (m *MongoDBStorage)Save(topic *group.Topic) error {
	return m.mongoDBHandler.Insert(topic)
}

func (m *MongoDBStorage)Query(q interface{}) ([]group.Topic, error) {
	return []group.Topic{}, nil
}
