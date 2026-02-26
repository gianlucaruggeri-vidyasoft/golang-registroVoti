package dto

import "time"

type GradeDTO struct {
	ID        string    `bson:"_id,omitempty"`
	StudentID string    `bson:"student_id"`
	Subject   string    `bson:"subject"`
	Value     float64   `bson:"value"`
	Date      time.Time `bson:"date"`
}
