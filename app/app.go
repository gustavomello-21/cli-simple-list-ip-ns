package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "command line apllication"
	app.Usage = "Search ips and server names"

	flag := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Usage: "enjoei.com.br",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flag,
			Action: searchIps,
		},
		{
			Name:   "ns",
			Usage:  "Busca servidor por ip",
			Flags:  flag,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) error {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func searchServer(c *cli.Context) error {
	ip := c.String("host")
	ns, err := net.LookupNS(ip)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ns {
		fmt.Println(s.Host)
	}
	return nil

}
