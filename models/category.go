package models

type Category struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ProductCount int    `json:"product_count,omitempty"`
}
