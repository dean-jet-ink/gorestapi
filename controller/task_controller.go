package controller

import (
	"gorestapi/model"
	"gorestapi/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	CreateTask(c echo.Context) error
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type TaskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &TaskController{
		tu: tu,
	}
}

func (tc *TaskController) CreateTask(c echo.Context) error {
	userId := tc.userId(c)

	task := &model.Task{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task.UserId = userId

	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *TaskController) GetTaskById(c echo.Context) error {
	userId := tc.userId(c)

	taskIdStr := c.Param("taskId")
	taskId, _ := strconv.Atoi(taskIdStr)

	taskRes, err := tc.tu.GetTaskById(uint(taskId), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *TaskController) GetAllTasks(c echo.Context) error {
	userId := tc.userId(c)

	taskResponses, err := tc.tu.GetAllTasks(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskResponses)
}

func (tc *TaskController) UpdateTask(c echo.Context) error {
	userId := tc.userId(c)

	taskIdStr := c.Param("taskId")
	taskId, _ := strconv.Atoi(taskIdStr)

	task := &model.Task{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.Id = uint(taskId)
	task.UserId = userId

	taskRes, err := tc.tu.UpdateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *TaskController) DeleteTask(c echo.Context) error {
	userId := tc.userId(c)

	taskIdStr := c.Param("taskId")
	taskId, _ := strconv.Atoi(taskIdStr)

	if err := tc.tu.DeleteTask(uint(taskId), userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (tc *TaskController) userId(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	return uint(userId.(float64))
}
