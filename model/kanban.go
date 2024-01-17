package model

import "time"

type KanbanColumn struct {
	ColName string
	Color   string
	Limit   int
}

type KanbanItem struct {
	Title       string
	Detail      string
	Item_status string
	Deadline    time.Time
}
