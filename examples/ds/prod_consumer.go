package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type page struct {
	associatedPages []page
	content string
}

func main() {
	runtime.GOMAXPROCS(4)
	lastPage := page{
		content: "10th level",
	}
	anotherPage := page{
		content: "another 10th level",
	}
	for i := 9; i >= 1; i-- {
		p1 := page{
			associatedPages: []page{lastPage, anotherPage},
			content: fmt.Sprintf("%vth level", i),
		}
		p2 := page{
			associatedPages: []page{lastPage},
			content: fmt.Sprintf("another %vth level", i),
		}
		lastPage = p1
		anotherPage = p2
	}
	root := page{
		associatedPages: []page{lastPage, anotherPage},
		content: "root page",
	}
	// we should crawl only to this level
	level := 4
	printPages(root, level)
	fmt.Println("---------compare------------")
	printPagesParallel(root, level)
}

func printPages(root page, level int) {
	i := 0
	queue := []pageTup{{p: root, l: 0}}
	for len(queue) > 0 {
		currPageTup := queue[0]
		queue = queue[1:]
		fmt.Printf("content: %v \n", currPageTup.p.content)
		i = currPageTup.l + 1
		if i < level {
			for _, page := range currPageTup.p.associatedPages {
				queue = append(queue, pageTup{p: page, l: i})
			}
		}
	}
}

type pageTup struct {
	p page
	l int
}

type driver struct {
	pageChan chan *pageTup
	loaderGroup sync.WaitGroup
	channelStopGroup sync.WaitGroup
}

func printPagesParallel(root page, level int) {
	/*
	1. the number of worker threads (goroutines) are fixed
	2. each worker thread waits on the channel
	when something is available, it takes the page off the channel and
	prints the content. It then takes the associatedPages and puts
	it on the channel
	3. In order to exit the worker, add a waitgroup that closes when the worker sees
	that the page has no more associated pages (this can then be modified to check for the level)
	 */
	numWorkers := 5
	d := driver{pageChan: make(chan *pageTup, numWorkers)}
	d.pageChan <- &pageTup{p: root, l: 0}
	for i := 0; i < numWorkers; i++ {
		d.add(i + 1, level)
	}
	d.loaderGroup.Wait()
	fmt.Println("main is scheduled to close channel")
	close(d.pageChan)
	d.channelStopGroup.Wait()
	fmt.Printf("in main active goroutines: %v \n", runtime.NumGoroutine())
}

func (d *driver) add(workerNum int, level int) {
	d.loaderGroup.Add(1)
	d.channelStopGroup.Add(1)
	go func() {
		for true {
			currPageTup, ok := <- d.pageChan
			// this will make sure that these goroutines get cleared up when scheduled
			if !ok {
				fmt.Printf("nothing more to be done for worker: %v \n", workerNum)
				d.channelStopGroup.Done()
				return
			}
			// process
			fmt.Printf("content: %v \n", currPageTup.p.content)
			var sleepTime time.Duration
			for i := 0; i < workerNum; i++ {
				sleepTime += time.Second
			}
			time.Sleep(sleepTime)
			if currPageTup.l + 1 < level {
				for i := range currPageTup.p.associatedPages {
					d.pageChan <- &pageTup{p: currPageTup.p.associatedPages[i], l: currPageTup.l + 1}
				}
			} else {
				// decrement the waitgroup as this particular worker will have nothing more to add
				d.loaderGroup.Done()
			}
		}
	}()
}