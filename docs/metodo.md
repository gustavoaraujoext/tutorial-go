# Método

## Composição

Abaixo um exemplo do uso de composição:

```go
package main

import "fmt"

type Animal struct{
    
}

func (a Animal) Comer() {
    fmt.Println("Comendo")
}

type MembroFamilia struct{
    
}

func (mf MembroFamilia) Nome() {
    fmt.Println("Meu nome não é Johnny")
}

type Cachorro struct {
    Animal        // Struct incorporada
    MembroFamilia // Struct incorporada
}

func main() {
    d := Cachorro{}
    d.Comer() // "Comendo"
    d.Nome()  // "Meu nome não é Johnny"
}
```