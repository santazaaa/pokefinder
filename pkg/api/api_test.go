package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByNameOrID_ShouldReturnDataIfFound(t *testing.T) {
	name := "ditto"
	res, err := FindByNameOrID(name)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Name, name)
}

func TestFindByNameOrID_ShouldReturnErrorIfNotFound(t *testing.T) {
	name := "some_invalid_name"
	_, err := FindByNameOrID(name)
	assert.NotNil(t, err)
}

func TestFindEncountersByNameOrID_ShouldReturnDataIfFound(t *testing.T) {
	name := "ditto"
	res, err := FindEncountersByNameOrID(name)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestFindEncountersByNameOrID_ShouldReturnErrorIfNotFound(t *testing.T) {
	name := "some_invalid_name"
	_, err := FindEncountersByNameOrID(name)
	assert.NotNil(t, err)
}
