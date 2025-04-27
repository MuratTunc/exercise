package main

import (
	"fmt"
	"sync"
)

type Car struct {
	Brand string
	Model string
	Year  int
}

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Worker", id, "started")
}

type Person struct {
	Name string
}

func (p *Person) updateName(newName string) {
	p.Name = newName
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()
	fmt.Println("All Workers Done")

	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)

	for value := range ch {
		fmt.Println(value)
	}

	person := Person{Name: "mutu"}
	fmt.Println(person)

	person.updateName("bro")

	fmt.Println(person)

}
