package main

import (
	"algorithm/queue"
	"fmt"
)

func main() {
	s := queue.LinkQueue{}
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue("caonima")
	fmt.Println(s.Size())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Dequeue())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
	fmt.Println(s.Dequeue())
	fmt.Println(s.IsEmpty())
}
