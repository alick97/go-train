package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Artist     string    `json:"artist"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedAt1 time.Time `json:"created_at1"`
	CreatedAt2 time.Time `json:"created_at2"`
	CreatedAt3 time.Time `json:"created_at3"`
}

const tmplStr = `{
	"id": "{{.ID}}",
	"title": "{{.Title}}",
	"artist": "{{.Artist}}",
	"price": {{.Price}},
	"created_at": "{{.CreatedAt}}",
	"created_at1": "{{.CreatedAt1}}",
	"created_at2": "{{.CreatedAt2}}",
	"created_at3": "{{.CreatedAt3}}"
}`

// func (a album) MarshalJSON() ([]byte, error) {
// 	// t, _ := template.New("t").Parse(tmplStr)
// 	// var s bytes.Buffer
// 	// t.Execute(&s, a)
//
// 	return []byte(fmt.Sprintf(`{"id":"%s","title":"%s",
// 	"artist": "%s",
// 	"price":%.3f,
// 	"created_at": "%s",
// 	"created_at1":"%s",
// 	"created_at2":"%s",
// 	"created_at3":"%s"
// 	}`, a.ID, a.Title, a.Title, a.Price,
// 		a.CreatedAt.Format(time.RFC3339),
// 		a.CreatedAt1.Format(time.RFC3339),
// 		a.CreatedAt2.Format(time.RFC3339),
// 		a.CreatedAt3.Format(time.RFC3339),
// 	)), nil
// }

//	func (a album) MarshalJSON() ([]byte, error) {
//		// t, _ := template.New("t").Parse(tmplStr)
//		// var s bytes.Buffer
//		// t.Execute(&s, a)
//
//		return []byte(fmt.Sprintf(`{"id":"%s","title":"%s",
//		"artist": "%s",
//		"price":%.3f,
//		"created_at": "%s"
//		}`, a.ID, a.Title, a.Title, a.Price,
//			a.CreatedAt.Format(time.RFC3339),
//		)), nil
//	}
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99,
		CreatedAt:  time.Now().UTC(),
		CreatedAt1: time.Now().UTC(),
		CreatedAt2: time.Now().UTC(),
		CreatedAt3: time.Now().UTC(),
	},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var aStr []byte

func getAlbums(c *gin.Context) {
	// p := profile.Start()
	// defer p.Stop()
	// c.JSON(http.StatusOK, albums)
	c.PureJSON(http.StatusOK, albums)
	// c.AsciiJSON(http.StatusOK, albums)
	// c.Data(http.StatusOK, "application/json", aStr)
}

func init() {
	b, _ := json.Marshal(albums[0])
	for i := 0; i < 5000; i++ {
		albums = append(albums, albums[0])
		aStr = append(aStr, b...)
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:3399")
}
