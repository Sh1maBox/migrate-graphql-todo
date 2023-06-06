package gqltodo

type (
	TodoRepository interface {
		GetAll() []Todo
		Create(title, ID string) Todo
		GetByID(ID string) Todo
	}

	InmemoryTodoRepository struct {
		// "title": "ご飯作る
		db []map[string]string
	}
)

func (repo *InmemoryTodoRepository) GetAll() []Todo {
	todosData := repo.db
	var todos []Todo
	for _, v := range todosData {
		todos = append(todos, NewTodo(
			v["title"], v["id"],
		))
	}
	return todos
}

func (repo *InmemoryTodoRepository) Create(title, ID string) Todo {
	todo := NewTodo(
		title, ID,
	)
	repo.db = append(repo.db, map[string]string{
		"title": todo.Title,
		"id":    todo.ID,
	})
	return todo
}

func (repo *InmemoryTodoRepository) GetByID(ID string) Todo {
	todosData := repo.db
	var todo Todo
	for _, v := range todosData {
		if v["id"] == ID {
			todo = NewTodo(
				v["title"], v["id"],
			)
			break
		}
	}
	return todo
}

func NewInmemoryTodoRepository(todoMap []map[string]string) InmemoryTodoRepository {
	return InmemoryTodoRepository{
		db: todoMap,
	}
}
