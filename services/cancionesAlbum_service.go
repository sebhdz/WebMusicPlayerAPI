package services

import (
	"log"
	"pruebaAPI/database"
	"pruebaAPI/models"
)

func ObtenerCancionesAlbum(idAlbum string, canciones []models.Cancion) (models.CancionesAlbum, error) {
	row := database.DB.QueryRow(`SELECT albumes.nombre_album, artistas.nombre_artista, albumes.fecha_publicacion, albumes.url_imagen FROM albumes JOIN artistas on albumes.id_artista = artistas.id WHERE albumes.id = ?`, idAlbum)

	var resultado models.CancionesAlbum
	err := row.Scan(&resultado.NombreDisco, &resultado.Artista, &resultado.Fecha, &resultado.Imagen)
	if err != nil {
		log.Println(err)
		return models.CancionesAlbum{}, err
	}
	resultado.Canciones = canciones
	return resultado, nil

}
