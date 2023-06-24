package main

import sentry_helper "github.com/evenyosua18/ego-util/tracing/sentry-helper"

func main() {
	//initialize sentry
	flushFunction, err := sentry_helper.InitializeSentry("https://9adf624ee89d40b1b06e7dadeac646ca@o4504592004415488.ingest.sentry.io/4505322347626496", "test")
	if err != nil {
		panic(err)
	}
	defer flushFunction("2")
}
