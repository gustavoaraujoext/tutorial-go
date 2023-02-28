# fmt

O pacote `fmt` implementa E/S formatada com funções análogas para C como `printf` e `scanf`. O formato dos 'verbos' é derivado de C, mas são mais simples.

## Verbos

__Geral__:

```txt
%v  - o valor em um formato padrão*, ao imprimir _structs_, o sinalizador "+" (%+v) adiciona nomes de campo.
%#v - uma representação da sintaxe Go do valor.
%T  - uma representação da sintaxe Go do tipo do valor.
%%  - um sinal de porcentagem literal; não consome nenhum valor (imprime somente um %).


* O formato padrão para %v é:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

__Booleano__:

```txt
%t - a palavra verdadeiro ou falso
```

__Integer__:

```txt
%b - base 2.
%c - o caractere representado pelo ponto de código Unicode correspondente.
%d - base 10.
%o - base 8.
%O - base 8 com prefixo 0o.
%q - um caractere entre aspas simples escapado com segurança.
%x - base 16, com letras minúsculas para a-f
%X - base 16, com letras maiúsculas para A-F
%U - formato Unicode: U+1234; o mesmo que "U+%04X"
```

__Ponto Flutuante__:

```txt
%b - notação científica sem decimal com expoente a potência de dois, na forma de strconv. FormatFloat com o formato 'b', por exemplo. -123456p-78.
%e - notação científica, por exemplo -1.234456e+78.
%E - notação científica, por exemplo -1.234456E+78.
%f - ponto decimal, mas sem expoente, ex. 123.456.
%F - sinônimo de %f.
%g %e - para grandes expoentes, %f caso contrário. A precisão é discutida abaixo.
%G %E - para grandes expoentes, %F caso contrário.
%x - notação hexadecimal (com potência decimal de dois expoentes), ex. -0x1.23abcp+20.
%X - notação hexadecimal maiúscula, por exemplo -0X1.23ABCP+20.
```

__String e Slice de bytes__:

```txt
%s - os bytes não interpretados da string ou fatia.
%q - uma string entre aspas duplas escapada com segurança.
%x - base 16, minúsculas, dois caracteres por byte.
%X - base 16, maiúsculas, dois caracteres por byte.
```

__Slice__:

```txt
%p - endereço do 0º elemento na notação de base 16, com 0x inicial.
```

__Ponteiro__:

```txt
%p - notação de base 16, com 0x à esquerda.
%b, %d, %o, %x e %X - também funcionam com ponteiros, formatando o valor exatamente como se fosse um número inteiro.
```

// TODO:

## Referências

- <https://pkg.go.dev/fmt>