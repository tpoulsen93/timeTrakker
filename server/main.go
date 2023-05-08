package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// get the development mode from the commandline
	mode := flag.Bool("dev", false, "run the server in development mode (default to production mode)")
	flag.Parse()

	// if we're in dev mode, load the test environment
	if *mode {
		err := godotenv.Load("../.testEnv")
		if err != nil {
			log.Fatalf("Failed to load .testEnv file: %s\n", err)
		} else {
			log.Print("Using test environment\n")
		}
	} else {
		log.Print("Using production environment\n")
	}

	// make a new db connection pool
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	} else {
		log.Print("Successfully connected to the database\n")
	}
	defer pool.Close()

	// ping the database to verify we are connected
	pingErr := pool.Ping(ctx)
	if pingErr != nil {
		log.Fatalf("Failed to ping the database: %v\n", pingErr.Error())
	} else {
		log.Print("Successfully pinged the database\n")
	}

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	// router.Run("localhost:8080")
}

// ### I'm only keeping this code below as an example for the future ###
// ### It will be deleted as soon as we implement our first few endpoints ###

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop through the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }
