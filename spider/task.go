package main

import (
	"github.com/soyking/douban-rent-tools/group"
	"log"
	"strings"
	"sync"
	"time"
)

func runTask() {
	groups := strings.Split(groupNames, ",")
	log.Printf("crawling groups: %s\n", groups)

	tick := time.Tick(time.Duration(frequency) * time.Second)
	log.Println("...start task...")
	count := 1
	for _ = range tick {
		log.Printf("\ttask %d\n", count)
		var wg sync.WaitGroup
		taskChan := make(chan int, groupsConcurrency)

		for _, g := range groups {
			taskChan <- 1
			wg.Add(1)

			go func(groupName string) {
				topics, err := group.GetTopics(groupName, topicsConcurrency)
				if err != nil {
					log.Printf("\t\t[Fail] fetch group: %s err: %s\n", groupName, err.Error())
				}
				topics = topicsFilter(topics)
				err = store.Save(topics)
				if err != nil {
					log.Printf("\t\t[Fail] save group: %s err: %s\n", groupName, err.Error())
				} else {
					log.Printf("\t\t[SUCCESS] group: %s topics %d\n", groupName, len(topics))
				}
				wg.Done()
				<-taskChan
			}(g)
		}

		wg.Wait()
		count += 1
	}
}
