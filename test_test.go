package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarContraseña(t *testing.T) {

	res1 := validarContraseña("test123")

	assert.Equal(t, "BUENA", res1)

	res2 := validarContraseña("test1234567")

	assert.Equal(t, "EXCELENTE", res2)

	res3 := validarContraseña("test")

	assert.Equal(t, "DEBIL", res3)

	res4 := validarContraseña("123456789")

	assert.Equal(t, "DEBIL", res4)

}
