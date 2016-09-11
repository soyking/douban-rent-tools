package main

import (
	"flag"
	"github.com/soyking/douban-rent-tools/filter"
	"time"
)

const (
	FLAG_GROUPS_NAME    = "groups"
	FLAG_GROUPS_DEFAULT = "beijingzufang"
	FLAG_GROUPS_USAGE   = "group name, split by ,"

	FLAG_MONGO_ADDR_NAME    = "mg_addr"
	FLAG_MONGO_ADDR_DEFAULT = "127.0.0.1:27017"
	FLAG_MONGO_ADDR_USAGE   = "MongoDB address, split by ,"

	FLAG_MONGO_USERNAME_NAME    = "mg_usr"
	FLAG_MONGO_USERNAME_DEFAULT = ""
	FLAG_MONGO_USERNAME_USAGE   = "MongoDB username"

	FLAG_MONGO_PASSWORD_NAME    = "mg_pwd"
	FLAG_MONGO_PASSWORD_DEFAULT = ""
	FLAG_MONGO_PASSWORD_USAGE   = "MongoDB password"

	FLAG_MONGO_DATABASE_NAME    = "mg_db"
	FLAG_MONGO_DATABASE_DEFAULT = "db_rent"
	FLAG_MONGO_DATABASE_USAGE   = "MongoDB database"

	FLAG_FREQUENCY_NAME    = "freq"
	FLAG_FREQUENCY_DEFUALT = 60
	FLAG_FREQUENCY_USAGE   = "spider frequency(in second)"

	FLAG_GROUPS_CONCURRENCY_NAME    = "g_con"
	FLAG_GROUPS_CONCURRENCY_DEFAULT = 1
	FLAG_GROUPS_CONCURRENCY_USAGE   = "concurrency for groups crawling"

	FLAG_TOPICS_CONCURRENCY_NAME    = "t_con"
	FLAG_TOPICS_CONCURRENCY_DEFAULT = 1
	FLAG_TOPICS_CONCURRENCY_USAGE   = " for topics crawling"

	FLAG_AUTHOR_FILTER_NAME    = "author_filter"
	FLAG_AUTHOR_FILTER_DEFAULT = "NO_FILE"
	FLAG_AUTHOR_FILTER_USAGE   = "author filter file path"

	FLAG_TITLE_FILTER_NAME    = "title_filter"
	FLAG_TITLE_FILTER_DEFAULT = "NO_FILE"
	FLAG_TITLE_FILTER_USAGE   = "title filter file path"

	FLAG_CONTENT_FILTER_NAME    = "content_filter"
	FLAG_CONTENT_FILTER_DEFAULT = "NO_FILE"
	FLAG_CONTENT_FILTER_USAGE   = "content filter file path"

	FLAG_REPLY_FILTER_NAME    = "reply_filter"
	FLAG_REPLY_FILTER_DEFAULT = 0
	FLAG_REPLY_FILTER_USAGE   = "max reply of a topic"

	FLAG_PIC_FILTER_NAME    = "pic_filter"
	FLAG_PIC_FILTER_DEFAULT = false
	FLAG_PIC_FILTER_USAGE   = "topic with picture"

	FLAG_LAST_UPDATE_TIME_FILTER_NAME    = "last_utime_filter"
	FLAG_LAST_UPDATE_TIME_FILTER_DEFAULT = ""
	FLAG_LAST_UPDATE_TIME_FILTER_USAGE   = "last update time filter, format: 2006-01-02 15:04:05"
)

var (
	groupNames string
	frequency  int

	mongoDBAddr     string
	mongoDBUsername string
	mongoDBPassword string
	mongoDBDatabase string

	groupsConcurrency int
	topicsConcurrency int

	authorFilterFile     string
	titleFilterFile      string
	contentFilterFile    string
	replyFilter          int
	picFilter            bool
	lastUpdateTimeFilter string

	topicsFilter filter.Filter
)

