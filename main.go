package main

import (
	"fmt"
	"sync"
)

func Gen(max int, ch chan<- int) {
	go func() {
		for i := range max {
			ch <- i + 1
		}
		close(ch)
	}()

}

func isNotFizzBuzz(n int) bool {
	return !(n%3 == 0 || n%5 == 0)
}

func isFizzBuzz(n int) bool {
	return n%3 == 0 && n%5 == 0
}

func number(n int) string {
	return fmt.Sprintf("%d", n)
}

func isFizz(n int) bool {
	return n%3 == 0 && n%5 != 0
}

func fizz(n int) string {
	return "Fizz"
}

func isBuzz(n int) bool {
	return n%5 == 0 && n%3 != 0
}

func buzz(n int) string {
	return "Buzz"
}

func fizzbuzz(n int) string {
	return "FizzBuzz"
}

type Mapper struct {
	Cond func(int) bool
	Map  func(int) string
}

type Response struct {
	WorkerId int
	Arg      int
	Value    string
	HasValue bool
}

func (r Response) String() string {
	return fmt.Sprintf("Response{WorkerId: %d, Arg: %d, Value: %s, HasValue: %v}", r.WorkerId, r.Arg, r.Value, r.HasValue)
}

func main() {
	wg := sync.WaitGroup{}
	const MaxNo = 10_000_000
	const numjobs = 1000

	var mappers []Mapper = []Mapper{
		{isNotFizzBuzz, number},
		{isFizzBuzz, fizzbuzz},
		{isFizz, fizz},
		{isBuzz, buzz},
	}

	genChan := make(chan int, 1000)
	valueChan := make(chan Response, 1000)
	sortedChan := make(chan Response, 1000)
	Gen(MaxNo, genChan)
	SortedBuffer(valueChan, sortedChan)
	for id := 0; id < numjobs; id++ {
		wg.Add(1)
		go worker(id, mappers, genChan, valueChan, &wg)
	}

	go func() {
		wg.Wait()
		close(valueChan)
	}()

	for s := range sortedChan {

		fmt.Println(s.Arg, "->", s.Value)
	}

}
