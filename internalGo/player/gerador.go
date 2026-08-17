package player

import "math/rand"

type Gerador1 struct {}

func (g *Gerador1) GerarNumero() (float32) {
	var number1 float32 = rand.Float32() * 100
	return number1
}