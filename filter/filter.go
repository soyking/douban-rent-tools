package filter

import "github.com/soyking/douban-rent-tools/group"

type Filter func([]*group.Topic) []*group.Topic

// 与逻辑过滤器，需要满足所有过滤要求
func NewFilter(filters ...FilterFunc) Filter {
	return func(topics []*group.Topic) []*group.Topic {
		filteredTopics := []*group.Topic{}
		for _, topic := range topics {
			valid := true
			for _, filter := range filters {
				if !filter(topic) {
					valid = false
					break
				}
			}
			if valid {
				filteredTopics = append(filteredTopics, topic)
			}
		}

		return filteredTopics
	}
}
