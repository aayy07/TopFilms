// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Results struct {
	Adult            *string  `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	ID               *int     `json:"id"`
	Title            *string  `json:"title"`
	OriginalLanguage *string  `json:"original_language"`
	OriginalTitle    *string  `json:"original_title"`
	Overview         *string  `json:"overview"`
	PosterPath       *string  `json:"poster_path"`
	MediaType        *string  `json:"media_type"`
	GenreIds         []*int   `json:"genre_ids"`
	Popularity       *float64 `json:"popularity"`
	ReleaseDate      *string  `json:"release_date"`
	Video            *bool    `json:"video"`
	VoteAverage      *float64 `json:"vote_average"`
	VoteCount        *int     `json:"vote_count"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type TrendingMoviesResponse struct {
	Page         *int       `json:"page"`
	Results      []*Results `json:"results"`
	TotalPages   *int       `json:"total_pages"`
	TotalResults *int       `json:"total_results"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
