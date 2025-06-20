package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

/*
* Gerar cria uma nova aplicação de linha de comando usando o pacote urfave/cli.
 */
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Busca IP's e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
			Usage: "Especifica o host para buscar o IP",
		},
	}

	/*
	 * @desc   Define o comando principal da aplicação.
	 *         Este comando é chamado quando o usuário executa a aplicação sem argumentos adicionais.
	 *
	 * @usage  go run main.go ips --parametro valor
	 *
	 * @example
	 *         go run main.go ips (retorna o(s) IP(s) do google.com)
	 *         go run main.go ips --host amazon.com.br (retorna o(S) IP(s) do amazon.com.br)
	 *
	 *         go run main.go servidores (retorna o(s) servidor(s) do google.com)
	 *         go run main.go servidores --host amazon.com.br (retorna o(s) servidor(es) do amazon.com.br)
	 */
	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		}, {
			Name:   "servidores",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)

	if erro != nil {

		log.Fatal("Erro ao buscar IPs:", erro)
	}

	for _, ip := range ips {
		fmt.Println("IP(s) encontrado(s):", ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host) //name server

	if erro != nil {
		log.Fatal("Erro ao buscar IPs:", erro)
	}

	for _, servidor := range servidores {
		fmt.Println("Servidor(es) encontrado:", servidor)
	}
}
