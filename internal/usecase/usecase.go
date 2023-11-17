package usecase

import (
	"context"
	"grpcTest/internal/entity"
	"log/slog"
)

type UseCase struct {
	Itn
}

type Itn interface {
	GetItnInfo(ctx context.Context, dto *entity.DTO) error
}

func New(logger *slog.Logger) *UseCase {
	return &UseCase{
		Itn: NewItnUseCase(logger),
	}
}
