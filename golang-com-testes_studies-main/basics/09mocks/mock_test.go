package main

import (
	"bytes"
	"testing"
)

type SleepSpy struct {
	Chamadas int
}

func (s *SleepSpy) Sleep() {
	s.Chamadas++
}

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}
	SleepSpy := &SleepSpy{}

	Contagem(buffer, SleepSpy)

	resultado := buffer.String()
	esperado := `3
2
1
Vai!`

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}

	if SleepSpy.Chamadas != 4 {
		t.Errorf("não houve chamadas suficientes do sleeper, esperado 4, resultado %d", SleepSpy.Chamadas)
	}
}
