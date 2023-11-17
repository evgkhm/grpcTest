package grpc

import (
	"context"
	"grpcTest/internal/controller/grpc/spec"
)

func (s *Service) GetItnInfo(ctx context.Context, req *spec.GetItnRequest) (*spec.GetItnResponse, error) {
	inputItn := req.Itn
	itnDTO, err := s.useCase.GetItnInfo(ctx, inputItn)
	if err != nil {
		return &spec.GetItnResponse{Success: false}, err
	}

	return &spec.GetItnResponse{
		Itn:         itnDTO.Itn,
		Kpp:         itnDTO.Kpp,
		NameCompany: itnDTO.NameCompany,
		NameLeader:  itnDTO.NameLeader,
		Success:     true,
	}, nil
}
