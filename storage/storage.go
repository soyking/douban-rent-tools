package storage

import "github.com/soyking/douban-rent-tools/group"

type QueryRequest struct {
	Page              int      `json:"page"`
	Size              int      `json:"size"`
	SortedBy          []string `json:"sorted_by"`
	Keywords          []string `json:"keywords"`
	FromUpdateTime    int64    `json:"from_update_time"`
	ToUpdateTime      int64    `json:"to_update_time"`
	FromLastReplyTime int64    `json:"from_last_reply_time"`
	ToLastReplyTime   int64    `json:"to_last_reply_time"`
}

type Storage interface {
	Save([]*group.Topic) error
	// 返回（总数，帖子，错误）
	Query(r *QueryRequest) (int, []group.Topic, error)
}
