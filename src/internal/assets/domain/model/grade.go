package model
import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"

type GradeModel struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    StudentID primitive.ObjectID `bson:"student_id"` 
    Subject   string             `bson:"subject"`
    Value     float64            `bson:"value"`
    Date      time.Time          `bson:"date"`
}