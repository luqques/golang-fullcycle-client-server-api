package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
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

	db, err := sql.Open("sqlite", "cotacoes.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:", err)
	}
	defer db.Close()
	criarTabelaCotacoes(db)

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()

		if r.URL.Path != "/cotacao" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cotacao, err := buscarCotacao(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = salvarCotacao(db, cotacao)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(cotacao)
	})

	http.ListenAndServe(":8080", nil)
}

func criarTabelaCotacoes(db *sql.DB) {
	sqlTabelaCotacoes := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"code" TEXT,
		"codein" TEXT,
		"name" TEXT,
		"high" TEXT,
		"low" TEXT,
		"varBid" TEXT,
		"pctChange" TEXT,
		"bid" TEXT,
		"ask" TEXT,
		"timestamp" TEXT,
		"createDate" TEXT
	);`
	_, err := db.Exec(sqlTabelaCotacoes)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
}

func buscarCotacao(ctx context.Context) (*Cotacao, error) {
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

func salvarCotacao(db *sql.DB, cotacao *Cotacao) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO cotacoes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		cotacao.Usdbrl.Code,
		cotacao.Usdbrl.Codein,
		cotacao.Usdbrl.Name,
		cotacao.Usdbrl.High,
		cotacao.Usdbrl.Low,
		cotacao.Usdbrl.VarBid,
		cotacao.Usdbrl.PctChange,
		cotacao.Usdbrl.Bid,
		cotacao.Usdbrl.Ask,
		cotacao.Usdbrl.Timestamp,
		cotacao.Usdbrl.CreateDate,
	)
	if err != nil {
		return err
	}

	return nil
}
