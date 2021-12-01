package main

type List struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

func (l *List) AddTask(t Task) {
	l.Tasks = append(l.Tasks, t)
}

func (l *List) RemoveTask(t Task) {
	for i, task := range l.Tasks {
		if task.Label == t.Label {
			l.Tasks = append(l.Tasks[:i], l.Tasks[i+1:]...)
		}
	}
}

func (l *List) UpdateTask(t Task) {
	for i, task := range l.Tasks {
		if task.Label == t.Label {
			l.Tasks[i] = t
		}
	}
}
