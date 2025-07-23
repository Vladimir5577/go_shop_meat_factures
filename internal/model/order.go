package model

type Order struct {
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

type OrderItem struct {
	ProductId uint    `json:"product_id"`
	Amount    uint    `json:"amount"`
	SummItem  float32 `json:"summ_item"`
}

type Ordering struct {
	UserId    uint        `json:"user_id"`
	Products  []OrderItem `json:"products"`
	TotalSumm float32     `json:"total_summ"`
	Status    string      `json:"status"`
	Comment   string      `json:"comment"`
}
