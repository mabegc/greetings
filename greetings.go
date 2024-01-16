package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a message.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func randomFormat() string {
	formats := []string{
		"|hola, %v! ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de verte!",
		"¡Qué gusto conocerte, %v!",
		"¡Te amo %v!",
		"%v",
	}

	return formats[rand.Intn(len(formats))]
}