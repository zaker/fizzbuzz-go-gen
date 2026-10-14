package main

import (
	"sort"
)

func SortedBuffer(in <-chan Response, out chan<- Response) {

	jobId := 1
	var responses []Response
	go func() {

		for n := range in {
			//fmt.Println("Processing", n)
			if n.HasValue {
				if len(responses) == 0 || n.Arg < responses[len(responses)-1].Arg {
					responses = append(responses, n)
				} else {
					responses = append(responses, n)
					sort.Slice(responses, func(i, j int) bool {
						return responses[i].Arg > responses[j].Arg
					})
				}

				//fmt.Println("Sending", jobId, responses[len(responses)-1].Arg)
				for len(responses) > 0 && responses[len(responses)-1].Arg <= jobId {
					// fmt.Println("Sending", responses[len(responses)-1])
					out <- responses[len(responses)-1]
					jobId++
					responses = responses[:len(responses)-1]
				}

			}
		}

		close(out)
	}()

}
