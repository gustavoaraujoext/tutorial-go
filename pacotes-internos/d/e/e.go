package e

import (
	"fmt"

	"github.com/kustavo/tutorial-go/pacotes-internos/internal/a"
)

func E(call string) {
	fmt.Println("Método E Chamado por", call)
	a.A("D")
}