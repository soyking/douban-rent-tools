package group

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"os"
	"testing"
)

const (
	TOPICS_NUMBER = 25
)

func getLocalContent() (*goquery.Document, error) {
	f, err := os.Open("./group_test.html")
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(f)
}

func testTopics(t *testing.T, doc *goquery.Document) {
	topics, err := GetTopics(doc)
	if err != nil {
		t.Error(err)
	} else if len(topics) != TOPICS_NUMBER {
		t.Errorf("topics number is err, should be %d but %d\n", TOPICS_NUMBER, len(topics))
	} else {
		b, _ := json.MarshalIndent(topics, "", "    ")
		t.Log(string(b))
	}
}

func TestGetTopics(t *testing.T) {
	content, err := getLocalContent()
	if err != nil {
		t.Error(err)
	} else {
		testTopics(t, content)
	}
}

func TestGetTopics2(t *testing.T) {
	content, err := GetGroup("beijingzufang", "25")
	if err != nil {
		t.Error(err)
	} else {
		testTopics(t, content)
	}
}
