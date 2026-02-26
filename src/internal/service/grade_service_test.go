package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"goApp/src/internal/assets/domain/model"
	"goApp/src/internal/assets/factory/dto"
)

type MockGradeRepo struct {
	mock.Mock
}

func (m *MockGradeRepo) Create(ctx context.Context, grade *model.GradeModel) error {
	args := m.Called(ctx, grade)
	return args.Error(0)
}

func (m *MockGradeRepo) GetByStudentID(ctx context.Context, studentID primitive.ObjectID) ([]model.GradeModel, error) {
	args := m.Called(ctx, studentID)
	return args.Get(0).([]model.GradeModel), args.Error(1)
}

func (m *MockGradeRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockStudentRepoForGrade struct {
	mock.Mock
}

func (m *MockStudentRepoForGrade) Insert(ctx context.Context, student model.StudentModel) (*model.StudentModel, error) {
	args := m.Called(ctx, student)
	if args.Get(0) != nil {
		return args.Get(0).(*model.StudentModel), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStudentRepoForGrade) GetAll(ctx context.Context) ([]model.StudentModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.StudentModel), args.Error(1)
}

func (m *MockStudentRepoForGrade) FindById(ctx context.Context, ID primitive.ObjectID) (*model.StudentModel, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) != nil {
		return args.Get(0).(*model.StudentModel), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStudentRepoForGrade) AddGrade(ctx context.Context, ID primitive.ObjectID, subject string, grade float64) error {
	args := m.Called(ctx, ID, subject, grade)
	return args.Error(0)
}

func TestGradeService_AddGrade(t *testing.T) {
	ctx := context.Background()

	t.Run("mapper error", func(t *testing.T) {
		gradeRepo := new(MockGradeRepo)
		studentRepo := new(MockStudentRepoForGrade)
		svc := NewGradeService(gradeRepo, studentRepo)

		dtoInput := dto.GradeDTO{
			ID:        "",
			StudentID: "not-a-valid-hex",
			Subject:   "Go",
			Value:     10,
		}

		res, err := svc.AddGrade(ctx, dtoInput)

		assert.Error(t, err)
		assert.Nil(t, res)
		gradeRepo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
		studentRepo.AssertNotCalled(t, "FindById", mock.Anything, mock.Anything)
	})

	t.Run("student not found", func(t *testing.T) {
		gradeRepo := new(MockGradeRepo)
		studentRepo := new(MockStudentRepoForGrade)
		svc := NewGradeService(gradeRepo, studentRepo)

		studentID := primitive.NewObjectID()
		dtoInput := dto.GradeDTO{
			ID:        "",
			StudentID: studentID.Hex(),
			Subject:   "Go",
			Value:     8,
		}

		studentRepo.On("FindById", ctx, studentID).Return(nil, errors.New("student not found"))

		res, err := svc.AddGrade(ctx, dtoInput)

		assert.Error(t, err)
		assert.Nil(t, res)
		gradeRepo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
		studentRepo.AssertExpectations(t)
	})

	t.Run("create grade error", func(t *testing.T) {
		gradeRepo := new(MockGradeRepo)
		studentRepo := new(MockStudentRepoForGrade)
		svc := NewGradeService(gradeRepo, studentRepo)

		studentID := primitive.NewObjectID()
		dtoInput := dto.GradeDTO{
			ID:        "",
			StudentID: studentID.Hex(),
			Subject:   "Go",
			Value:     7,
		}

		student := &model.StudentModel{ID: studentID}
		studentRepo.On("FindById", ctx, studentID).Return(student, nil)
		gradeRepo.On("Create", ctx, mock.AnythingOfType("*model.GradeModel")).Return(errors.New("create error"))

		res, err := svc.AddGrade(ctx, dtoInput)

		assert.Error(t, err)
		assert.Nil(t, res)
		studentRepo.AssertExpectations(t)
		gradeRepo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		gradeRepo := new(MockGradeRepo)
		studentRepo := new(MockStudentRepoForGrade)
		svc := NewGradeService(gradeRepo, studentRepo)

		studentID := primitive.NewObjectID()
		dtoInput := dto.GradeDTO{
			ID:        "",
			StudentID: studentID.Hex(),
			Subject:   "Go",
			Value:     9,
		}

		student := &model.StudentModel{ID: studentID}
		studentRepo.On("FindById", ctx, studentID).Return(student, nil)
		gradeRepo.On("Create", ctx, mock.AnythingOfType("*model.GradeModel")).Return(nil)

		res, err := svc.AddGrade(ctx, dtoInput)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, dtoInput.Subject, res.Subject)
		assert.Equal(t, dtoInput.Value, res.Value)
		studentRepo.AssertExpectations(t)
		gradeRepo.AssertExpectations(t)
	})
}

