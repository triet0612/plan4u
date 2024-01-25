package handler

import (
	"log"
	"plan4u/model"

	"github.com/labstack/echo/v4"
)

func (a *AppHandler) Kanban(c echo.Context) error {
	rows, err := a.Db.Query("SELECT colname, color, col_limit FROM kanban_column;")
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(500, "database error")
	}
	kCols := []*model.KanbanColumn{}
	defer rows.Close()
	for rows.Next() {
		tempRow := &model.KanbanColumn{}
		if err := rows.Scan(&tempRow.ColName, &tempRow.Color, &tempRow.Limit); err != nil {
			log.Println(err)
			continue
		}
		itemRows, err := a.Db.Query("select title, detail, item_status, deadline from kanban_item where item_status=?;", tempRow.ColName)
		if err != nil {
			log.Println(err)
			return echo.ErrInternalServerError
		}
		tempRow.Items = []*model.KanbanItem{}
		defer rows.Close()
		for itemRows.Next() {
			kItem := &model.KanbanItem{}
			if err = itemRows.Scan(&kItem.Title, &kItem.Detail, &kItem.ItemStatus, &kItem.Deadline); err != nil {
				log.Println(err)
				continue
			}
			tempRow.Items = append(tempRow.Items, kItem)
		}
		kCols = append(kCols, tempRow)
	}
	return c.Render(200, "kanban.html", kCols)
}

func (a *AppHandler) KanbanDetail(c echo.Context) error {
	return c.Render(200, "kanban-detail.html", nil)
}

func (a *AppHandler) KanbanCreateForm(c echo.Context) error {
	rows, err := a.Db.Query("select colname from kanban_column")
	if err != nil {
		return echo.ErrInternalServerError
	}
	colList := []string{}
	for rows.Next() {
		var tempStr string
		if err := rows.Scan(&tempStr); err != nil {
			log.Println(err)
			continue
		}
		colList = append(colList, tempStr)
	}
	return c.Render(200, "kanban-create.html", colList)
}

func (a *AppHandler) KanbanCreate(c echo.Context) error {
	var newItem model.KanbanItem
	if err := c.Bind(&newItem); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(400, "invalid form inputs")
	}
	if _, err := a.Db.Exec(
		"insert into kanban_item values (?, ?, ?, ?);",
		newItem.Title, newItem.Detail, newItem.ItemStatus, newItem.Deadline,
	); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.HTML(200, "ok")
}
