package todo

type Task struct {
	ID        int    `json:"id"`
	Tarefa    string `json:"title"`
	Completed bool   `json:"completed"`
}
