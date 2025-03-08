package handlers

import (
	"context"
	"log"

	"github.com/equaerdist/simple-api/internal/services"
	car "github.com/equaerdist/simple-api/pkg/github.com/equaerdist/simple-api/pkg/car"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CarGrpcService struct {
	car.CarServiceServer
}

func (s *CarGrpcService) GetCar(ctx context.Context, req *car.GetCarRequest) (*car.GetCarResponse, error) {
	serv := services.NewCarService()
	carRes, err := serv.GetCar(ctx, int(req.CarId))

	if err != nil {
		log.Printf("Error ocurred when try get car from service: %v", err)
		return nil, err
	}

	return &car.GetCarResponse{
		CarId:     int64(carRes.ID),
		ModelName: carRes.ModelName,
		CreatedAt: timestamppb.New(carRes.CreatedAt),
	}, nil
}

func (s *CarGrpcService) CreateCar(ctx context.Context, req *car.CreateCarRequest) (*car.CreateCarResponse, error) {
	serv := services.NewCarService()
	id, err := serv.CreateCar(ctx, req.ModelName)

	if err != nil {
		log.Printf("Error ocurred when try create car with car service: %v", err)
		return nil, err
	}

	return &car.CreateCarResponse{
		CarId: int64(id),
	}, nil
}

func (s *CarGrpcService) UpdateCar(ctx context.Context, req *car.UpdateCarRequest) (*car.UpdateCarResponse, error) {
	serv := services.NewCarService()
	err := serv.UpdateCar(ctx, int(req.CarId), req.ModelName)
	if err != nil {
		return nil, err
	}

	return &car.UpdateCarResponse{}, nil
}

func (s *CarGrpcService) DeleteCar(ctx context.Context, req *car.DeleteCarRequest) (*car.DeleteCarResponse, error) {
	serv := services.NewCarService()
	err := serv.DeleteCar(ctx, int(req.CarId))
	if err != nil {
		return nil, err
	}
	return &car.DeleteCarResponse{}, nil
}
