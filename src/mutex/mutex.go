package main

import(
	"fmt"
	"sync"
)

type Incrementor struct {
	counter int
	mu sync.Mutex
}

func (inc *Incrementor) Inc(appendValue int) {
	inc.mu.Lock()
	inc.counter += appendValue;
	defer inc.mu.Unlock()
}

func (inc *Incrementor) Get() int {
	inc.mu.Lock()
	defer inc.mu.Unlock()
	return inc.counter
}


func main() {
	var inc = Incrementor{}

	for x:= 0; x< 100; x++ {
		go inc.Inc(1);
		go inc.Inc(2);
		go inc.Inc(20);
		fmt.Println(inc.Get())
	}
}
