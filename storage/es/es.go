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
          "query_string": {
            "fields": [
              "title",
              "topic_content.content"
            ],
            "analyze_wildcard": true,
            "default_operator": "OR",
            "query": "*三家* *六号线*"
          }
        },
        {
          "range": {
            "last_reply_time": {
              "gte": "2016-09-16T12:21:00",
              "lte": "2016-09-16T12:21:00"
            }
          }
        },
        {
          "range": {
            "topic_content.update_time": {
              "gte": "2016-09-01T08:00:00",
              "lte": "2016-09-16T21:00:00"
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

const WILDCARD = "*"

func (e *ElasticSearchStorage) Query(r *storage.QueryRequest) (int, []group.Topic, error) {
	searchService := e.client.Search().Index(e.index).Type(TOPIC_TYPE)

	var queries []elastic.Query
	queries = newTimeRange(queries, "last_reply_time", r.FromLastReplyTime, r.ToLastReplyTime)
	queries = newTimeRange(queries, "topic_content.update_time", r.FromUpdateTime, r.ToUpdateTime)
	if len(r.Keywords) != 0 {
		for i := range r.Keywords {
			r.Keywords[i] = WILDCARD + r.Keywords[i] + WILDCARD
		}
		keywords := strings.Join(r.Keywords, " ")
		// 关键词查询使用 contain 式查询，所以对应的 field 设置为 not_analyzed
		queries = append(queries,
			elastic.NewQueryStringQuery(keywords).
				AnalyzeWildcard(true).
				DefaultOperator("OR").
				Field("title").
				Field("topic_content.content"),
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

	fsc := elastic.NewFetchSourceContext(true)
	fsc.Exclude("topic_content.content", "topic_content.pic_urls")
	searchService.FetchSourceContext(fsc)

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

		topic.URL = result.Hits.Hits[i].Id
		topics = append(topics, *topic)
	}

	return int(result.Hits.TotalHits), topics, nil
}
