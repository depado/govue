package entry

import (
	"net/http"
	"strconv"

	"github.com/Depado/govue/hateoas"
	"github.com/gin-gonic/gin"
)

// Bucket is the name of the bucket storing all the entries
const (
	Bucket = "entries"
	Type   = "entry"
)

// Entry is the main struct
type Entry struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

// Validate validates that all the required files are not empty.
func (e Entry) Validate() hateoas.Errors {
	var errors hateoas.Errors
	if e.Title == "" {
		errors = append(errors, hateoas.Error{
			Status: http.StatusBadRequest,
			Title:  "title field is required",
		})
	}
	if e.Markdown == "" {
		errors = append(errors, hateoas.Error{
			Status: http.StatusBadRequest,
			Title:  "markdown field is required",
		})
	}
	return errors
}

// Data contains the Type of the request and the Attributes
type Data struct {
	Type       string `json:"type,omitempty"`
	Attributes *Entry `json:"attributes,omitempty"`
	Links      *Links `json:"links,omitempty"`
}

// Links represent a list of links
type Links map[string]string

// Wrapper is the HATEOAS wrapper
type Wrapper struct {
	Data   *Data           `json:"data,omitempty"`
	Errors *hateoas.Errors `json:"errors,omitempty"`
}

// MultiWrapper is a wrapper that can accept multiple Data
type MultiWrapper struct {
	Data   *[]Data         `json:"data,omitempty"`
	Errors *hateoas.Errors `json:"errors,omitempty"`
}

// Post is the handler to POST a new Entry
func Post(c *gin.Context) {
	var err error
	var json = Wrapper{}

	if err = c.BindJSON(&json); err == nil {
		errors := json.Data.Attributes.Validate()
		if len(errors) > 0 {
			c.JSON(http.StatusBadRequest, Wrapper{Errors: &errors})
			return
		}
		if err = json.Data.Attributes.Save(); err != nil {
			json.Data = nil
			json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "could not save entry"}}
			c.JSON(http.StatusInternalServerError, json)
		} else {
			json.Data.Links = &Links{"self": c.Request.URL.RequestURI() + strconv.Itoa(json.Data.Attributes.ID)}
			c.JSON(http.StatusCreated, json)
		}
	} else {
		json.Data = nil
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "Bad json format"}}
		c.JSON(http.StatusBadRequest, json)
	}
}

// List lists the entries
func List(c *gin.Context) {
	var json = MultiWrapper{}
	var datas = []Data{}

	entr, err := All()
	if err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "could not retrieve entries"}}
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	for i := range entr {
		datas = append(datas, Data{Type: Type, Attributes: &entr[i]})
	}
	json.Data = &datas
	c.JSON(http.StatusOK, json)
}

// Get is the handler to GET an existing entry
func Get(c *gin.Context) {
	var err error
	var e Entry
	var json = Wrapper{}

	id := c.Param("id")
	if err = e.Get(id); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusNotFound, Title: "id could not be found"}}
		c.JSON(http.StatusNotFound, json)
		return
	}
	if e.ID, err = strconv.Atoi(id); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "id can't be parsed"}}
		c.JSON(http.StatusInternalServerError, json)
	}
	json.Data = &Data{Type: Type, Attributes: &e}
	c.JSON(http.StatusOK, json)
}

// Patch is used to update a resource.
func Patch(c *gin.Context) {
	var err error
	var e Entry
	var json = Wrapper{}

	id := c.Param("id")
	if err = e.Get(id); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusNotFound, Title: "id could not be found"}}
		c.JSON(http.StatusNotFound, json)
		return
	}
	if e.ID, err = strconv.Atoi(id); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "id can't be parsed"}}
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	json.Data = &Data{Type: Type, Attributes: &e}
	if err = c.BindJSON(&json); err == nil {
		if err = json.Data.Attributes.Save(); err != nil {
			json.Data = nil
			json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "could not save entry"}}
			c.JSON(http.StatusInternalServerError, json)
		} else {
			c.JSON(http.StatusCreated, json)
		}
	} else {
		json.Data = nil
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "Bad json format"}}
		c.JSON(http.StatusBadRequest, json)
	}
}

// Delete deletes a resource
func Delete(c *gin.Context) {
	var err error
	var e Entry
	var json = Wrapper{}

	id := c.Param("id")
	if err = e.Get(id); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusNotFound, Title: "id could not be found"}}
		c.JSON(http.StatusNotFound, json)
		return
	}
	if e.ID, err = strconv.Atoi(id); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "id can't be parsed"}}
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	if err = e.Delete(); err != nil {
		json.Errors = &hateoas.Errors{hateoas.Error{Status: http.StatusInternalServerError, Title: "couldn't delete resource"}}
		c.JSON(http.StatusInternalServerError, json)
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}
