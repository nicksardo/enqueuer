package main

import (
	"fmt"
	"log"
	"time"

	"github.com/iron-io/iron_go3/config"
	"github.com/iron-io/iron_go3/mq"
)

func main() {
	settings := &config.Settings{
		Token:     "x",
		ProjectId: "y",
		Host:      "abc.iron.io",
	}
	q := mq.ConfigNew("sampleQueue", settings)

	FlatLoad(&q, 1000)
}

func FlatLoad(q *mq.Queue, msgsPerSecond int) {
	msgs := make([]string, 0, msgsPerSecond)
	count := 0
	for {
		go func() {
			for x := 0; x < msgsPerSecond; x++ {
				msgs = append(msgs, fmt.Sprintf("%d", count+x))
			}
			s := time.Now()
			_, err := q.PushStrings(msgs...)
			d := time.Since(s)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(d)
			msgs = msgs[:0]
		}()
		count += msgsPerSecond
		time.Sleep(time.Second)
	}
}
