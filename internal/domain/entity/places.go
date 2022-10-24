package entity
import(
	"context"
)
type Places struct {
	Id int `json:"id"`
	Latitude string `json:"latitude" validate:"required`
	Longitude string `json:"longitude" validate:"required"`
}

type PlacesUsecase interface {
	GetById(ctx context.Context, id int) (Places, error)
	UpdatePlaces(ctx context.Context, pl *Places) error
	DeletePlaces(ctx context.Context, id Places) error
	SaveNewPlaces(ctx context.Context, pl Places) error
}

type PlacesRepository interface {
	GetById(ctx context.Context, id int) (model Places, err error)
	UpdatePlaces(ctx context.Context, pl *Places) (err error)
	DeletePlaces(ctx context.Context, id Places) (err error)
	SaveNewPlaces(ctx context.Context, pl Places) (err error)
}




// --Long ID
// --String lat, lad (Координаты разрешенного места)
// --latitude and longitude


//Добавить структуру с кодами ошибок, которые будут возвращаться вместо дефолтных ошибок 