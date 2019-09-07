package product

// Payload JSON of a product
type Payload struct {
	Name   string  `json:"name"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Weight int     `json:"weight"`
	Price  float64 `json:"price"`
}
