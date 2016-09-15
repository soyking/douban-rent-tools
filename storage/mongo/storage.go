package mongo

import (
	"github.com/soyking/douban-rent-tools/group"
	"github.com/soyking/douban-rent-tools/storage"
	"gopkg.in/mgo.v2/bson"
	"time"
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
	return &MongoDBStorage{m}, nil
}

func (m *MongoDBStorage) Save(topics []*group.Topic) error {
	pairs := []interface{}{}
	for i := range topics {
		pairs = append(pairs, bson.M{"_id": topics[i].URL}, topics[i])
	}
	_, err := m.mongoDBHandler.BulkUpsert(pairs...)
	return err
}

const ANY = ".*"

func getRegex(keyword string) string {
	return ANY + keyword + ANY
}

func getTimeQuery(queries []bson.M, timestamp int64, field string, cond string) []bson.M {
	if timestamp != 0 {
		// mongo 时间查询需要 utc 时间，要加上时区时间差
		timestamp = timestamp + 8*60*60
		return append(
			queries,
			bson.M{field: bson.M{cond: time.Unix(timestamp, 0).UTC()}},
		)
	}
	return queries
}

func (m *MongoDBStorage) Query(r *storage.QueryRequest) (int, []group.Topic, error) {
	var queries []bson.M

	var keywordQueries []bson.M
	for _, keyword := range r.Keywords {
		// 搜索 title content 中包含关键字的帖子
		// MongoDB 3.2 中文搜索支持比较麻烦，所以只做正则
		keywordQueries = append(
			keywordQueries,
			bson.M{"title": bson.M{"$regex": getRegex(keyword)}},
			bson.M{"topic_content.content": bson.M{"$regex": getRegex(keyword)}},
		)
	}
	if len(keywordQueries) != 0 {
		queries = append(
			queries,
			bson.M{"$or": keywordQueries},
		)
	}

	queries = getTimeQuery(queries, r.FromUpdateTime, "topic_content.update_time", "$gte")
	queries = getTimeQuery(queries, r.ToUpdateTime, "topic_content.update_time", "$lte")
	queries = getTimeQuery(queries, r.FromLastReplyTime, "last_reply_time", "$gte")
	queries = getTimeQuery(queries, r.ToLastReplyTime, "last_reply_time", "$lte")

	if len(r.SortedBy) == 0 {
		// 默认按最晚回复时间排序
		r.SortedBy = []string{"-last_reply_time"}
	}

	var query bson.M
	if len(queries) > 0 {
		query = bson.M{"$and": queries}
	}

	selector := bson.M{
		"topic_content.content":  0,
		"topic_content.pic_urls": 0,
	}
	var result []group.Topic
	count, err := m.mongoDBHandler.FindAll(
		query,
		selector,
		r.Page,
		r.Size,
		&result,
		r.SortedBy...,
	)
	return count, result, err
}
