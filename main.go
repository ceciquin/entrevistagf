package main

// Implementar un funcionamiento que valide una contraseña dada. Existen dos validaciones:
// 1) Mínimo 8 caracteres
// 2) Debe tener al menos 2 números y una letra
// Si la contraseña satisface solo una de las validaciones, devuelve "BUENA"
// Si la contraseña satisface ambas validaciones, devuelve "EXCELENTE"
// Si la contraseña no satisface ninguna, devuelve "DEBIL"

import (
	"unicode"
)

func validarContraseña(contraseña string) string {

	validacion := ""

	if validarNumerosYletra(contraseña) {

		if len(contraseña) >= 8 {
			validacion = "EXCELENTE"
		} else {
			validacion = "BUENA"
		}
	} else {
		validacion = "DEBIL"
	}

	return validacion

}

func validarNumerosYletra(contraseña string) bool {

	segundaValidacion := false
	sum := 0
	letras := false

	for _, c := range contraseña {

		if unicode.IsNumber(c) {
			sum++
		}
		if !letras && unicode.IsLetter(c) {
			letras = true
		}
	}

	if sum >= 2 && letras {
		segundaValidacion = true
	}

	return segundaValidacion
}

func main() {

}
