# Método

// TODO:

Um método é apenas uma função com um argumento de receptor. É declarado com a mesma sintaxe com a adição do receptor.

```go
type Person struct {
    Name string
    Age  int
}

func (p *Person) isAdult bool {
  return p.Age > 18
}
```

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

## Referências
