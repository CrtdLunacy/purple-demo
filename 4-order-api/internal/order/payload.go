package order

type CreateRequest struct {
	Products   []uint  `JSON:"products" validate:"required"`
	TotalPrice float64 `json:"total_price" validate:"required"`
}
