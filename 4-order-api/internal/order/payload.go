package order

type CreateRequest struct {
	Products   []uint  `json:"products" validate:"required"`
	TotalPrice float64 `json:"total_price" validate:"required"`
}
