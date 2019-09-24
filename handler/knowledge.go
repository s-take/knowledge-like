package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/s-take/knowledge-like/model"
)

// Knowledge functions

func AddKnowledge(c echo.Context) error {
	Knowledge := new(model.Knowledge)
	if err := c.Bind(Knowledge); err != nil {
		return err
	}

	if Knowledge.Title == "" || Knowledge.Author == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid fields",
		}
	}

	model.CreateKnowledge(Knowledge)

	return c.JSON(http.StatusCreated, Knowledge)
}

func GetKnowledges(c echo.Context) error {
	Knowledges := model.ListKnowledges()
	return c.JSON(http.StatusOK, Knowledges)
}

func DeleteKnowledge(c echo.Context) error {

	KnowledgeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteKnowledge(&model.Knowledge{ID: KnowledgeID}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
