package main

type Album struct {
	Id     int64
	Title  string
	Artist string
	Price  float32
}

func AlbumByArtist(artist string) ([]Album, error) {
	var albums []Album
	albums = append(albums, Album{
		Id:     1,
		Title:  "album-title-1",
		Artist: "album-artist-1",
		Price:  25.5,
	})
	albums = append(albums, Album{
		Id:     2,
		Title:  "album-title-2",
		Artist: "album-artist-2",
		Price:  25.5,
	})
	albums = append(albums, Album{
		Id:     3,
		Title:  "album-title-3",
		Artist: "album-artist-3",
		Price:  25.5,
	})
	return albums, nil
}
