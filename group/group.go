package group

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	groupURLFormat   = "http[s]?://www.douban.com/group/(.*?)/discussion$"
	groupURLRegexp   *regexp.Regexp
	ErrorBadGroupURL = errors.New("bad group url, format: " + groupURLFormat)
)

func init() {
	groupURLRegexp, _ = regexp.Compile(groupURLFormat)
}

// 获取豆瓣小组的内容 start 表示从第几条开始 默认 0
// 返回 25 条结果的网页内容
func GetGroup(url string, start ...int) ([]byte, error) {
	if !groupURLRegexp.MatchString(url) {
		return nil, ErrorBadGroupURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
