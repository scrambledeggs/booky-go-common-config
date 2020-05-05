package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigDev(t *testing.T) {
	NewConfig()
	assert.Equal(t, 1,1)
}