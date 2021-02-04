package main

import (
	"fmt"
	"testing"
)

type testQueue struct{}

func (tq *testQueue) AddJob(name string, groups []string) string {
	fmt.Println("Add job to test queue.")
	return "someJobIDFromTestQueue"
}

func TestProcessRequest(t *testing.T) {
	q := &testQueue{} //use test implementation defined above
	name := "Jeff"
	groups := []string{"Group1", "Group3"}
	id := processRequest(name, groups, q)
	if id != "someJobIDFromTestQueue" {
		t.Errorf("Didn't get expected test ID.")
	}
}
