package group

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	TOPIC_CONTENT_NOT_FOUND = "呃...你想要的东西不在这儿"
)

var (
	EMPTY_WORD = []string{
		" ",
		"\n",
	}
	emptyReplacer *strings.Replacer

	ErrorTopicDelete = errors.New("topic has been deleted")

	httpClient = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return ErrorTopicDelete
		},
	}
)

func init() {
	oldnew := []string{}
	for _, w := range EMPTY_WORD {
		// 替换为空
		oldnew = append(oldnew, w, "")
	}
	emptyReplacer = strings.NewReplacer(oldnew...)
}

type TopicContent struct {
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
	Content    string    `json:"content" bson:"content"`
	WithPic    bool      `json:"with_pic" bson:"with_pic"`
	PicURLs    []string  `json:"pic_urls" bson:"pic_urls"`
	Like       int       `json:"like" bson:"like"`
}

func GetTopicContent(url string) (*TopicContent, error) {
	t, err := getTopicContent(url)
	if err != nil {
		return nil, errors.New("topic: " + url + "; " + err.Error())
	}

	return t, nil
}

func getTopicContent(url string) (*TopicContent, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusFound {
		return nil, ErrorTopicDelete
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	topicContent, err := ParseTopicContent(doc)
	if err != nil {
		return nil, err
	}

	return topicContent, nil
}

func ParseTopicContent(doc *goquery.Document) (*TopicContent, error) {
	wholeContent := doc.Find("html body div#wrapper div#content div.grid-16-8.clearfix div.article div.topic-content.clearfix")

	// TODO: FULL TITLE
	updateTimeStr := wholeContent.Find("div.topic-doc h3 span.color-green").Text()
	if updateTimeStr == "" {
		// 存在蓝色状态的帖子，感觉是一种预删除的状态，页面结构不一样，这里作帖子被删除处理
		return nil, ErrorTopicDelete
	}
	updateTime, err := time.Parse("2006-01-02 15:04:05", updateTimeStr)
	if err != nil {
		return nil, err
	}

	topicContent := wholeContent.Find("div.topic-doc div#link-report div.topic-content")
	if topicContent.Length() == 0 {
		return nil, errors.New("without content")
	}

	content := []string{}
	topicContent.Find("p").Each(func(i int, s *goquery.Selection) {
		content = append(content, emptyReplacer.Replace(s.Text()))
	})

	picBlock := topicContent.Find("div.topic-figure.cc")
	withPic := false
	picURLs := []string{}
	if picBlock.Length() > 0 {
		withPic = true
		picBlock.Each(func(i int, s *goquery.Selection) {
			picURL, exist := s.Find("img").Attr("src")
			if exist && picURL != "" {
				picURLs = append(picURLs, picURL)
			}
		})
	}

	likeStr := wholeContent.Find("div#sep.sns-bar div.sns-bar-fav span.fav-num a").Text()
	like := 0
	if likeStr != "" {
		like, err = strconv.Atoi(strings.TrimRight(likeStr, "人"))
		if err != nil {
			return nil, err
		}
	}

	return &TopicContent{
		UpdateTime: updateTime,
		Content:    strings.Join(content, ""),
		WithPic:    withPic,
		PicURLs:    picURLs,
		Like:       like,
	}, nil
}
