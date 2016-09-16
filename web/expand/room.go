package expand

import "strings"

const (
	ROOM_PREFIX = "房间："
)

var (
	roomToHan = map[string][]string{
		"1":  []string{"一", "单"},
		"2":  []string{"二", "两", "俩"},
		"3":  []string{"三", "仨"},
		"4":  []string{"四"},
		"5":  []string{"五"},
		"6":  []string{"六"},
		"7":  []string{"七"},
		"8":  []string{"八"},
		"9":  []string{"九"},
		"10": []string{"十"},
	}

	roomDescs = []string{
		"居",
		"室",
		"房间",
		"个房间",
	}
)

// 输入 房间：2
func RoomCond(keyword string) bool {
	return strings.HasPrefix(keyword, ROOM_PREFIX)
}

func roomDesc(room string) []string {
	d := []string{}
	for i := range roomDescs {
		d = append(d, room+roomDescs[i])
	}
	return d
}

// 返回 [2居 2室 2房间 2个房间 二居 二室 二房间 二个房间 两居 两室 两房间 两个房间 俩居 俩室 俩房间 俩个房间]
func RoomExpand(room string) []string {
	room = strings.TrimLeft(room, ROOM_PREFIX)
	r := roomDesc(room)
	roomHan, ok := roomToHan[room]
	if ok {
		for _, h := range roomHan {
			r = append(r, roomDesc(h)...)
		}
	}
	return r
}
