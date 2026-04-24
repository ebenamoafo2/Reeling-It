package models

type Movie struct {
	ID          int
	TMDB_ID     int
	Title       string
	Tagline     *string
	ReleaseYear int
	Genres      []Genre
	Overview    *string  //it can be null in the database, so we use a pointer to string
	Score       *float32 //it can be null in the database, so we use a pointer to float32
	Popularity  *float32 //it can be null in the database, so we use a pointer to float32
	Keywords    []string
	Language    *string
	PosterURL   *string
	TrailerURL  *string
	Casting     []Actor
}
