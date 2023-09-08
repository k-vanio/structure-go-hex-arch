package arithmetic_test

import (
	"testing"

	"github.com/k-vanio/structure-go-hex-arch/internal/adapters/core/arithmetic"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	ar := arithmetic.NewAdapter()

	result, err := ar.Addition(2, 5)

	assert.Nil(t, err)
	assert.Equal(t, int32(7), result)
}

func TestSubtraction(t *testing.T) {
	ar := arithmetic.NewAdapter()

	result, err := ar.Subtraction(20, 5)

	assert.Nil(t, err)
	assert.Equal(t, int32(15), result)
}

func TestMultiplication(t *testing.T) {
	ar := arithmetic.NewAdapter()

	result, err := ar.Multiplication(20, 5)

	assert.Nil(t, err)
	assert.Equal(t, int32(100), result)
}

func TestDivision(t *testing.T) {
	ar := arithmetic.NewAdapter()

	result, err := ar.Division(20, 5)

	assert.Nil(t, err)
	assert.Equal(t, int32(4), result)
}
