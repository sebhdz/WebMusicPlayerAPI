package services

import (
	"pruebaAPI/database"
	"pruebaAPI/models"
)

func ObtenerCanciones(idAlbum string) ([]models.Cancion, error) {
	rows, err := database.DB.Query("SELECT canciones.id, posicion, nombre_cancion, albumes.url_imagen, artistas.nombre_artista ,url FROM canciones JOIN albumes on canciones.id_album = albumes.id JOIN artistas on albumes.id_artista = artistas.id WHERE id_album = ?", idAlbum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var canciones []models.Cancion
	for rows.Next() {
		var cancion models.Cancion
		err := rows.Scan(&cancion.ID, &cancion.Posicion, &cancion.NombreCanion, &cancion.ImagenUrl, &cancion.Artista, &cancion.Url)
		if err != nil {
			return nil, err
		}
		canciones = append(canciones, cancion)
	}
	return canciones, nil
}