func init() {
	flag.StringVar(&groupNames, FLAG_GROUPS_NAME, FLAG_GROUPS_DEFAULT, FLAG_GROUPS_USAGE)
	flag.IntVar(&frequency, FLAG_FREQUENCY_NAME, FLAG_FREQUENCY_DEFUALT, FLAG_FREQUENCY_USAGE)

	flag.StringVar(&mongoDBAddr, FLAG_MONGO_ADDR_NAME, FLAG_MONGO_ADDR_DEFAULT, FLAG_MONGO_ADDR_USAGE)
	flag.StringVar(&mongoDBUsername, FLAG_MONGO_USERNAME_NAME, FLAG_MONGO_USERNAME_DEFAULT, FLAG_MONGO_USERNAME_USAGE)
	flag.StringVar(&mongoDBPassword, FLAG_MONGO_PASSWORD_NAME, FLAG_MONGO_PASSWORD_DEFAULT, FLAG_MONGO_PASSWORD_USAGE)
	flag.StringVar(&mongoDBDatabase, FLAG_MONGO_DATABASE_NAME, FLAG_MONGO_DATABASE_DEFAULT, FLAG_MONGO_DATABASE_USAGE)

	flag.IntVar(&groupsConcurrency, FLAG_GROUPS_CONCURRENCY_NAME, FLAG_GROUPS_CONCURRENCY_DEFAULT, FLAG_GROUPS_CONCURRENCY_USAGE)
	flag.IntVar(&topicsConcurrency, FLAG_TOPICS_CONCURRENCY_NAME, FLAG_TOPICS_CONCURRENCY_DEFAULT, FLAG_TOPICS_CONCURRENCY_USAGE)

	flag.StringVar(&authorFilterFile, FLAG_AUTHOR_FILTER_NAME, FLAG_AUTHOR_FILTER_DEFAULT, FLAG_AUTHOR_FILTER_USAGE)
	flag.StringVar(&titleFilterFile, FLAG_TITLE_FILTER_NAME, FLAG_TITLE_FILTER_DEFAULT, FLAG_TITLE_FILTER_USAGE)
	flag.StringVar(&contentFilterFile, FLAG_CONTENT_FILTER_NAME, FLAG_CONTENT_FILTER_DEFAULT, FLAG_CONTENT_FILTER_USAGE)
	flag.IntVar(&replyFilter, FLAG_REPLY_FILTER_NAME, FLAG_REPLY_FILTER_DEFAULT, FLAG_REPLY_FILTER_USAGE)
	flag.BoolVar(&picFilter, FLAG_PIC_FILTER_NAME, FLAG_PIC_FILTER_DEFAULT, FLAG_PIC_FILTER_USAGE)
	flag.StringVar(&lastUpdateTimeFilter, FLAG_LAST_UPDATE_TIME_FILTER_NAME, FLAG_LAST_UPDATE_TIME_FILTER_DEFAULT, FLAG_LAST_UPDATE_TIME_FILTER_USAGE)

	flag.Parse()

	if frequency <= 0 {
		frequency = FLAG_FREQUENCY_DEFUALT
	}
	if groupsConcurrency <= 0 {
		groupsConcurrency = FLAG_GROUPS_CONCURRENCY_DEFAULT
	}
	if topicsConcurrency <= 0 {
		groupsConcurrency = FLAG_TOPICS_CONCURRENCY_DEFAULT
	}

	filterFuncs := []filter.FilterFunc{}
	if authorFilterFile != FLAG_AUTHOR_FILTER_DEFAULT {
		authors, err := readLines(authorFilterFile)
		if err != nil {
			panic(err)
		}
		filterFuncs = append(filterFuncs, filter.AuthorFilter(authors))
	}
	if titleFilterFile != FLAG_AUTHOR_FILTER_DEFAULT {
		titles, err := readLines(titleFilterFile)
		if err != nil {
			panic(err)
		}
		filterFuncs = append(filterFuncs, filter.TitleFilter(titles))
	}
	if contentFilterFile != FLAG_AUTHOR_FILTER_DEFAULT {
		contents, err := readLines(contentFilterFile)
		if err != nil {
			panic(err)
		}
		filterFuncs = append(filterFuncs, filter.ContentFilter(contents))
	}
	if replyFilter > FLAG_REPLY_FILTER_DEFAULT {
		filterFuncs = append(filterFuncs, filter.ReplyLimitFilter(replyFilter))
	}
	if picFilter {
		// 只对有图片要求的过滤
		filterFuncs = append(filterFuncs, filter.PicFilter(true))
	}
	if lastUpdateTimeFilter != FLAG_LAST_UPDATE_TIME_FILTER_DEFAULT {
		t, err := time.Parse("2006-01-02 15:04:05", lastUpdateTimeFilter)
		if err != nil {
			panic("check your time format: " + err.Error())
		}
		filterFuncs = append(filterFuncs, filter.LastUpdateTimeFilter(t))
	}

	topicsFilter = filter.NewFilter(filterFuncs...)
}
