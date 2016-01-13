package entry

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Bucket is the name of the bucket storing all the entries
const Bucket = "entries"

// Entry is the main struct
type Entry struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

// APIError represents a single API Error
type APIError struct {
	ID     int    `json:"id,omitempty"`
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// APIErrors represent multiple API Errors
type APIErrors []APIError

// Post is the handler to POST a new Entry
func Post(c *gin.Context) {
	var err error
	var errors APIErrors
	var json Entry

	if err = c.BindJSON(&json); err == nil {
		if json.Markdown == "" {
			errors = append(errors, APIError{
				Status: http.StatusBadRequest,
				Title:  "markdown field is required",
			})
		}
		if json.Title == "" {
			errors = append(errors, APIError{
				Status: http.StatusBadRequest,
				Title:  "title field is required",
			})
		}
		if len(errors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		if err = json.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": json})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"errors": fmt.Sprintf("%s", err)})
	}
}
