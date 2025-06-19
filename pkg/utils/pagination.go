package utils

const (
	DEFAULT_LIMIT = 10
	DEFAULT_PAGE  = 1
)

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
