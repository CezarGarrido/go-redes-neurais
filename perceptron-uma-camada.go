package main

import (
    "fmt"
    "math"
)

func main() {

    // Operador AND
    var entradas = [][]int{
        []int{0, 0},
        []int{0, 1},
        []int{1, 0},
        []int{1, 1},
    }
    var saidas = []int{0, 0, 0, 1}

    // Operador OR
    /*var entradas = [][]int{
          []int{0, 0},
          []int{0, 1},
          []int{1, 0},
          []int{1, 1},
      }
      var saidas = []int{0, 1, 1, 1}
    */

    //operador XOR, entra em loop
    /*
       var entradas = [][]int{
               []int{0, 0},
               []int{0, 1},
               []int{1, 0},
               []int{1, 1},
           }
       var saidas = []int{0, 1, 1, 0}
    */

    var pesos = []float64{0.0, 0.0}

    var taxaAprendizagem float64 = 0.1

    treinar(entradas, pesos, saidas, taxaAprendizagem)

    fmt.Println("Rede neural treinada")

    fmt.Println(calculaSaida(entradas[0], pesos))
    fmt.Println(calculaSaida(entradas[1], pesos))
    fmt.Println(calculaSaida(entradas[2], pesos))
    fmt.Println(calculaSaida(entradas[3], pesos))

}

func soma(entradas []int, pesos []float64) float64 {
    var s float64 = 0

    for _, e := range entradas {
        for _, p := range pesos {
            s += float64(e) * p
        }
    }
    return s
}

func stepFunction(soma float64) int {
    if soma >= 1.0 {
        return 1
    }
    return 0
}

func calculaSaida(entradas []int, pesos []float64) int {
    return stepFunction(soma(entradas, pesos))
}

func treinar(entradas [][]int, pesos []float64, saidas []int, taxaAprendizagem float64) {
    var erroTotal int = 1

    for erroTotal != 0 {
        erroTotal = 0
        for i, _ := range saidas {
            saidaCalculada := calculaSaida(entradas[i], pesos)

            erro := math.Abs(float64(saidas[i]) - float64(saidaCalculada))

            erroTotal += int(erro)

            for j, _ := range pesos {
                pesos[j] = pesos[j] + (taxaAprendizagem * float64(entradas[i][j]) * erro)
                fmt.Println("Peso atualizado:", pesos[j])
            }

        }
        fmt.Println("Total de erros:", erroTotal)
    }

}
