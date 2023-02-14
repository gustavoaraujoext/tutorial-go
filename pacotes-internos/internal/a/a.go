package a

import (
	"fmt"

	"github.com/kustavo/tutorial-go/pacotes-internos/internal/c"
)

func A(call string) {
	fmt.Println("Método A Chamado por", call)
	c.C("A")
}
