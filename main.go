package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Kind of Blue", Artist: "Miles Davis", Price: 45.99},
	{ID: "5", Title: "A Love Supreme", Artist: "John Coltrane", Price: 49.99},
	{ID: "6", Title: "Time Out", Artist: "Dave Brubeck", Price: 22.99},
	{ID: "7", Title: "Mingus Ah Um", Artist: "Charles Mingus", Price: 34.99},
	{ID: "8", Title: "The Shape of Jazz to Come", Artist: "Ornette Coleman", Price: 29.99},
	{ID: "9", Title: "Head Hunters", Artist: "Herbie Hancock", Price: 31.99},
	{ID: "10", Title: "Bitches Brew", Artist: "Miles Davis", Price: 40.99},
	{ID: "11", Title: "Giant Steps", Artist: "John Coltrane", Price: 37.99},
	{ID: "12", Title: "Ella and Louis", Artist: "Ella Fitzgerald & Louis Armstrong", Price: 26.99},
	{ID: "13", Title: "Take Five", Artist: "Dave Brubeck Quartet", Price: 23.99},
	{ID: "14", Title: "Maiden Voyage", Artist: "Herbie Hancock", Price: 28.99},
	{ID: "15", Title: "My Favorite Things", Artist: "John Coltrane", Price: 33.99},
	{ID: "16", Title: "Moanin'", Artist: "Art Blakey & The Jazz Messengers", Price: 20.99},
	{ID: "17", Title: "Kind of Blue", Artist: "Miles Davis", Price: 45.99},
	{ID: "18", Title: "Birth of the Cool", Artist: "Miles Davis", Price: 24.99},
	{ID: "19", Title: "Black Saint and the Sinner Lady", Artist: "Charles Mingus", Price: 36.99},
	{ID: "20", Title: "Miles Smiles", Artist: "Miles Davis", Price: 27.99},
	{ID: "21", Title: "Budapest Concert", Artist: "Keith Jarrett", Price: 32.99},
	{ID: "22", Title: "Bird and Diz", Artist: "Charlie Parker & Dizzy Gillespie", Price: 19.99},
	{ID: "23", Title: "The Sidewinder", Artist: "Lee Morgan", Price: 25.99},
	{ID: "24", Title: "Saxophone Colossus", Artist: "Sonny Rollins", Price: 21.99},
	{ID: "25", Title: "Porgy and Bess", Artist: "Miles Davis", Price: 38.99},
	{ID: "26", Title: "The Epic", Artist: "Kamasi Washington", Price: 42.99},
	{ID: "27", Title: "Empyrean Isles", Artist: "Herbie Hancock", Price: 30.99},
	{ID: "28", Title: "Money Jungle", Artist: "Duke Ellington, Charles Mingus, Max Roach", Price: 35.99},
	{ID: "29", Title: "Go!", Artist: "Dexter Gordon", Price: 18.99},
	{ID: "30", Title: "Empyrean Isles", Artist: "Herbie Hancock", Price: 30.99},
	{ID: "31", Title: "Somethin' Else", Artist: "Cannonball Adderley", Price: 26.99},
	{ID: "32", Title: "A Night in Tunisia", Artist: "Art Blakey & The Jazz Messengers", Price: 29.99},
	{ID: "33", Title: "Speak No Evil", Artist: "Wayne Shorter", Price: 31.99},
	{ID: "34", Title: "Brilliant Corners", Artist: "Thelonious Monk", Price: 33.99},
	{ID: "35", Title: "Miles Ahead", Artist: "Miles Davis", Price: 28.99},
	{ID: "36", Title: "The Bridge", Artist: "Sonny Rollins", Price: 24.99},
	{ID: "37", Title: "The Black Saint and the Sinner Lady", Artist: "Charles Mingus", Price: 36.99},
	{ID: "38", Title: "Maiden Voyage", Artist: "Herbie Hancock", Price: 28.99},
	{ID: "39", Title: "Ella Fitzgerald Sings the Cole Porter Song Book", Artist: "Ella Fitzgerald", Price: 37.99},
	{ID: "40", Title: "Out to Lunch", Artist: "Eric Dolphy", Price: 23.99},
	{ID: "41", Title: "The Köln Concert", Artist: "Keith Jarrett", Price: 32.99},
	{ID: "42", Title: "My Funny Valentine", Artist: "Miles Davis", Price: 26.99},
	{ID: "43", Title: "The Real McCoy", Artist: "McCoy Tyner", Price: 29.99},
	{ID: "44", Title: "Chet Baker Sings", Artist: "Chet Baker", Price: 21.99},
	{ID: "45", Title: "Soul Station", Artist: "Hank Mobley", Price: 25.99},
	{ID: "46", Title: "Ellington at Newport", Artist: "Duke Ellington", Price: 27.99},
	{ID: "47", Title: "Kind of Blue", Artist: "Miles Davis", Price: 45.99},
	{ID: "48", Title: "The Shape of Jazz to Come", Artist: "Ornette Coleman", Price: 29.99},
	{ID: "49", Title: "Giant Steps", Artist: "John Coltrane", Price: 37.99},
	{ID: "50", Title: "A Love Supreme", Artist: "John Coltrane", Price: 49.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
