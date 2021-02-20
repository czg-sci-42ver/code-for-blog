package task

import (
	"net/http"
	"time"

	"example.com/internal/taskstore"
	"github.com/labstack/echo/v4"
)

type TaskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *TaskServer {
	store := taskstore.New()
	return &TaskServer{store: store}
}

func (ts *TaskServer) GetDueYearMonthDay(ctx echo.Context, year int, month int, day int) error {
	tasks := ts.store.GetTasksByDueDate(year, time.Month(month), day)
	return ctx.JSON(http.StatusOK, tasks)
}

func (ts *TaskServer) GetTagTagname(ctx echo.Context, tagname string) error {
	tasks := ts.store.GetTasksByTag(tagname)
	return ctx.JSON(http.StatusOK, tasks)
}

func (ts *TaskServer) GetTask(ctx echo.Context) error {
	allTasks := ts.store.GetAllTasks()
	return ctx.JSON(http.StatusOK, allTasks)
}

func (ts *TaskServer) PostTask(ctx echo.Context) error {
	var taskBody PostTaskJSONBody
	err := ctx.Bind(&taskBody)
	if err != nil {
		return err
	}

	// TODO: check non-nil on these fields?!
	// TODO: do I need additional error checking here?
	id := ts.store.CreateTask(*taskBody.Text, *taskBody.Tags, *taskBody.Due)
	type ResponseId struct {
		Id int `json:"id"`
	}
	return ctx.JSON(http.StatusOK, ResponseId{Id: id})
}

func (ts *TaskServer) DeleteTaskId(ctx echo.Context, id int) error {
	return ts.store.DeleteTask(id)
}

func (ts *TaskServer) DeleteAllTasks(ctx echo.Context) error {
	ts.store.DeleteAllTasks()
	return nil
}

func (ts *TaskServer) GetTaskId(ctx echo.Context, id int) error {
	task, err := ts.store.GetTask(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, task)
}
