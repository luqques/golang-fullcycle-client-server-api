package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func StartServer() {
	fmt.Println("Iniciando servidor...")
	defer fmt.Println("Encerrando servidor...")

	http.HandleFunc("/cotacao", cotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cotacao, err := buscaCotacao(ctx)
	if err != nil {
		log.Fatalln("Erro ao buscar cotação:", err)
		http.Error(w, "Erro ao buscar cotação: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cotacao)
}

func buscaCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cotacaoJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	print(string(cotacaoJson))

	var cotacao Cotacao
	err = json.Unmarshal(cotacaoJson, &cotacao)

	return &cotacao, err
}
