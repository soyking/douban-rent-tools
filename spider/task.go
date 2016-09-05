package main

import (
	"github.com/soyking/douban-rent-tools/group"
	"log"
	"strings"
	"time"
)

func runTask() {
	groups := strings.Split(groupNames, ",")
	tick := time.Tick(time.Duration(frequency) * time.Second)
	log.Println("...start task...")
	count := 1
	for _ = range tick {
		// TODO: MULTI THREAD CHOICE, INCLUDING GET TOPICS
		log.Printf("\ttask %d\n", count)
		for _, g := range groups {
			log.Printf("\t\tgroup %s\n", g)
			topics, err := group.GetTopics(g)
			if err != nil {
				// TODO: BETTER LOGGER
				log.Printf("\t\t\t[Fail] group: %s err: %s\n", g, err.Error())
			}
			store.Save(topics)
			log.Printf("\t\t\t[SUCCESS] topics %d\n", len(topics))
		}
		count += 1
	}
}
