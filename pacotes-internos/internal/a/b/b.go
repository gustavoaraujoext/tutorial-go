package b

import (
	"fmt"

	"github.com/kustavo/tutorial-go/pacotes-internos/internal/c"
)

func B(call string) {
	fmt.Println("Método B Chamado por", call)
	c.C("B")
}