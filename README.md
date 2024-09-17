# Meu Primeiro Projeto em Go

Este é o meu primeiro projeto em Go. O programa é uma aplicação simples que pede ao usuário seu nome e, em seguida, solicita uma resposta sobre como ele está. O programa valida a resposta e continua perguntando até que uma resposta válida seja fornecida.

## Como Funciona

1. O programa solicita ao usuário que digite seu nome.
2. Após receber o nome, o programa pergunta como o usuário está.
3. O usuário deve responder com "bem" ou "mal". 
4. Se a resposta for "bem", o programa responde com "Que ótimo!" e encerra.
5. Se a resposta for "mal", o programa responde com "Ah, que pena." e encerra.
6. Se a resposta não for "bem" ou "mal", o programa informa que não entendeu e pergunta novamente como o usuário está.

## Código

```go
package main

import (
	"fmt"
	"strings"
)

func main() {

	var resposta string

	fmt.Print("Digite o seu nome: ")
	_, err := fmt.Scanln(&resposta)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}

	fmt.Println("Olá", resposta, "como você está?")

	var resposta2 string
	resposta2 = strings.ToLower(resposta2)

	for {
		_, err2 := fmt.Scanln(&resposta2)

		if err2 != nil {
			fmt.Println("Erro ao ler a entrada:", err2)
			return
		}

		if resposta2 == "bem" {
			fmt.Println("Que ótimo!")
			break
		} else if resposta2 == "mal" {
			fmt.Println("Ah, que pena.")
			break
		} else {
			fmt.Println("Não entendi sua resposta, responda com 'bem' ou 'mal'.")
			fmt.Println(resposta, "como você está?")
		}
	}
}
```

## Como Executar

1. Certifique-se de ter o Go instalado em seu sistema.
2. Compile o código para gerar o executável `main.exe` usando o comando: `go build -o main.exe`.
3. Navegue até a pasta onde o `main.exe` está salvo.
4. Para executar o programa, você pode:
   - **No Terminal:** Use o comando `./main.exe` (ou `main.exe` no Windows).
   - **Na Interface Gráfica:** Dê um duplo clique no `main.exe` na pasta do projeto.
