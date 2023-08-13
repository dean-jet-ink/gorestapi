package main

import (
	"gorestapi/controller"
	"gorestapi/db"
	"gorestapi/engine"
	"gorestapi/repository"
	"gorestapi/usecase"
	"gorestapi/validator"
)

func main() {
	db := db.NewDB()
	ur := repository.NewUserRepository(db)
	uv := validator.NewUserValidator()
	uu := usecase.NewUserUsecase(ur, uv)
	uc := controller.NewUserController(uu)

	tr := repository.NewTaskRepository(db)
	tv := validator.NewTaskValidator()
	tu := usecase.NewTaskUsecase(tr, tv)
	tc := controller.NewTaskController(tu)

	e := engine.New(uc, tc)
	e.Logger.Fatal(e.Start(":8080"))
}
