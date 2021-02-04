package main

import (
	"fmt"

	"github.com/joshuaswickirl/interface-example/myqueue" //import implementation
)

type Queue interface {
	AddJob(string, []string) string
}

func main() {
	var q Queue
	q = &myqueue.Queue{} //use myqueue implementation imported above
	name := "Eli"
	groups := []string{"Group1", "Group2"}
	_ = processRequest(name, groups, q)

}

func processRequest(name string, groups []string, q Queue) string {
	jobID := q.AddJob(name, groups)
	fmt.Println("Job ID:", jobID)
	return jobID
}
