package main

import "testing"

func TestAddTask(t *testing.T) {
	var list List

	task := Task{Label: "task1"}

	list.AddTask(task)
	if list.Tasks[0].Label != "task1" {
		t.Error("Expected 'task1' but got ", list.Tasks[0].Label)
	}
}
