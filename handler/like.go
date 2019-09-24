package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/s-take/knowledge-like/model"
)

func AddLike(c echo.Context) error {
	Like := new(model.Like)
	if err := c.Bind(Like); err != nil {
		return err
	}

	if Like.KnowledgeID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	Like.UserID = uid
	model.CreateLike(Like)

	return c.JSON(http.StatusCreated, Like)
}

func GetLikes(c echo.Context) error {
	Likes := model.ListLikes()
	return c.JSON(http.StatusOK, Likes)
}

func GetLikesByUserID(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	Likes := model.FindLikes(&model.Like{UserID: uid})
	return c.JSON(http.StatusOK, Likes)
}
func GetLikesByKnowledgeID(c echo.Context) error {
	KnowledgeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	Likes := model.FindLikes(&model.Like{KnowledgeID: KnowledgeID})
	return c.JSON(http.StatusOK, Likes)
}

func DeleteLikeByUserID(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	if err := model.DeleteLike(&model.Like{UserID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
