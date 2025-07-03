package models

type Album struct {
	ID            int    `json:"id"`
	NombreArtista string `json:"nombre_artista"`
	Nombre        string `json:"nombre"`
	Fecha         string `json:"fecha"`
	Imagen        string `json:"imagen"`
}
