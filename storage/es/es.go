package es

import (
	"encoding/json"
	"github.com/soyking/douban-rent-tools/group"
	"github.com/soyking/douban-rent-tools/storage"
	"gopkg.in/olivere/elastic.v3"
	"strings"
	"time"
)

const (
	TOPIC_TYPE = "topic"
)

type ElasticSearchStorage struct {
	client *elastic.Client
	index  string
}

func NewElasticSearchStorage(addr, index string) (*ElasticSearchStorage, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {
		return nil, err
	}

	exist, err := client.IndexExists(index).Do()
	if err != nil {
		return nil, err
	}

	if !exist {
		_, err := client.CreateIndex(index).Do()
		if err != nil {
			return nil, err
		}
		// mapping
		_, err = client.PutMapping().Index(index).Type(TOPIC_TYPE).BodyString(mappings).Do()
		if err != nil {
			return nil, err
		}
	}

	return &ElasticSearchStorage{
		client: client,
		index:  index,
	}, nil
}

func (e *ElasticSearchStorage) Save(topics []*group.Topic) error {
	bulkService := e.client.Bulk()
	for _, topic := range topics {
		id := topic.URL
		topic.URL = ""
		topicIndex := elastic.
			NewBulkIndexRequest().
			Index(e.index).
			Type(TOPIC_TYPE).
			Id(id).
			Doc(topic)
		bulkService.Add(topicIndex)
	}
	_, err := bulkService.Do()

	return err
}

/*
es query:

GET db_rent/topic/_search
{
  "from": 0,
  "size": 1,
  "query": {
    "bool": {
      "must": [
        {
          "bool": {
            "should": [
              {
                "match": {
                  "topic_content.content": "北京 小区"
                }
              },
              {
                "match": {
                  "title": "北京 小区"
                }
              }
            ]
          }
        },
        {
          "range": {
            "last_reply_time": {
              "gte": "2016-09-12T21:00:00",
              "lte": "2016-09-12T22:29:00"
            }
          }
        },
        {
          "range": {
            "topic_content.update_time": {
              "gte": "2016-09-11T08:00:00",
              "lte": "2016-09-12T21:00:00"
            }
          }
        }
      ]
    }
  },
  "sort": [
    {
      "reply": {
        "order": "desc"
      }
    }
  ]
}

*/

func getTimeFormat(ts int64) string {
	// 时间格式
	return time.Unix(ts, 0).Format("2006-01-02T15:04:05")
}

func newTimeRange(eq []elastic.Query, field string, from, to int64) []elastic.Query {
	if from == 0 && to == 0 {
		return eq
	}
	rangeQuery := elastic.NewRangeQuery(field)
	if from != 0 {
		rangeQuery.Gte(getTimeFormat(from))
	}
	if to != 0 {
		rangeQuery.Lte(getTimeFormat(to))
	}
	return append(eq, rangeQuery)
}

func (e *ElasticSearchStorage) Query(r *storage.QueryRequest) (int, []group.Topic, error) {
	searchService := e.client.Search().Index(e.index).Type(TOPIC_TYPE)

	var queries []elastic.Query
	queries = newTimeRange(queries, "last_reply_time", r.FromLastReplyTime, r.ToLastReplyTime)
	queries = newTimeRange(queries, "topic_content.update_time", r.FromUpdateTime, r.ToUpdateTime)
	if len(r.Keywords) != 0 {
		keywords := strings.Join(r.Keywords, " ")
		// 标题 或者 帖子内容包含，bool.should
		queries = append(queries,
			elastic.NewBoolQuery().Should(
				elastic.NewMatchQuery("title", keywords),
				elastic.NewMatchQuery("topic_content.content", keywords),
			),
		)
	}
	// 时间，内容同时满足，bool.must
	searchService = searchService.Query(elastic.NewBoolQuery().Must(queries...))

	if len(r.SortedBy) == 0 {
		searchService = searchService.Sort("last_reply_time", false)
	} else {
		for _, sort := range r.SortedBy {
			asc := true
			// -表示降序
			if len(sort) > 0 && sort[0] == '-' {
				sort = sort[1:]
				asc = false
			}
			searchService = searchService.Sort(sort, asc)
		}
	}

	if r.Page != 0 || r.Size != 0 {
		// 如果都未0，es默认返回10条数据，与mongo不同
		from := r.Page * r.Size
		searchService = searchService.From(from).Size(r.Size)
	}

	result, err := searchService.Do()
	if err != nil {
		return 0, nil, err
	}

	topics := []group.Topic{}
	for i := range result.Hits.Hits {
		topic := new(group.Topic)
		topicBytes, err := result.Hits.Hits[i].Source.MarshalJSON()
		if err != nil {
			return 0, nil, err
		}

		err = json.Unmarshal(topicBytes, topic)
		if err != nil {
			return 0, nil, err
		}

		topics = append(topics, *topic)
	}

	return int(result.Hits.TotalHits), topics, nil
}
