package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate vai retornar a aplicação de cli para ser executada
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação cli"
	app.Usage = "Busca Ips e Nomes de servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "dev-araujo.com.br", // valor default
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca servidores",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
