package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Search APIs and Server Names"
	app.Usage = "Searching for apis and servers names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for ips on the internet",
			Flags:  flags,
			Action: searchIPs,
		},

		{
			Name:   "serverNames",
			Usage:  "Search for server names",
			Flags:  flags,
			Action: searchSNs,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchSNs(c *cli.Context) {
	host := c.String("host")

	serverNames, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, serverName := range serverNames {
		fmt.Println(serverName)
	}
}
