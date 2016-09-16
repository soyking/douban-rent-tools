package router

import "strings"

const (
	SUBWAY_PREFIX = "地铁："
	LINE          = "号线"
)

var subwayToHan = map[string]string{
	"1":  "一",
	"2":  "二",
	"4":  "四",
	"5":  "五",
	"6":  "六",
	"7":  "七",
	"8":  "八",
	"9":  "九",
	"10": "十",
	"13": "十三",
	"14": "十四",
	"15": "十五",
}

// 输入 地铁：5
func SubwayCond(keyword string) bool {
	return strings.HasPrefix(keyword, SUBWAY_PREFIX)
}

func subwayName(line string) string {
	return line + LINE
}

// 返回 [5号线, 五号线]
func SubwayExpand(line string) []string {
	line = strings.TrimLeft(line, SUBWAY_PREFIX)
	r := []string{subwayName(line)}
	hanLine, ok := subwayToHan[line]
	if ok {
		r = append(r, subwayName(hanLine))
	}
	return r
}
