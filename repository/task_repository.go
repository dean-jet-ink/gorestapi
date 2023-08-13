package repository

import (
	"fmt"
	"gorestapi/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepotisoty interface {
	Create(task *model.Task) error
	FindById(task *model.Task, taskId uint, userId uint) error
	FindAll(tasks *[]*model.Task, userId uint) error
	Update(task *model.Task) error
	Delete(taskId uint, userId uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepotisoty {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) Create(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) FindById(task *model.Task, taskId, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) FindAll(tasks *[]*model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) Update(task *model.Task) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", task.Id, task.UserId).Update("title", task.Title)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not existed")
	}

	return nil
}

func (tr *TaskRepository) Delete(taskId, userId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not existed")
	}

	return nil
}
