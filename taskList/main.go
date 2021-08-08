package main

type IdTraker struct {
	id int
}

func (it *IdTraker) Get() int {
	it.id++
	return it.id
}

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskList struct {
	tasks map[int]*Task
	id    *IdTraker
}

func (c *TaskList) Create(task *Task) {
	task.ID = c.id.Get()
	c.tasks[task.ID] = task
}

func (c *TaskList) Update(task *Task) {
	c.tasks[task.ID] = task
}

func (c *TaskList) Get(id int) *Task {
	return c.tasks[id]
}

func (c *TaskList) GetAll() map[int]*Task {
	return c.tasks
}

func (c *TaskList) Delete(id int) {
	delete(c.tasks, id)
}

func main() {

}
