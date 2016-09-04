package group

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"time"
	"strings"
)

type Topic struct {
	URL           string
	Titile        string
	AuthorURL     string
	Author        string
	Reply         int
	LastReplyTime time.Time
	TopicContent  *TopicContent
}

// 从文档树中获取 Topic
func GetTopics(doc *goquery.Document) ([]*Topic, error) {
	var topics []*Topic
	var outErr error
	doc.Find("html body #wrapper div#content div.grid-16-8.clearfix div.article div table.olt tbody tr").
		Each(func(i int, s *goquery.Selection) {
			topic, err := GetTopic(s)
			if err != nil {
				outErr = errors.New("group: " + doc.Url.String() + " #" + strconv.Itoa(i) + "; " + err.Error())
			}

			if topic != nil {
				topics = append(topics, topic)
			}
		},
		)
	return topics, outErr
}

// 解析成自定义的 Topic
func GetTopic(s *goquery.Selection) (*Topic, error) {
	titleBlock := s.Find("td.title")
	if titleBlock.Length() == 0 {
		// 存在非小组话题
		return nil, nil
	}

	authorBlock := titleBlock.Next()
	replyBlock := authorBlock.Next()
	timeBlock := replyBlock.Next()

	titleBlock = titleBlock.Find("a")
	url, exist := titleBlock.Attr("href")
	if !exist {
		return nil, errors.New("without url")
	}
	topicContent, err := GetTopicContent(url)
	if err != nil {
		if strings.Contains(err.Error(), ErrorTopicDelete.Error()) {
			return nil, nil
		}
		return nil, err
	}

	title := titleBlock.Text()
	if title == "" {
		return nil, errors.New("without title")
	}

	authorBlock = authorBlock.Find("a")
	authorURL, exist := authorBlock.Attr("href")
	if !exist {
		return nil, errors.New("without author url")
	}
	author := authorBlock.Text()
	if author == "" {
		return nil, errors.New("without author")
	}

	replyStr := replyBlock.Text()
	reply := 0
	if replyStr != "" {
		var err error
		reply, err = strconv.Atoi(replyStr)
		if err != nil {
			return nil, err
		}
	}

	replyTimeStr := timeBlock.Text()
	if replyTimeStr == "" {
		return nil, errors.New("without last reply time")
	}
	// 时间格式是 08-31 23:42 加上当前年份方便解析成 time.Time
	lastReplyTime, err := time.Parse("2006-01-02 15:04", strconv.Itoa(time.Now().Year())+"-"+replyTimeStr)
	if err != nil {
		return nil, err
	}

	return &Topic{
		URL:           url,
		Titile:        title,
		AuthorURL:     authorURL,
		Author:        author,
		Reply:         reply,
		LastReplyTime: lastReplyTime,
		TopicContent:  topicContent,
	}, nil
}
