package models

type Task struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (t *Task) Update(newTask Task) bool {

	if t.Id != newTask.Id {
		return false
	}

	t.Title = newTask.Title
	t.Author = newTask.Author
	return true
}
