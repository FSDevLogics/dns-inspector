package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli/v2" // Importa a versão 2 do pacote de criação de CLI
)

// Create inicializa e configura a nossa aplicação de linha de comando.
// Retorna um ponteiro para a estrutura do aplicativo cli.App.
func Create() *cli.App {
	// Cria uma nova instância vazia do aplicativo
	app := cli.NewApp()
	app.Name = "CLI App"                                     // Nome do aplicativo
	app.Usage = "Busca IPs e Nomes de Servidor na internet"  // Descrição do que ele faz (aparece no help)

	// Definimos as flags (parâmetros opcionais) que o usuário pode passar no terminal.
	// Exemplo de uso no terminal: --host google.com
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",            // Nome da flag
			Value: "parse3d.com.br",  // Valor padrão caso o usuário não passe a flag "--host"
			Usage: "O domínio que você deseja pesquisar", // Descrição para a tela de ajuda
		},
	}

	// Commands define os comandos que o aplicativo aceita (neste caso, 'ip' e 'server')
	app.Commands = []*cli.Command{
		{
			Name:   "ip",                                      // Comando digitado pelo usuário
			Usage:  "Busca os endereços IP de um domínio na internet", // Descrição
			Flags:  flags,                                     // Vincula a flag "--host" a este comando
			Action: searchIPs,                                 // Função que será disparada ao usar este comando
		},
		{
			Name:   "server",                                  // Comando digitado pelo usuário
			Usage:  "Busca os servidores de nome (DNS) de um domínio", // Descrição
			Flags:  flags,                                     // Vincula a mesma flag "--host"
			Action: searchServer,                              // Função que será disparada ao usar este comando
		},
	}

	// Retorna o aplicativo pronto e configurado para o arquivo main.go
	return app
}

// searchIPs é a função responsável por consultar os IPs.
// Ela recebe o contexto (c) do CLI, que contém as flags e argumentos que o usuário digitou.
func searchIPs(c *cli.Context) error {
	// Pega o valor em texto (String) associado à flag "host"
	host := c.String("host")

	// net.LookupIP faz a consulta real na internet buscando os IPs associados ao domínio
	ips, err := net.LookupIP(host)
	if err != nil {
		return err // Se o domínio for inválido ou estiver sem internet, repassa o erro para o main.go
	}

	// O retorno de LookupIP é uma lista de IPs. O 'for' percorre e imprime cada um na tela.
	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil // Retorna nil (nulo) indicando que a função executou com sucesso
}

// searchServer é a função responsável por consultar os servidores DNS.
func searchServer(c *cli.Context) error {
	// Pega o valor da flag "host"
	host := c.String("host")

	// net.LookupNS busca os Name Servers (Servidores de Nome) associados ao domínio
	servers, err := net.LookupNS(host)
	if err != nil {
		return err // Repassa qualquer erro encontrado
	}

	// Percorre a lista de servidores encontrados.
	// Como a variável 'server' é uma struct, acessamos a propriedade '.Host' para pegar o nome
	for _, server := range servers {
		fmt.Println(server.Host)
	}

	return nil // Sucesso
}