package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"goApp/src/internal/assets/domain/model"
)

type MockStudentRepo struct {
	mock.Mock
}

func (m *MockStudentRepo) Insert(ctx context.Context, student model.StudentModel) (*model.StudentModel, error) {
	args := m.Called(ctx, student)
	if args.Get(0) != nil {
		return args.Get(0).(*model.StudentModel), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStudentRepo) GetAll(ctx context.Context) ([]model.StudentModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.StudentModel), args.Error(1)
}

func (m *MockStudentRepo) FindById(ctx context.Context, ID primitive.ObjectID) (*model.StudentModel, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) != nil {
		return args.Get(0).(*model.StudentModel), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStudentRepo) AddGrade(ctx context.Context, ID primitive.ObjectID, subject string, grade float64) error {
	args := m.Called(ctx, ID, subject, grade)
	return args.Error(0)
}

func TestStudentService_Get(t *testing.T) {
	ctx := context.Background()
	validID := primitive.NewObjectID()

	t.Run("id non valido", func(t *testing.T) {
		mockRepo := new(MockStudentRepo)
		svc := NewStudentService(mockRepo)

		res, err := svc.Get(ctx, "id-non-esadecimale")

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("studente non trovato", func(t *testing.T) {
		mockRepo := new(MockStudentRepo)
		svc := NewStudentService(mockRepo)

		mockRepo.On("FindById", ctx, validID).Return(nil, errors.New("studente non trovato"))

		res, err := svc.Get(ctx, validID.Hex())

		assert.Error(t, err)
		assert.Equal(t, "studente non trovato", err.Error())
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("studente valido", func(t *testing.T) {
		mockRepo := new(MockStudentRepo)
		svc := NewStudentService(mockRepo)

		studenteFinto := &model.StudentModel{
			ID:      validID,
			Name:    "Mario",
			Surname: "Rossi",
		}

		mockRepo.On("FindById", ctx, validID).Return(studenteFinto, nil)

		res, err := svc.Get(ctx, validID.Hex())

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "Mario", res.Name)
		mockRepo.AssertExpectations(t)
	})
}

