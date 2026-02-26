package mapper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goApp/src/internal/assets/domain/model"
	"goApp/src/internal/assets/factory/dto"
)

func ToStudentDTO(m model.StudentModel) *dto.StudentDTO {
	

	return &dto.StudentDTO{
		ID:      m.ID.Hex(),
		Name:    m.Name,
		Surname: m.Surname,
	}
}

func ToStudentModel(d dto.StudentDTO) (*model.StudentModel, error) {
	//  converto l'id solo se c'è altrimenti ci pensa mongo
	var id primitive.ObjectID
	if d.ID != "" {
		var err error
		id, err = primitive.ObjectIDFromHex(d.ID)
		if err != nil {
			return nil, err
		}
	}

	return &model.StudentModel{
		ID:      id,
		Name:    d.Name,
		Surname: d.Surname,
	}, nil
}

func ToStudentDTOs(models []model.StudentModel) []dto.StudentDTO {
	dtos := make([]dto.StudentDTO, 0, len(models))
	for _, m := range models {
		dtos = append(dtos, *ToStudentDTO(m))
	}
	return dtos
}