package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"goApp/src/internal/assets/factory/dto"
	"goApp/src/internal/assets/factory/mapper"
	"goApp/src/internal/repository"
)

type GradeService interface {
	AddGrade(ctx context.Context, gradeDTO dto.GradeDTO) (*dto.GradeDTO, error)
	ListBySubject(ctx context.Context, studentID string, subject string) ([]dto.GradeDTO, error)
	AverageBySubject(ctx context.Context, studentID string, subject string) (float64, error)
}

type gradeService struct {
	repo        repository.GradeRepository
	studentRepo repository.StudentRepository
}

func NewGradeService(r repository.GradeRepository, sr repository.StudentRepository) GradeService {
	return &gradeService{repo: r, studentRepo: sr}
}

func (s *gradeService) AddGrade(ctx context.Context, gradeDTO dto.GradeDTO) (*dto.GradeDTO, error) {
	m, err := mapper.ToGradeModel(gradeDTO)
	if err != nil {
		return nil, err
	}

	_, err = s.studentRepo.FindById(ctx, m.StudentID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	return mapper.ToGradeDTO(*m), nil
}

func (s *gradeService) ListBySubject(ctx context.Context, studentID string, subject string) ([]dto.GradeDTO, error) {
	oid, err := primitive.ObjectIDFromHex(studentID)
	if err != nil {
		return nil, err
	}

	grades, err := s.repo.GetByStudentID(ctx, oid)
	if err != nil {
		return nil, err
	}

	var filtered []dto.GradeDTO
	for _, g := range grades {
		if g.Subject == subject {
			d := mapper.ToGradeDTO(g)
			filtered = append(filtered, *d)
		}
	}

	return filtered, nil
}

func (s *gradeService) AverageBySubject(ctx context.Context, studentID string, subject string) (float64, error) {
	grades, err := s.ListBySubject(ctx, studentID, subject)
	if err != nil {
		return 0, err
	}
	if len(grades) == 0 {
		return 0, nil
	}

	var sum float64
	for _, g := range grades {
		sum += g.Value
	}

	return sum / float64(len(grades)), nil
}
