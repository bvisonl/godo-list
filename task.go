package main

type Task struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Date        string `json:"date"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"`
	Color       string `json:"color"`
}
