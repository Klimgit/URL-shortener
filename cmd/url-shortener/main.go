package main

import (
	"URL-shortener/internal/config"
	"fmt"
	_ "log/slog"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Printf("%+v\n\n", cfg)

	//TODO init logger: slog

	//TODO: init storage: postgress

	//TODO: init router: chi, chi-render

	//TODO: run server
}
