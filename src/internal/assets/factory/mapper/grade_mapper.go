package mapper

import (
	"goApp/src/internal/assets/domain/model"
	"goApp/src/internal/assets/factory/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToGradeDTO(m model.GradeModel) *dto.GradeDTO {
	return &dto.GradeDTO{
		ID:        m.ID.Hex(),
		StudentID: m.StudentID.Hex(),
		Subject:   m.Subject,
		Value:     m.Value,
		Date:      m.Date,
	}
}

func ToGradeModel(d dto.GradeDTO) (*model.GradeModel, error) {
	var id primitive.ObjectID
	if d.ID != "" {
		var err error
		id, err = primitive.ObjectIDFromHex(d.ID)
		if err != nil {
			return nil, err
		}
	}

	studentID, err := primitive.ObjectIDFromHex(d.StudentID)
	if err != nil {
		return nil, err
	}

	return &model.GradeModel{
		ID:        id,
		StudentID: studentID,
		Subject:   d.Subject,
		Value:     d.Value,
		Date:      d.Date,
	}, nil
}
