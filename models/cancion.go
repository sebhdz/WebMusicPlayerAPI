package models

type Cancion struct {
	ID            int    `json:"id"`
	NombreAlbum   string `json:"nombre_album"`
	NombreArtista string `json:"nombre_artista"`
	NombreCanion  string `json:"nombre_canion"`
	Url           string `json:"url"`
}
