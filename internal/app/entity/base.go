package entity

import "math"

type (
	Auth struct {
		User User
	}

	Filter struct {
		page     int
		pageSize int
	}

	Page[T any] struct {
		Items      []T
		Page       int
		PageSize   int
		TotalPages int
		TotalItems int
	}
)

func NewPage[T any](items []T, filter Filter, totalItems int) Page[T] {
	return Page[T]{
		Items:      items,
		TotalPages: int(math.Ceil((float64(totalItems) / float64(filter.pageSize)))),
		Page:       filter.page,
		PageSize:   filter.pageSize,
		TotalItems: totalItems,
	}
}

func NewFilter(page int, pageSize int) Filter {
	filter := Filter{
		page:     1,
		pageSize: 10,
	}

	if page > 1 {
		filter.page = page
	}
	if pageSize >= 1 && pageSize <= 100 {
		filter.pageSize = pageSize
	} else if pageSize > 100 {
		filter.pageSize = 100
	}

	return filter
}

func (f Filter) Page() int {
	return f.page
}

func (f Filter) PageSize() int {
	return f.pageSize
}
