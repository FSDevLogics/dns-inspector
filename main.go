package main

import (
	"cli_app/app" // Importa o nosso pacote interno (ajuste 'cli_app' se o nome do seu módulo for diferente)
	"log"
	"os"
)

func main() {
	// Chama a nossa função Create() para montar e receber toda a configuração do app
	application := app.Create()

	// O método Run() pega o aplicativo configurado e o executa.
	// O os.Args repassa para o CLI tudo o que foi digitado no terminal (ex: ["ip", "--host", "google.com"]).
	// Se alguma função Action (como searchIPs) retornar um erro, ele entra no bloco 'if'.
	if err := application.Run(os.Args); err != nil {
		// log.Fatal exibe o erro na tela e encerra o programa abruptamente de forma segura
		log.Fatal(err)
	}
}