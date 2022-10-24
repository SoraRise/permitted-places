package entity

type Places struct {
	Id int `json:"id"`
	Latitude string `json:"latitude" validate:"required`
	Longitude string `json:"longitude" validate:"required"`
}





// --Long ID
// --String lat, lad (Координаты разрешенного места)
// --latitude and longitude


//Добавить структуру с кодами ошибок, которые будут возвращаться вместо дефолтных ошибок 