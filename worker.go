package main

import "sync"

func worker(id int, mappers []Mapper, in <-chan int, out chan<- Response, wg *sync.WaitGroup) {

	defer wg.Done()
	for n := range in {
		for _, mapper := range mappers {
			if mapper.Cond(n) {
				out <- Response{
					WorkerId: id,
					Arg:      n,
					Value:    mapper.Map(n),
					HasValue: true,
				}
			} else {
				out <- Response{
					WorkerId: id,
					Arg:      n,
					Value:    "",
					HasValue: false,
				}
			}
		}
	}
}
