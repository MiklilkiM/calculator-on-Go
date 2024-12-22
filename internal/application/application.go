package application

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/MiklilkiM/calculator-on-go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}
	return &Config{Addr: addr}
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

// запрос на вычисление.
type Request struct {
	Expression string `json:"expression"`
}

// успешный ответ с результатом.
type ResponseSuccess struct {
	Result float64 `json:"result"`
}

// ответ с ошибкой.
type ResponseError struct {
	Error string `json:"error"`
}

// CalcHandler обрабатывает POST-запросы на вычисление выражений.
func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ResponseError{Error: "Expression is not valid"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ResponseError{Error: "Internal server error"})
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseSuccess{Result: result})
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Printf("Server is running on port %s\n", a.config.Addr)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
