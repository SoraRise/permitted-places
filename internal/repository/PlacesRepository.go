package repository

import(
	"context"
	entity "github.com/SoraRise/permitted-places/internal/domain/entity"
)

type PlacesRepository interface {
	GetById(ctx context.Context, id int) (model entity.Places, err error)
	UpdatePlaces(ctx context.Context, pl *entity.Places) (err error)
	DeletePlaces(ctx context.Context, id entity.Places) (err error)
	SaveNewPlaces(ctx context.Context, pl entity.Places) (err error)
}