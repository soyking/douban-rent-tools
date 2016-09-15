package main

const (
	APP_NAME    = "DOUBAN RENT TOOLS - SPIDER"
	APP_VERSION = "0.0.1"
)

func main() {
	println(APP_NAME + "\t" + APP_VERSION)
	initStorage()
	runTask()
}
