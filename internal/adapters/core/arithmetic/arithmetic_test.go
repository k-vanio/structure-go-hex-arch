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
