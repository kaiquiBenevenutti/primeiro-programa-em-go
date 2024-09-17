package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {

		var resposta string

		fmt.Print("Digite o seu nome: ")
		_, err := fmt.Scanln(&resposta)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			waitForUserInput()
			continue
		}

		fmt.Println("Ola", resposta, "como voce esta?")

		var resposta2 string
		resposta2 = strings.ToLower(resposta2)

		for {
			_, err2 := fmt.Scanln(&resposta2)

			if err2 != nil {
				fmt.Println("Erro ao ler a entrada:", err2)
				waitForUserInput()
				break
			}

			if resposta2 == "bem" {
				fmt.Println("Que otimo!")
				break
			} else if resposta2 == "mau" {
				fmt.Println("Ah, que pena.")
				break
			} else {
				fmt.Println("Não entendi sua resposta, responda com 'bem' ou 'mau'.")
				fmt.Println(resposta, "como voce esta?")
			}
		}

		fmt.Println("Pressione Enter para sair...")
		waitForUserInput()
		return
	}
}

func waitForUserInput() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
