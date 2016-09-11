package mongo

import (
	"gopkg.in/mgo.v2"
	"strings"
)

type MongoDBHandler struct {
	database   string
	collection string
	session    *mgo.Session
}

func NewMongoDBHandler(addr, username, password, database, collection string) (*MongoDBHandler, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(addr, ","),
		Username: username,
		Password: password,
		Database: database,
	}
	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	err = sess.Ping()
	if err != nil {
		return nil, err
	}

	return &MongoDBHandler{
		database:   database,
		collection: collection,
		session:    sess,
	}, nil
}

func (m *MongoDBHandler) Upsert(selector, update interface{}) error {
	sess := m.session.Copy()
	defer sess.Close()
	_, err := sess.DB(m.database).C(m.collection).Upsert(selector, update)
	return err
}

func (m *MongoDBHandler) EnsureIndex(keys ...string) error {
	sess := m.session.Copy()
	defer sess.Close()
	return sess.DB(m.database).C(m.collection).EnsureIndex(
		mgo.Index{
			Key:        keys,
			Background: true,
		},
	)
}

func (m *MongoDBHandler) FindAll(query interface{}, page, size int, result interface{}, sortedBy ...string) (int, error) {
	sess := m.session.Copy()
	defer sess.Close()
	q := sess.DB(m.database).C(m.collection).Find(query)
	count, err := q.Count()
	if err != nil {
		return 0, err
	}
	return count, q.Sort(sortedBy...).Skip(page * size).Limit(size).All(result)
}
