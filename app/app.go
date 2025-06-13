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

	/*
	 * @desc   Define o comando principal da aplicação.
	 *         Este comando é chamado quando o usuário executa a aplicação sem argumentos adicionais.
	 *
	 * @usage  go run main.go ip --parametro valor
	 *
	 * @example
	 *         go run main.go ip (retorna o IP do google.com)
	 *         go run main.go ip --host amazon.com.br (retorna o IP do amazon.com.br)
	 */
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
					Usage: "Especifica o host para buscar o IP",
				},
			},
			Action: buscarIps,
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
		fmt.Println("IP encontrado:", ip)
	}
}
