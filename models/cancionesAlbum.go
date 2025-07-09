package models

type CancionesAlbum struct {
	NombreDisco string    `json:"nombre_disco"`
	Imagen      string    `json:"imagen"`
	Artista     string    `json:"artista"`
	Fecha       string    `json:"fecha"`
	Canciones   []Cancion `json:"canciones"`
}
