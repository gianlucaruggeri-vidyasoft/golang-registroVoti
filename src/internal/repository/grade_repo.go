package repository

import (
	"context"
	"goApp/src/internal/assets/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GradeRepository interface {
	Create(ctx context.Context, grade *model.GradeModel) error
	GetByStudentID(ctx context.Context, studentID primitive.ObjectID) ([]model.GradeModel, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type mongoGradeRepo struct {
	collection *mongo.Collection
}

func NewGradeRepository(db *mongo.Database) GradeRepository {
	return &mongoGradeRepo{
		collection: db.Collection("grades"),
	}
}

func (r *mongoGradeRepo) Create(ctx context.Context, grade *model.GradeModel) error {
	_, err := r.collection.InsertOne(ctx, grade)
	return err
}

func (r *mongoGradeRepo) GetByStudentID(ctx context.Context, studentID primitive.ObjectID) ([]model.GradeModel, error) {
	var grades []model.GradeModel
	
	cursor, err := r.collection.Find(ctx, bson.M{"student_id": studentID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}
	return grades, nil
}

func (r *mongoGradeRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}