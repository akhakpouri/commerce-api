package tax

import dto "commerce/api/internal/dto/tax"

type TaxServiceI interface {
	GetStates() ([]dto.Tax, error)
	Calculate(amount float32, state string) (float32, error)
}

type TaxService struct {
}

func NewTaxService() TaxServiceI {
	return &TaxService{}
}

// Calculate implements [TaxServiceI].
func (t *TaxService) Calculate(amount float32, state string) (float32, error) {
	panic("unimplemented")
}

// GetStates implements [TaxServiceI].
func (t *TaxService) GetStates() ([]dto.Tax, error) {
	return stateTaxes, nil
}

var stateTaxes = []dto.Tax{
	{State: "AL", Amount: 0.04},
	{State: "AK", Amount: 0.00},
	{State: "AZ", Amount: 0.056},
	{State: "AR", Amount: 0.065},
	{State: "CA", Amount: 0.0725},
	{State: "CO", Amount: 0.029},
	{State: "CT", Amount: 0.0635},
	{State: "DE", Amount: 0.00},
	{State: "FL", Amount: 0.06},
	{State: "GA", Amount: 0.04},
	{State: "HI", Amount: 0.04},
	{State: "ID", Amount: 0.06},
	{State: "IL", Amount: 0.0625},
	{State: "IN", Amount: 0.07},
	{State: "IA", Amount: 0.06},
	{State: "KS", Amount: 0.065},
	{State: "KY", Amount: 0.06},
	{State: "LA", Amount: 0.0445},
	{State: "ME", Amount: 0.055},
	{State: "MD", Amount: 0.06},
	{State: "MA", Amount: 0.0625},
	{State: "MI", Amount: 0.06},
	{State: "MN", Amount: 0.06875},
	{State: "MS", Amount: 0.07},
	{State: "MO", Amount: 0.04225},
	{State: "MT", Amount: 0.00},
	{State: "NE", Amount: 0.055},
	{State: "NV", Amount: 0.0685},
	{State: "NH", Amount: 0.00},
	{State: "NJ", Amount: 0.06625},
	{State: "NM", Amount: 0.04875},
	{State: "NY", Amount: 0.04},
	{State: "NC", Amount: 0.0475},
	{State: "ND", Amount: 0.05},
	{State: "OH", Amount: 0.0575},
	{State: "OK", Amount: 0.045},
	{State: "OR", Amount: 0.00},
	{State: "PA", Amount: 0.06},
	{State: "RI", Amount: 0.07},
	{State: "SC", Amount: 0.06},
	{State: "SD", Amount: 0.045},
	{State: "TN", Amount: 0.07},
	{State: "TX", Amount: 0.0625},
	{State: "UT", Amount: 0.0595},
	{State: "VT", Amount: 0.06},
	{State: "VA", Amount: 0.053},
	{State: "WA", Amount: 0.065},
	{State: "WV", Amount: 0.06},
	{State: "WI", Amount: 0.05},
	{State: "WY", Amount: 0.04},
	{State: "DC", Amount: 0.06},
}
