package storage

import "github.com/soyking/douban-rent-tools/group"

type Storage interface {
	Save(*group.Topic) error
	Query(interface{}) ([]group.Topic, error)
}
