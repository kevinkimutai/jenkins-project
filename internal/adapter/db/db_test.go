package db_test

import (
	"testing"

	config "github.com/kevinkimutai/metadata"

	"github.com/kevinkimutai/metadata/internal/adapter/db"
	"github.com/kevinkimutai/metadata/internal/adapter/queries"
	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {

	// Define test database URL
	testDBUrl := config.DATABASE_URL

	// Initialize a new DBAdapter
	adapter := db.NewDB(testDBUrl)

	// Test database initialization
	if adapter == nil {
		t.Fatal("DBAdapter initialization failed")
	}

}

func TestGetMovieById(t *testing.T) {

	testDBUrl := config.DATABASE_URL

	// Initialize a new DBAdapter
	adapter := db.NewDB(testDBUrl)

	// Test database initialization
	if adapter == nil {
		t.Fatal("DBAdapter initialization failed")
	}

	//Check no movie with ID
	movieID := int64(0)
	movie, err := adapter.GetMovieById(movieID)

	if err != nil {
		t.Fatalf("Error getting movie by id:%v", err)
	}

	assert.Equal(t, movie, queries.Movie{})

	//Check Valid ID
	movieID = int64(1)
	movie, err = adapter.GetMovieById(movieID)

	if err != nil {
		t.Fatalf("Error getting movie by id:%v", err)
	}

	assert.Equal(t, movie.Title, "Dark Knight 3")

}
