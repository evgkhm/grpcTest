package grpc

import (
	"context"
	"grpcTest/internal/controller/grpc/spec"
	"grpcTest/internal/entity"
)

func (s *Service) GetItnInfo(ctx context.Context, req *spec.GetItnRequest) (*spec.GetItnResponse, error) {
	itnDTO := entity.DTO{Itn: req.Itn}

	err := s.useCase.GetItnInfo(ctx, &itnDTO)
	if err != nil {
		return &spec.GetItnResponse{Success: false}, err
	}

	return &spec.GetItnResponse{
		Itn:         itnDTO.Itn,
		Kpp:         itnDTO.Kpp,
		NameCompany: itnDTO.NameCompany,
		NameLeader:  itnDTO.NameLeader,
		Success:     true}, nil
}
