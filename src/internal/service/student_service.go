package service

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"goApp/src/internal/assets/factory/dto"
	"goApp/src/internal/assets/factory/mapper" 
	"goApp/src/internal/repository"
)

type StudentServiceImpl struct {
	repo repository.StudentRepository
}

type StudentService interface {
	Create(ctx context.Context, dto dto.StudentDTO) (*dto.StudentDTO, error)
	List(ctx context.Context) ([]dto.StudentDTO, error)
	Get(ctx context.Context, id string) (*dto.StudentDTO, error)
}

func (s *StudentServiceImpl) Create(ctx context.Context, d dto.StudentDTO) (*dto.StudentDTO, error) {
	start := time.Now()
	
	model, err := mapper.ToStudentModel(d)
	if err != nil {
		log.Errorf("Model conversion error: %v", err)
		return nil, err
	}

	res, err := s.repo.Insert(ctx, *model)
	if err != nil {
		return nil, err
	}

	log.Infof("Created in %v", time.Since(start))
	return mapper.ToStudentDTO(*res), nil
}

func (s *StudentServiceImpl) List(ctx context.Context) ([]dto.StudentDTO, error) {
	res, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToStudentDTOs(res), nil
}

func (s *StudentServiceImpl) Get(ctx context.Context, id string) (*dto.StudentDTO, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	
	res, err := s.repo.FindById(ctx, oid)
	if err != nil {
		return nil, err
	}
	return mapper.ToStudentDTO(*res), nil
}



func NewStudentService(repo repository.StudentRepository) StudentService {
	return &StudentServiceImpl{repo: repo}
}