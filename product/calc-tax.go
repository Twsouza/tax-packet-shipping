package product

import (
	"fmt"

	"github.com/twsouza/tax-packet-shipping/common"
)

const (
	defaultFactor float64 = 300.0
	minimumKg     float64 = 100.0
	minimumTax    float64 = 0.05
	gramsToKg     float64 = 1000.0
)

// CalcTax calculate shipping tax
func (p *Payload) CalcTax() float64 {
	volume := p.Height * p.Width * p.Length
	calculatedWeight := volume * defaultFactor

	fmt.Printf("%+v volume is %v and the calculated weight is %v\n", p.Name, volume, calculatedWeight)

	if calculatedWeight <= minimumKg {
		fmt.Printf("calculated weight is less than minimum kg\n")
		return p.Price * minimumTax
	}
	fmt.Printf("calculated weight is greater than minimum kg\n")

	weightInKg := float64(p.Weight) / gramsToKg
	tax := (calculatedWeight * weightInKg)

	return common.ToFixed(tax, 2)
}
