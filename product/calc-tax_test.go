package product

import "testing"

func TestCalcTax(t *testing.T) {
	p := &Payload{
		Name:   "Magic Mouse",
		Height: 0.75,
		Length: 1.1,
		Width:  0.6,
		Weight: 400,
		Price:  150,
	}
	tax := p.CalcTax()

	if tax != 59.4 {
		t.Error("Expected 59.4 but got ", tax)
	}

	p2 := &Payload{
		Name:   "Magic Mouse",
		Height: 0.5,
		Length: 0.5,
		Width:  1.1,
		Weight: 300,
		Price:  170,
	}
	tax2 := p2.CalcTax()

	if tax2 != 8.5 {
		t.Error("Expected 8.5 but got ", tax2)
	}

}
