package main

type List struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Date        string `json:"date"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"`
	Color       string `json:"color"`
}

func (l *List) AddTask(t Task) {
	l.Tasks = append(l.Tasks, t)
}
