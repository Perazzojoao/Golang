# Linguagem GO
> REsumo das principais funcionalidades da linguagem Golang.

## Func main ():
**func main():** Função principal à ser executada. Todo o código executável deve estar sujeito à ela.

**Sintaxe:**

``` 
  func main() {
    .
    .
    .
  }
```

## Variáveis
**Criação:** 
  1. `var` <`nome`> `type` -> Criação padrão.

  2. `var` <`nome`> = <`valor`> -> Inferência de tipo.
  
  3. <`nome`> := <`valor`> -> Inferência de var e tipo.

#### Obs: Variáveis inicializadas sem valor assumem um valor padrão para seu próprio tipo.

## Tipos
- **string**

- **int, int8, int16, int32, int64**

- **float32 ou float64**

- **bool**

#### **OBS-1:** adicionar `u` no início do tipo exclui números negativos.
#### **OBS-2:** adicionar `[]` no início do tipo implica em um array do mesmo tipo.

### Valor padrão de cada tipo
- **string:** `" "`

- **int:** `0`

- **float:** `0`

- **bool:** `false`

## print no console
**Comando:** `println()` ou `fmt.Println()`

**Ex:**

    fmt.Println("Site:", site, "-> Carregado com sucesso! StatusCode:", resp.StatusCode)

## scan
**Comando:** `fmt.Scanf("%d", &<var>)` ou `fmt.Scan(&<var>)`

## Funções
**Sintaxe:** func `<nome>`(`<param>` `type`) `type retorno` {. . .}

**Obs:** Deixar o retorno em branco implica em tipo `void`.

### Retornando mais de um valor:
**Sintaxe:** func `<nome>`(`<param>` `type`) (`type1`, `type2`, ...) {return var1, var2, . . .}

**Recebendo os valores:** var1, var2 := retornaDois()

## for
Na linguagem go não existe o comando `while`. Ao invés disso, utilizamos o `for` sem nenhum parâmetro.

**Ex:**

```
  for {

  }
```
Sendo assim, geramos um loop infinito. Para interromper o programa: `os.Exit(<código>)` -> `0` para bem sucedido ou `-1` para mau sucedido.