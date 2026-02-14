package main

import (
	"fmt"
	"sync"
)

// type post struct {
// 	views int
// }

// func (p *post) inc() {
// 	p.views += 1
// }

// func main() {
// 	myPost := post{views: 0}
// 	myPost.inc()
// 	myPost.inc()
// 	fmt.Println(myPost.views)
// }

//
// type post struct {
// 	views int
// }

// func (p *post) inc() {
// 	p.views += 1
// }

// func main() {
// 	myPost := post{views: 0}
// 	for i := 0; i < 100; i++ {
// 		myPost.inc()
// 	}
// 	fmt.Println(myPost.views)
// }

//
// type post struct {
// 	views int
// }

// func (p *post) inc() {
// 	p.views += 1
// }

// func main() {
// 	myPost := post{views: 0}
// 	for i := 0; i < 100; i++ {
// 		go myPost.inc()
// 	}
// 	fmt.Println(myPost.views)
// }

//
// type post struct {
// 	views int
// }

// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	p.views += 1
// }
// func main() {
// 	var wg sync.WaitGroup
// 	myPost := post{views: 0}
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go myPost.inc(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(myPost.views)
// }

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
