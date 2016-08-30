package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	content, err := ioutil.ReadFile("./group.html")
	if err != nil {
		log.Fatal(err)
	}

	//content, err := GetGroup("https://www.douban.com/group/beijingzufang/discussion?start=0")
	//if err != nil {
	//	log.Fatal(err)
	//}

	topics, err := GetTopics(content)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.MarshalIndent(topics, "", "    ")
	println(string(b))
}

func GetGroup(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

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

type Topic struct {
	URL           string
	Titile        string
	AuthorURL     string
	Author        string
	Reply         int
	LastReplyTime time.Time
}

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
