package gqltodo

type (
	GetTodosUseCase struct {
		repo TodoRepository
	}

	CreateTodoUsecase struct {
		repo TodoRepository
	}

	GetTodoUseCase struct {
		repo TodoRepository
	}
)

func (r *GetTodosUseCase) GetTodos() []Todo {
	return r.repo.GetAll()
}

func NewGetTodosUseCase(repo TodoRepository) GetTodosUseCase {
	return GetTodosUseCase{repo}
}

func (r *CreateTodoUsecase) CreateTodo(title string) Todo {
	uuid := MakeUUID([]byte(title))
	return r.repo.Create(title, uuid)
}

func NewCreateTodoUsecase(repo TodoRepository) CreateTodoUsecase {
	return CreateTodoUsecase{repo}
}

func (r *GetTodoUseCase) GetTodoByID(id string) Todo {
	return r.repo.GetByID(id)
}

func NewGetTodoByIDUseCase(repo TodoRepository) GetTodoUseCase {
	return GetTodoUseCase{repo}
}
