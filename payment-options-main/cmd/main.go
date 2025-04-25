package main

import (
	"fmt"
	"payment-options/config"
	"payment-options/internal/http"
	"payment-options/internal/repository"
	"payment-options/internal/usecase"
	"runtime"

	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.LoadConfig()

	// Memastikan semua core CPU digunakan
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("GOMAXPROCS set to", runtime.GOMAXPROCS(0))

	// Inisialisasi Echo
	e := echo.New()

	// Setup dependency injection
	repo := repository.NewPaymentRepo()
	uc := usecase.NewPaymentUsecase(repo)
	http.NewPaymentHandler(e, uc)

	e.Logger.Fatal(e.Start(cfg.ListenPort))
}
