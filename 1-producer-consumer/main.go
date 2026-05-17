//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, out chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			out <- nil
			return
		}
		out <- tweet
	}
}

func consumer(in chan *Tweet) {
	for t := range in {
		if t == nil {
			return
		}
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	wg := sync.WaitGroup{}
	pipe := make(chan *Tweet)

	wg.Go(func() {
		producer(stream, pipe)
	})
	wg.Go(func() {
		consumer(pipe)
	})

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
