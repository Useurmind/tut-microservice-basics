package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/useurmind/tut-microservice-basics/pkg/database"
	"github.com/useurmind/tut-microservice-basics/pkg/model"
)


func handleWithDB(c *gin.Context, handler func(c *gin.Context, db *sqlx.DB)) {
	db, err := database.GetConnection()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	handler(c, db)
}

func HandlePeople(route string, r gin.IRouter) {
	r.POST(route, func(c *gin.Context) { handleWithDB(c, handleInsertPeople) })
	r.GET(route, func(c *gin.Context) { handleWithDB(c, handleGetPeople) })
	r.GET(route + "/:id", func(c *gin.Context) { handleWithDB(c, handleGetPeopleById) })
	r.DELETE(route + "/:id", func(c *gin.Context) { handleWithDB(c, handleDeletePeopleById) })
}

func handleInsertPeople(c *gin.Context, db *sqlx.DB) {
	people := &model.People{}
	err := c.BindJSON(&people)
	if err != nil {
		return
	}
	people, err = database.InsertPeople(c, db, people) 
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(200, people)
}

func handleGetPeople(c *gin.Context, db *sqlx.DB) {
	people, err := database.GetPeople(c, db)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, people)
}

func handleGetPeopleById(c *gin.Context, db *sqlx.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	people, err := database.GetPeopleById(c, db, id)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if people == nil {
		c.Status(404)
		return
	}

	c.JSON(200, people)
}

func handleDeletePeopleById(c *gin.Context, db *sqlx.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	deleted, err := database.DeletePeopleById(c, db, id)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if !deleted {
		c.Status(404)
		return
	}

	c.Status(200)

}