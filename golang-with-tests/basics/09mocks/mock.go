package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}

	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}

	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
