package app

import "github.com/urfave/cli"

/*
Gerar cria uma nova aplicação de linha de comando usando o pacote urfave/cli.
*/

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Busca IP's e nomes de servidor na internet"

	return app
}
