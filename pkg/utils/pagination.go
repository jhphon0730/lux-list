package utils

import "strings"

const (
	DEFAULT_LIMIT = 10
	DEFAULT_PAGE  = 1
)

// CreatePaginationQuery는 검색 쿼리에서 limit와 page를 추출하여 반환하는 함수
func CreatePaginationQuery(search_query map[string]interface{}) (int, int) {
	limit := DEFAULT_LIMIT
	page := DEFAULT_PAGE

	if val, ok := search_query["limit"]; ok {
		if l, ok := val.(int); ok && l > 0 {
			limit = l
		}
	}

	if val, ok := search_query["page"]; ok {
		if p, ok := val.(int); ok && p > 0 {
			page = p
		}
	}

	return limit, page
}

// CreateOrderByQuery는 검색 쿼리에서 order_by를 추출하여 반환하는 함수
func CreateOrderByQuery(search_query map[string]interface{}) string {
	// (due_date DESC) ASC || DESC
	orderBy := "due_date DESC" // 기본값
	if val, ok := search_query["order_by"]; ok {
		if order, ok := val.(string); ok && (order == "ASC" || order == "DESC" || order == "asc" || order == "desc") {
			orderBy = "due_date " + strings.ToUpper(order)
		}
	}
	return orderBy
}
