package gqltodo

type Todo struct {
	Title string
	ID    string
}

func NewTodo(title string, ID string) Todo {
	return Todo{Title: title, ID: ID}
}
