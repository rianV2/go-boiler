package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/remnv/go-boiler/internal/model"
)

type PlayerCreateRequest struct {
	Name  *string `json:"name"`
	Level *string `json:"level"`
	Job   *string `json:"job"`
}

func (p PlayerCreateRequest) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Level, validation.Required),
		validation.Field(&p.Job, validation.Required),
	)
}

func (p PlayerCreateRequest) ToModel() (*model.Player, error) {
	return &model.Player{
		Name:  p.Name,
		Level: p.Level,
		Job:   p.Job,
	}, nil
}
