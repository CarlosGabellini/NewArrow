package player

import "math/rand"

type Gerador1 struct {}

func (g *Gerador1) GerarNumero() []float32 {
	geradorRun := make([]float32, 6)

	for i := range geradorRun {
		geradorRun[i] = rand.Float32() * 100
	}

	return geradorRun
}