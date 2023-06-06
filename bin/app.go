package main

import (
	"fmt"
	gqltodo "gql-todo"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	var todoMap []map[string]string
	for i := 0; i < 100; i++ {
		todoTitle := fmt.Sprintf("todo %d", i)
		todoMap = append(todoMap, map[string]string{"title": todoTitle, "id": gqltodo.MakeUUID([]byte(todoTitle))})
	}

	todoRepo := gqltodo.NewInmemoryTodoRepository(todoMap)

	e.GET("/todos", NewGetTodosController(gqltodo.NewGetTodosUseCase(&todoRepo)).handle)
	e.GET("/todo/:id", NewGetTodoController(gqltodo.NewGetTodoByIDUseCase(&todoRepo)).handle)
	e.POST("/todo", NewCreateTodoController(gqltodo.NewCreateTodoUsecase(&todoRepo)).handle)

	e.Logger.Fatal(e.Start(":1323"))
}

type (
	GetTodosResponse struct {
		Todos      []gqltodo.Todo `json:"todos"`
		TotalCount int64          `json:"totalCount"`
	}
	GetTodosController struct {
		usecase gqltodo.GetTodosUseCase
	}

	GetTodoRequest struct {
		ID string `param:"id"`
	}

	GetTodoResponse struct {
		Todo gqltodo.Todo `json:"todo"`
	}

	GetTodoController struct {
		usecase gqltodo.GetTodoUseCase
	}

	CreateTodoRequest struct {
		Title string `json:"title"`
	}

	CreateTodoResponse struct {
		Todo gqltodo.Todo `json:"todo"`
	}

	CreateTodoController struct {
		usecase gqltodo.CreateTodoUsecase
	}
)

func (r GetTodosController) handle(c echo.Context) error {
	todos := r.usecase.GetTodos()

	resp := NewGetTodosResponse(todos, int64(len(todos)))

	return c.JSON(http.StatusOK, resp)
}

func NewGetTodosResponse(todos []gqltodo.Todo, totalCount int64) GetTodosResponse {
	return GetTodosResponse{todos, totalCount}
}

func NewGetTodosController(usecase gqltodo.GetTodosUseCase) GetTodosController {
	return GetTodosController{usecase}
}

func (r CreateTodoController) handle(c echo.Context) error {
	req := new(CreateTodoRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
	}
	todo := r.usecase.CreateTodo(req.Title)

	resp := NewCreateTodoResponse(todo)

	return c.JSON(http.StatusCreated, resp)
}

func (r GetTodoController) handle(c echo.Context) error {
	req := new(GetTodoRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
	}
	todo := r.usecase.GetTodoByID(req.ID)

	resp := NewGetTodoResponse(todo)

	return c.JSON(http.StatusOK, resp)
}

func NewGetTodoResponse(todo gqltodo.Todo) GetTodoResponse {
	return GetTodoResponse{todo}
}

func NewGetTodoController(usecase gqltodo.GetTodoUseCase) GetTodoController {
	return GetTodoController{usecase}
}

func NewCreateTodoResponse(todo gqltodo.Todo) CreateTodoResponse {
	return CreateTodoResponse{todo}
}

func NewCreateTodoController(usecase gqltodo.CreateTodoUsecase) CreateTodoController {
	return CreateTodoController{usecase}
}
