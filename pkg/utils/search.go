package utils

import (
	"github.com/gin-gonic/gin"
)

func GetTasksSearchQuery(c *gin.Context) map[string]interface{} {
	query := make(map[string]interface{})

	if limit := c.Query("limit"); limit != "" {
		query["limit"] = InterfaceToInt(limit)
	}
	if page := c.Query("page"); page != "" {
		query["page"] = InterfaceToInt(page)
	}
	if title := c.Query("title"); title != "" {
		query["title"] = title
	}
	if isCompleted := c.Query("is_completed"); isCompleted != "" {
		query["is_completed"] = InterfaceToBool(isCompleted)
	}
	if priority := c.Query("priority"); priority != "" {
		query["priority"] = priority
	}
	if dueDate := c.Query("due_date"); dueDate != "" {
		query["due_date"] = dueDate
	}
	if orderBy := c.Query("order_by"); orderBy != "" {
		query["order_by"] = orderBy
	}

	return query
}
