package middlewaresUsecases

import (
	"github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresRepositories"
)

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepositories.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepositories.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}


