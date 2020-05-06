package climacell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataFieldString(t *testing.T) {
	v := Temp
	assert.Equal(t, "temp", v.String())
}

func TestUnitSystemString(t *testing.T) {
	v := US
	assert.Equal(t, "us", v.String())
}
