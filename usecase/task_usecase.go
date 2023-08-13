package usecase

import (
	"gorestapi/model"
	"gorestapi/repository"
	"gorestapi/validator"
	"time"
)

type ITaskUsecase interface {
	CreateTask(task *model.Task) (*model.TaskResponse, error)
	GetTaskById(taskId, userId uint) (*model.TaskResponse, error)
	GetAllTasks(userId uint) ([]*model.TaskResponse, error)
	UpdateTask(task *model.Task) (*model.TaskResponse, error)
	DeleteTask(taskId, userId uint) error
}

type TaskUsecase struct {
	tr repository.ITaskRepotisoty
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepotisoty, tv validator.ITaskValidator) ITaskUsecase {
	return &TaskUsecase{
		tr: tr,
		tv: tv,
	}
}

func (tu *TaskUsecase) CreateTask(task *model.Task) (*model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return nil, err
	}

	task.CreatedAt = time.Now()

	if err := tu.tr.Create(task); err != nil {
		return nil, err
	}

	taskRes := &model.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskRes, nil
}

func (tu *TaskUsecase) GetTaskById(taskId, userId uint) (*model.TaskResponse, error) {
	task := &model.Task{}
	if err := tu.tr.FindById(task, taskId, userId); err != nil {
		return nil, err
	}

	taskRes := &model.TaskResponse{
		Id:        taskId,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskRes, nil
}

func (tu *TaskUsecase) GetAllTasks(userId uint) ([]*model.TaskResponse, error) {
	tasks := make([]*model.Task, 0)
	if err := tu.tr.FindAll(&tasks, userId); err != nil {
		return nil, err
	}

	taskResponses := make([]*model.TaskResponse, 0)
	for _, task := range tasks {
		taskResponses = append(taskResponses, &model.TaskResponse{
			Id:        task.Id,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}

	return taskResponses, nil
}

func (tu *TaskUsecase) UpdateTask(task *model.Task) (*model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return nil, err
	}

	if err := tu.tr.Update(task); err != nil {
		return nil, err
	}

	taskRes := &model.TaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskRes, nil
}

func (tu *TaskUsecase) DeleteTask(taskId, userId uint) error {
	return tu.tr.Delete(taskId, userId)
}
