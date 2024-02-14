package entity

import "math"

type (
	Auth struct {
		User User
	}

	Filter struct {
		Page     int
		PageSize int
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
		TotalPages: int(math.Ceil((float64(totalItems) / float64(filter.PageSize)))),
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalItems: totalItems,
	}
}

func (f *Filter) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f *Filter) Limit() int {
	return f.PageSize
}
