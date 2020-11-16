package main

import (
	"log"
	"os"
	"fmt"
	"net"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network tool"
	app.Usage = "Query IPs, CNAMEs, MX records and Name servers."

	myFlags := []cli.Flag{
		&cli.StringFlag {
			Name: "host",
			Value: "pranavjoglekarcodes.web.app",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name: "ns",
			Usage: "Looks up the nameservers for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
