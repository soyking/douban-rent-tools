package group

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

type Topic struct {
	URL           string
	Titile        string
	AuthorURL     string
	Author        string
	Reply         int
	LastReplyTime time.Time
}

// 从网页内容获取 Topic
func GetTopics(content []byte) ([]*Topic, error) {
	r, err := regexp.Compile(`<tr class="">([\w\W]*?)</tr>`)
	if err != nil {
		return nil, err
	}

	rets := r.FindAllSubmatch(content, -1)
	topics := []*Topic{}
	for i := range rets {
		t, err := GetTopic(rets[i][1])
		if err != nil {
			return nil, err
		}
		topics = append(topics, t)
	}

	return topics, nil
}

// 解析出每个片段中的 Topic
func GetTopic(topicContent []byte) (*Topic, error) {
	r, err := regexp.Compile(`<a href="(.*?)" title="(.*?)"([\w\W]*?)<a href="(.*?)" class="">(.*?)</a>([\w\W]*?)<td nowrap="nowrap" class="">([0-9]*?)</td>([\w\W]*?)<td nowrap="nowrap" class="time">(.*?)</td>`)
	if err != nil {
		return nil, err
	}

	rets := r.FindSubmatch(topicContent)
	if len(rets) != 10 {
		return nil, errors.New("bad topic block")
	}

	var reply int
	replyStr := string(rets[7])
	if replyStr == "" {
		reply = 0
	} else {
		reply, err = strconv.Atoi(replyStr)
		if err != nil {
			return nil, err
		}
	}

	// 时间格式是 08-31 23:42 加上当前年份方便解析成 time.Time
	lastReplyTime, err := time.Parse("2006-01-02 15:04", strconv.Itoa(time.Now().Year())+"-"+string(rets[9]))
	if err != nil {
		return nil, err
	}

	return &Topic{
		URL:           string(rets[1]),
		Titile:        string(rets[2]),
		AuthorURL:     string(rets[4]),
		Author:        string(rets[5]),
		Reply:         reply,
		LastReplyTime: lastReplyTime,
	}, nil
}
