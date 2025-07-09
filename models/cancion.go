package models

type Cancion struct {
	ID           int    `json:"id"`
	Posicion     int    `json:"posicion"`
	NombreCanion string `json:"nombre_cancion"`
	ImagenUrl    string `json:"imagen_cancion"`
	Artista      string `json:"artista"`
	Url          string `json:"url"`
}
