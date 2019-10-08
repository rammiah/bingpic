package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	assert.NotEqual(t, nil, db, "db should not null")
}
