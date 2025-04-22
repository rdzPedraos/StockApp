package paginate

import (
	"app/database/connection"
	"app/database/scopes"
	"app/resources"
	"math"

	"gorm.io/gorm"
)

type statement func(*gorm.DB) *gorm.DB

type PaginateResolver[T interface{}] struct {
	stmt     statement
	resource resources.TypedResource[T]
	page     int
	perPage  int
}

type Meta struct {
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PerPage  int   `json:"per_page"`
	LastPage int   `json:"last_page"`
}

type PaginationResponse struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

func (s *PaginateResolver[T]) Stmt(stmt statement) *PaginateResolver[T] {
	s.stmt = stmt
	return s
}

func (s *PaginateResolver[T]) Resource(resource resources.TypedResource[T]) *PaginateResolver[T] {
	s.resource = resource
	return s
}

func (s *PaginateResolver[T]) Paginate(page int, perPage int) *PaginateResolver[T] {
	s.page = page
	s.perPage = perPage
	return s
}

func (s PaginateResolver[T]) Resolve() PaginationResponse {
	return PaginationResponse{
		Data: s.getData(),
		Meta: s.getMeta(),
	}
}

func (s PaginateResolver[T]) getQuery() *gorm.DB {
	query := connection.DB.Model(new(T))

	if s.stmt != nil {
		query = s.stmt(query)
	}

	return query
}

func (s PaginateResolver[T]) getMeta() Meta {
	var total int64

	countQuery := s.getQuery().Session(&gorm.Session{})
	countQuery.Count(&total)

	lastPage := math.Ceil(float64(total) / float64(s.perPage))

	return Meta{
		Total:    total,
		Page:     s.page,
		PerPage:  s.perPage,
		LastPage: int(lastPage),
	}
}

func (s PaginateResolver[T]) getData() interface{} {
	query := s.getQuery()

	var results []T
	query.Scopes(scopes.Pagination(s.page, s.perPage)).Find(&results)

	if s.resource != nil {
		return resources.FormatArrayWith(s.resource, results)
	}

	return results
}
