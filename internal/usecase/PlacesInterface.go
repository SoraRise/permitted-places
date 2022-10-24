package usecase

import (
	entity "github.com/SoraRise/permitted-places/internal/domain/entity"

	"context"
)
type PlacesUsecase interface {
	GetById(ctx context.Context, id int) (entity.Places, error)
	UpdatePlaces(ctx context.Context, pl *entity.Places) error
	DeletePlaces(ctx context.Context, id entity.Places) error
	SaveNewPlaces(ctx context.Context, pl entity.Places) error
}