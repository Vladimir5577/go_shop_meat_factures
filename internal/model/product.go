package model

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CategoryId  uint    `json:"category_id"`
	InStock     bool    `json:"in_stock"`
	IsActive    bool    `json:"is_active"`
	Created     string  `json:"created_at"`
	Updated     string  `json:"updated_at"`
}

type ProductResponse struct {
	Id          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float32          `json:"price"`
	Category    CategoryResponse `json:"category"`
	InStock     bool             `json:"in_stock"`
	IsActive    bool             `json:"is_active"`
	Created     string           `json:"created_at"`
	Updated     string           `json:"updated_at"`
}
