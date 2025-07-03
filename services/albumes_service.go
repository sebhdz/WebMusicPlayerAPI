package services

import (
	"log"
	"pruebaAPI/database"
	"pruebaAPI/models"
)

func ObtenerAlbumes() ([]models.Album, error) {
	rows, err := database.DB.Query("SELECT albumes.id, albumes.nombre_album, artistas.nombre_artista, albumes.fecha_publicacion, albumes.url_imagen FROM albumes JOIN artistas ON albumes.id_artista = artistas.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var albumes []models.Album
	for rows.Next() {
		var album models.Album
		err := rows.Scan(&album.ID, &album.Nombre, &album.NombreArtista, &album.Fecha, &album.Imagen)
		if err != nil {
			log.Fatal("ERROR AL OBTENER ALBUMES")
			return nil, err
		}
		albumes = append(albumes, album)
	}
	return albumes, nil
}
