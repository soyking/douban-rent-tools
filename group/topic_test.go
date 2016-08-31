package group

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

const (
	TOPICS_NUMBER = 25
)

func getLocalContent() ([]byte, error) {
	return ioutil.ReadFile("./group_test.html")
}

func testTopics(t *testing.T, content []byte) {
	topics, err := GetTopics(content)
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
	content, err := GetGroup("https://www.douban.com/group/beijingzufang/discussion", 25)
	if err != nil {
		t.Error(err)
	} else {
		testTopics(t, content)
	}
}
