package page

import "backend-go/internal/pkg/config"

type Page struct {
	PageSize int `form:"page_size" json:"page_size"`
	Page     int `form:"page" json:"page"`
}

func (p Page) GetLimit() int {
	maxPageSize := config.Config.App.MaxPageSize
	if p.PageSize > maxPageSize {
		return maxPageSize
	} else if p.PageSize <= 0 {
		return config.Config.App.DefaultPageSize
	}
	return p.PageSize
}

func (p Page) GetOffset() int {
	if p.Page >= 1 {
		return (p.Page - 1) * p.PageSize
	}
	return 0
}
