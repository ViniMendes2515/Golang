package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	// TDD - Test Driven Development
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{10, 15}
		area := retangulo.Area()
		if area != 150 {
			t.Fatalf("Esperado: 150, Recebido: %0.2f", area)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		area := circulo.Area()
		areaEsperada := (math.Pi * 100)
		if area != areaEsperada {
			t.Errorf("Esperado: %0.2f, Recebido: %0.2f", areaEsperada, area)
		}
	})
}
