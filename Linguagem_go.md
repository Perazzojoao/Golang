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

### for padrão

**Sintaxe:** for i := 0; i < `<var>`; i++ {. . .}

### for range (forEach)

**Sintaxe:** for `<índice>`, `<elemento>` := `range` `<slice>` {. . .}

**Ex:**

```
  for _, site := range sites {
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "-> Carregado com sucesso! StatusCode:", resp.StatusCode)
		} else {
			fmt.Println("Site:", site, "-> Fora do ar! StatusCode:", resp.StatusCode)
		}
	}
```

## Slice (Array)

Slice é um array de tamanho dinâmico, em que a própria linguagem infere com base nos ítens presentes.

### Por de baixo dos panos:

Ao criarmos um slice, a linguagem cria um array de `3 posições por padrão`. Caso ultrapassemos essa capacidade, seu tamanho é `dobrado automaticamente`. Sendo assim, um slice de tamanho 4 contém capacidade para 6 ítens e um slice de tamanho 9 terá capacidade para 12 ítens.

#### Obs: O slice usa como base a capacidade em que foi inicializada para dobrar de capacidade. Assim um slice inicializado com 5 ítens, ao estrapolar seu tamanho, será aumentado para 10 posições.

### Sintaxe:

`<variável>` := `[]type` {var1, var2, . . .}

**Ex:**

```
  nome := []string {"Douglas", "João", "Nathália", "Daniel"}
```

#### Identificar quantidade de ítens:

**len(nome):** retorna a quantidade de ítens no array. -> // 4

**cap(nome):** retorna a capacidade atual do array. -> // 6

#### Funções slice

1.  `append(<slice>, novo_ítem)`: retorna um novo slice com o ítem adicionado.

    **Ex:**

        exemploSlice = append(exemploSlice, "Hello")

## Ler arquivos txt

Existem mais de uma forma de ler arquivos txt. Cada método resulta em diferentes possibilidades de manipulação.

1.  **Ler arquivo inteiro:** `os.ReadFile("<path>")` -> Apenas leitura

2.  **Abrir conexão:** `os.Open("<path>")`

      **Sintaxe:** `file`, `err` := `os.Open("arquivo.txt")`

      **Criar leitor:** `leitor` := `bufio.NewReader(file)`

      **Ler linha por linha:** `linha`, `err` := `leitor.ReadString('<fim_linha>')`

      **Ex:**

        func lerSites() []string {
          file, err := os.Open("sites.txt")

          if err != nil {
            fmt.Println("Ocorreu um erro ->", err)
          }

          leitor := bufio.NewReader(file)

          var lista []string
          for {
            linha, err := leitor.ReadString('\n')

            if err != nil {
              if err == io.EOF {
                break
              }
              fmt.Println("Ocorreu um erro ->", err)
            }
            linha = strings.TrimSpace(linha) -> Método de string para remover espaços em branco
            lista = append(lista, linha)
          }
          file.Close()
          return lista
        }

      #### Obs: Ao abrirmos uma conexão com um arquivo, é boa prática fechar-la ao final: `<file>.Close()`