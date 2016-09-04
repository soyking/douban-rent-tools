package group

import "github.com/PuerkitoBio/goquery"

func getGroupURL(name string, start ...string) string {
	s := "0"
	if len(start) > 0 {
		s = start[0]
	}
	return "https://www.douban.com/group/" + name + "/discussion?start=" + s
}

// 获取豆瓣小组的内容 start 表示从第几条开始 默认 0
// 返回 25 条结果的网页内容
func GetGroup(name string, start ...string) (*goquery.Document, error) {
	return goquery.NewDocument(getGroupURL(name, start...))
}
