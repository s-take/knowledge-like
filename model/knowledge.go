package model

import "fmt"

type Knowledge struct {
	// id, title, url, author
	ID     int    `json:"id" gorm:"praimaly_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Knowledges []Knowledge

func CreateKnowledge(Knowledge *Knowledge) {
	db.Create(Knowledge)
}

func ListKnowledges() Knowledges {
	var knowledges Knowledges
	db.Find(&knowledges)
	return knowledges
}

func DeleteKnowledge(t *Knowledge) error {
	if rows := db.Where(t).Delete(&Knowledge{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Knowledge (%v) to delete", t)
	}
	return nil
}
