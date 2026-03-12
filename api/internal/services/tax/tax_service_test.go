package tax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStates(t *testing.T) {
	svc := NewTaxService()

	states := svc.GetStates()
	fmt.Printf("total number of states is %d \n", len(states))
	assert.Equal(t, 51, len(states), "They should be equal")
}

func TestCalculate(t *testing.T) {
	svc := NewTaxService()
	amount, state := 100.00, "MD"
	tax, err := svc.Calculate(amount, state)
	assert.NoError(t, err)
	assert.Equal(t, 6.0, *tax, "They should be equal")
}

func TestInvalidStateCalculate(t *testing.T) {
	svc := NewTaxService()
	amount, state := 100.00, "BC"
	tax, err := svc.Calculate(amount, state)
	assert.NoError(t, err)
	assert.Nil(t, tax)
}
