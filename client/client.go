package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	fmt.Println("Iniciando cliente...")
	defer fmt.Println("Encerrando cliente...")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Resposta inesperada do servidor:", resp.Status)
		return
	}

	var cotacao Cotacao
	cotacaoJson, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta do servidor:", err)
		return
	}

	err = json.Unmarshal(cotacaoJson, &cotacao)
	if err != nil {
		fmt.Println("Erro ao deserializar cotação:", err)
		return
	}
	fmt.Println("Cotação recebida:", cotacao.Usdbrl.Bid)

	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", cotacao.Usdbrl.Bid))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
}
