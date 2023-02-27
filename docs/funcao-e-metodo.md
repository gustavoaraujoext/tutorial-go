# Função e Método

## Função

Uma função é um grupo de instruções que existem dentro de um programa com a finalidade de executar uma tarefa específica. Em um alto nível, uma função recebe uma entrada e retorna uma saída.

Funções são geralmente o bloco de códigos ou instruções em um programa que dá ao usuário a capacidade de reutilizar o mesmo código que, em última análise, economiza o uso excessivo de memória, atua como uma economia de tempo e, mais importante, fornece melhor legibilidade do código. Então, basicamente, uma função é uma coleção de instruções que executam alguma tarefa específica e retornam o resultado para o chamador. Uma função também pode executar alguma tarefa específica sem retornar nada.

Abaixo o exemplo de uma função em Go:

```go
func nome_funcao(lista_parametros)(tipo_retorno){
    // corpo da função.....
}
```

A declaração da função contém:

- __func__: É uma palavra-chave na linguagem Go, que é usada para criar uma função.
- __nome_funcao__: É o nome da função.
- __lista_parametros__: Ele contém o nome e o tipo dos parâmetros de função.
- __tipo_retorno__: É opcional e contém os tipos dos valores que a função retorna. Se você estiver usando __tipo_retorno__ em sua função, então é necessário usar uma instrução `return` em sua função.

A Invocação de Função ou a Chamada de Função é feita quando o usuário deseja executar a função. A função precisa ser chamada para usar sua funcionalidade. Como mostrado no exemplo abaixo, temos uma função chamada `area()` com dois parâmetros. Agora chamamos essa função na função principal usando seu nome, ou seja, `area(12, 10)` com dois parâmetros.

```go
package main
import "fmt"
 
func area(length, width int)int{
     
    Ar := length * width
    return Ar
}

func main() {
   fmt.Printf("Area of rectangle is : %d", area(12, 10))
}

// Saída:
// Area of rectangle is : 120
```

> __Note__
> Se as funções com nomes que começam com uma letra maiúscula serão exportadas para outros pacotes. Se o nome da função começar com uma letra minúscula, ela não será exportada para outros pacotes, mas você poderá chamar essa função dentro do mesmo pacote.

### Argumentos

Na linguagem Go, os parâmetros passados para uma função são chamados de __parâmetros reais__, enquanto os parâmetros recebidos por uma função são chamados de __parâmetros formais__.

A linguagem Go oferece suporte a duas maneiras de passar argumentos para sua função:

- __Chamada por valor__: Neste método de passagem de parâmetros, os valores dos parâmetros reais são copiados para os parâmetros formais da função e os dois tipos de parâmetros são armazenados em diferentes locais de memória. Portanto, quaisquer alterações feitas dentro das funções __não são refletidas nos parâmetros reais__ do chamador.

    Exemplo:

    ```go
    package main
    
    import "fmt"
    
    func swap(a, b int) int {
        var o int
        o = a
        a = b
        b = o
        return o
    }
    
    func main() {
        var p int = 10
        var q int = 20
        
        fmt.Printf("p = %d and q = %d", p, q)
        swap(p, q) // chamada por valor
        fmt.Printf("\np = %d and q = %d",p, q)
    }

    // Saída:
    // p = 10 and q = 20
    // p = 10 and q = 20
    ```

- __Chamada por referência__: Ambos os parâmetros reais e formais referem-se aos mesmos locais, de modo que quaisquer alterações feitas dentro da função __são realmente refletidas nos parâmetros reais__ do chamador. Isto é realizado utilizando ponteiros, que passa o endereço de um tipo para a função. A pilha de funções tem uma referência ao objeto original. Portanto, quaisquer modificações no objeto passado modificarão o objeto original.

    Exemplo:

    ```go
    package main
    
    import "fmt"
    
    func swap(a, b *int) int {
        var o int
        o = *a
        *a = *b
        *b = o
        return o
    }
    
    func main() {
        var p int = 10
        var q int = 20

        fmt.Printf("p = %d and q = %d", p, q)
        swap(&p, &q) // chamada por referência
        fmt.Printf("\np = %d and q = %d",p, q)
    }

    // Saída:
    // p = 10 and q = 20
    // p = 20 and q = 10
    ```

### Funções variádicas

TODO:

## Método

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

### Composição

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

- <https://www.geeksforgeeks.org/functions-in-go-language>
- <https://www.geeksforgeeks.org/variadic-functions-in-go>