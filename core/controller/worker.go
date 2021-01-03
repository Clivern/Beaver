// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"sync"
	"time"
)

const WorkersCount = 10000

func main() {
	Work(Generate(1000000))
}

func Generate(max int) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			for i := 0; i < max; i++ {
				channel <- fmt.Sprintf("Record %d!\n", i)
			}

			break
		}

		close(channel)
	}()

	return channel
}

func Work(inputChannel <-chan string) {
	wg := &sync.WaitGroup{}

	for t := 0; t < WorkersCount; t++ {
		wg.Add(1)
		go Worker(inputChannel, wg)
	}

	wg.Wait()
}

func Worker(inputChannel <-chan string, wg *sync.WaitGroup) {
	for input := range inputChannel {
		fmt.Printf("Started Processing %s\n", input)

		time.Sleep(3 * time.Second)
	}

	wg.Done()
}
