package handlers

import u "git.garena.com/aldino.rahman/go-redis-example/internal/usecases"

type Handler struct {
	studentUseCase u.IstudentUseCase
}

type Opts struct {
	StudentUseCase u.IstudentUseCase
}

func New(opts *Opts) *Handler {
	return &Handler{
		studentUseCase: opts.StudentUseCase,
	}
}