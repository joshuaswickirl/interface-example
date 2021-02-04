package myqueue

import "fmt"

type Queue struct{}

func (q *Queue) AddJob(name string, groups []string) string {
	fmt.Println("Added job to my queue.")
	return "someJobIDFromMyQueue"
}
