package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"goApp/src/internal/assets/domain/model"
)

type StudentRepository interface {
	Insert(ctx context.Context, student model.StudentModel) (*model.StudentModel, error)
	GetAll(ctx context.Context) ([]model.StudentModel, error)
	FindById(ctx context.Context, ID primitive.ObjectID) (*model.StudentModel, error)
	AddGrade(ctx context.Context, ID primitive.ObjectID, subject string, grade float64) error
}

type StudentRepositoryImpl struct {
	collection *mongo.Collection
}

func (r *StudentRepositoryImpl) Insert(ctx context.Context, student model.StudentModel) (*model.StudentModel, error) {
	result, err := r.collection.InsertOne(ctx, student)
	if err != nil {
		return nil, fmt.Errorf("Failed Insert: %w", err)
	}
	student.ID = result.InsertedID.(primitive.ObjectID)
	return &student, nil
}

func (r *StudentRepositoryImpl) FindById(ctx context.Context, ID primitive.ObjectID) (*model.StudentModel, error) {
	var student model.StudentModel
	err := r.collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&student)
	if err != nil {
		return nil, fmt.Errorf("Failed search by ID: %w", err)
	}
	return &student, nil
}

func (r *StudentRepositoryImpl) GetAll(ctx context.Context) ([]model.StudentModel, error) {
	list, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("Failed Get List: %w", err)
	}
	defer list.Close(ctx)

	var students []model.StudentModel
	if err = list.All(ctx, &students); err != nil {
		return nil, fmt.Errorf("Invalid model data: %w", err)
	}
	return students, nil
}

func (r *StudentRepositoryImpl) AddGrade(ctx context.Context, ID primitive.ObjectID, subject string, grade float64) error {
	filter := bson.M{"_id": ID}
	update := bson.M{
		"$push": bson.M{
			"voti": bson.M{
				"materia": subject,
				"voto":    grade,
			},
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Failed Update: %w", err)
	}
	return nil
}

func NewStudentRepository(db *mongo.Database) StudentRepository {
	return &StudentRepositoryImpl{
		collection: db.Collection("students"),
	}
}
