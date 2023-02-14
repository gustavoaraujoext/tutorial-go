package d

import (
	"fmt"

	"github.com/kustavo/tutorial-go/pacotes-internos/internal/a"
)

func D(call string) {
	fmt.Println("Método D Chamado por", call)
	a.A("D")
}