package model

type KanbanColumn struct {
	ColName string
	Color   string
	Limit   int
	Items   []*KanbanItem
}

type KanbanItem struct {
	Title      string `form:"Title"`
	Detail     string `form:"Detail"`
	ItemStatus string `form:"ItemStatus"`
	Deadline   string `form:"Deadline"`
}
