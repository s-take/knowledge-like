package model

import "fmt"

type Like struct {
	ID          int `json:"id" gorm:"praimaly_key"`
	UserID      int `json:"userId"`
	User        User
	KnowledgeID int `json:"knowledgeId"`
	Knowledge   Knowledge
	Reason      string `json:"reason"`
	Comment     string `json:"comment"`
}

type Likes []Like

func CreateLike(Like *Like) {
	db.Create(Like)
}

func ListLikes() Likes {
	var likes Likes
	// db.Find(&likes)
	db.Preload("Knowledge").Find(&likes)
	return likes
}

func FindLikes(t *Like) Likes {
	var Likes Likes
	db.Where(t).Find(&Likes)
	return Likes
}

func DeleteLike(t *Like) error {
	if rows := db.Where(t).Delete(&Like{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Like (%v) to delete", t)
	}
	return nil
}
