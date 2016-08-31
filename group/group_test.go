package group

import "testing"

func TestGetGroup(t *testing.T) {
	badURLs := []string{
		"https://www.douban.com/group/discussion",
		"https://www.douban.com/group/beijingzufang/discussion?abc=def",
	}

	for _, url := range badURLs {
		_, err := GetGroup(url)
		if err != ErrorBadGroupURL {
			t.Errorf("should be bad url: %s", url)
		}
	}
}

func TestGetGroup2(t *testing.T) {
	content, err := GetGroup("https://www.douban.com/group/beijingzufang/discussion")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(content))
	}
}
