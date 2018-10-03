package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // para numeros aleatorios
}

func main() {
	vitorias := 0
	perdas := 0
	numeroLetras := rand.Intn(2) + 4 // de 4 ~ 6
	novamente, ganhou := jogarForca(numeroLetras)

	for {
		if ganhou == true {
			vitorias++
			numeroLetras = rand.Intn(11) + 4
		} else {
			perdas++
			numeroLetras = rand.Intn(11) + 4
		}
		if novamente == "y" {
			limparTela()
			fmt.Printf("------------------------\n")
			fmt.Printf("    Placar\n")
			fmt.Printf("  %d: vitorias, %d: perdas\n", vitorias, perdas)
			fmt.Printf("------------------------\n")
			novamente, ganhou = jogarForca(numeroLetras)
		} else if novamente == "n" {
			break
		}
	}
}

func jogarForca(numeroLetras int) (jogarnovamente string, isGanhador bool) {
	estagioMorte := 0
	chutouUmaLetra := false
	ganhou := false
	chute := ""
	letrasChutadas := ""
	novamente := ""
	sublinhado := ""
	novoSublinhado := ""

	fmt.Printf("F O R C A\n")
	fmt.Printf("Digite a palavra secreta: ")
	palavraEscondida, _ := gopass.GetPasswdMasked()
	palavra := string(palavraEscondida)
	fmt.Printf("O tamanho da palavra e: %v\n", len(palavra))

	for {
		desenharForca(estagioMorte)
		if estagioMorte == 7 {
			fmt.Printf("Eita...o enforcado morreu!\n")
			fmt.Printf("A palavra que poderia sava-lo era: %s\n", palavra)
			for {
				fmt.Printf("Jogar novamente? (y/n) \n")
				fmt.Scanln(&novamente)
				isYorN, err := regexp.MatchString("^y|Y|n|N", novamente)
				if err != nil {
					fmt.Printf("Eita aconteceu algo muito errado. Saindo com erro de regex match %v", novamente)
					return
				}
				if isYorN == false {
					fmt.Printf("Voce nao digitou 'y' or 'n'! Tente novamente\n")
				} else if len(novamente) > 1 {
					fmt.Printf("Voce digitou mais de uma letra! Tente novamente\n")
				} else if strings.ToLower(novamente) == "y" {
					return "y", false
				} else {
					return "n", false
				}

			}
		}
		if chutouUmaLetra == false {
			sublinhado = esconderPalavra(len(palavra))
			fmt.Printf("%s\n", sublinhado)
		} else {
			fmt.Printf("%s\n", novoSublinhado)
		}
		fmt.Printf("Chute uma letra: ")
		fmt.Scanln(&chute)

		isALetter, err := regexp.MatchString("^[a-zA-Z]", chute)
		if err != nil {
			limparTela()
			fmt.Printf("Eita aconteceu algo muito errado. Saindo com erro de regex match %v", chute)
			return
		}

		if isALetter == false {
			limparTela()
			fmt.Printf("O que voce digitou nao e uma letra! Tente novamente\n")
		} else if len(chute) > 1 {
			limparTela()
			fmt.Printf("Voce digitou mais de uma letra! Tente novamente\n")
		} else if strings.Contains(letrasChutadas, chute) {
			limparTela()
			fmt.Printf("Voce ja chutou esta letra! Tente novamente\n")
		} else if strings.Contains(palavra, chute) {
			limparTela()
			fmt.Printf("A letra que voce chutou esta na palavra\n")
			letrasChutadas += chute

			if chutouUmaLetra == false {
				sublinhadoAtualizado := revelarSublinhado(palavra, chute, sublinhado)
				novoSublinhado = sublinhadoAtualizado
			} else {
				sublinhadoAtualizado := revelarSublinhado(palavra, chute, novoSublinhado)
				novoSublinhado = sublinhadoAtualizado
			}

			chutouUmaLetra = true
			if novoSublinhado == palavra {
				ganhou = true
			}
			if ganhou == true {
				limparTela()
				fmt.Printf("-= P A R A B E N S =-\n")
				fmt.Printf("Voce ganhou! A palavra era: %s\n", palavra)
				for {
					fmt.Printf("Jogar novamente? (y/n) \n")
					fmt.Scanln(&novamente)
					isYorN, err := regexp.MatchString("^y|Y|n|N", novamente)
					if err != nil {
						fmt.Printf("Eita aconteceu algo muito errado. Saindo com erro de regex match %v", novamente)
						return
					}
					if isYorN == false {
						fmt.Printf("Voce nao digitou 'y' or 'n'! Tente novamente\n")
					} else if len(novamente) > 1 {
						fmt.Printf("Voce digitou mais do que uma letra! Tente novamente\n")
					} else if strings.ToLower(novamente) == "y" {
						return "y", true
					} else {
						return "n", true
					}
				}
			}
		} else {
			limparTela()
			fmt.Printf("A palavra que digitou nao esta na palavra\n")
			estagioMorte++
			letrasChutadas += chute
		}
	}
}

func desenharForca(estagioMorte int) {
	switch estagioMorte {
	case 0:
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 1:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 2:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 3:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 4:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 5:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 6:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 7:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("MORREU|\n")
		fmt.Printf("========\n")
	}
}

func esconderPalavra(tamanhoPalavra int) string {
	sublinhado := ""
	for i := 0; i < tamanhoPalavra; i++ {
		sublinhado += "_"
	}
	return sublinhado
}

func revelarSublinhado(palavra string, chute string, sublinhado string) string {
	novoSublinhado := ""
	for i, r := range sublinhado {
		if c := string(r); c != "_" {
			novoSublinhado += c

		} else {
			var letra = string(palavra[i])
			if chute == letra {
				novoSublinhado += chute
			} else {
				novoSublinhado += "_"
			}
		}
	}
	return novoSublinhado
}

func verificarVitoria(novoSublinhado string, palavra string) bool {
	if novoSublinhado == palavra {
		return true
	}
	return false
}

func limparTela() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
