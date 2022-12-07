package models

type Pagination[T any] struct {
	Total       int64 `json:"total"`
	Data        []T   `json:"data"`
	PerPage     int64 `json:"per_page"`
	TotalPages  int64 `json:"total_pages"`
	IsLastPage  bool  `json:"is_last_page"`
	IsFirstPage bool  `json:"is_first_page"`
}

func Paginate[T any](total int64, count int64, page int64, data []T) Pagination[T] {
	totalPages := total / count

	if total%count != 0 {
		totalPages = totalPages + 1
	}

	if page == 0 {
		page = 1
	}

	return Pagination[T]{
		Data:        data,
		Total:       total,
		PerPage:     count,
		IsLastPage:  page >= totalPages,
		IsFirstPage: page == 1,
		TotalPages:  totalPages,
	}
}
