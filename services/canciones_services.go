package services

import (
	"fmt"
	"pruebaAPI/database"
	"pruebaAPI/models"
)

func ObtenerCanciones(id_album string) ([]models.Cancion, error) {
	rows, err := database.DB.Query("SELECT canciones.id, albumes.nombre_album, artistas.nombre_artista, nombre_cancion, url FROM canciones JOIN albumes ON id_album = albumes.id JOIN artistas ON albumes.id_artista = artistas.id WHERE id_album = ? ORDER BY canciones.posicion", id_album)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var canciones []models.Cancion
	for rows.Next() {
		var cancion models.Cancion
		err := rows.Scan(&cancion.ID, &cancion.NombreAlbum, &cancion.NombreArtista, &cancion.NombreCanion, &cancion.Url)
		if err != nil {
			return nil, err
		}
		canciones = append(canciones, cancion)
	}
	fmt.Println(canciones)
	return canciones, nil
}
