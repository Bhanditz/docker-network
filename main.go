package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/docker/docker-network/namespace"
)

var (
	networkCommand = cli.Command{
		Name:      "network",
		ShortName: "n",
		Usage:     "Manage networks",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "Shows list of created networks",
			},
			{
				Name:  "add",
				Usage: "Add network",
			},
			{
				Name:  "del",
				Usage: "Delete network",
			},
			{
				Name:  "show",
				Usage: "Shows info about network",
			},
		},
	}
	epCommand = cli.Command{
		Name:      "endpoint",
		ShortName: "ep",
		Usage:     "Manage endpoints in network",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "Shows list of endpoinds in network",
			},
			{
				Name:  "add",
				Usage: "Add endpoint to network",
			},
			{
				Name:  "del",
				Usage: "Delete endpoint from network",
			},
			{
				Name:  "show",
				Usage: "Shows info about endpoint",
			},
		},
	}
	nsCommand = cli.Command{
		Name:      "namespace",
		ShortName: "ns",
		Usage:     "Manage network namespaces",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "List of network namespaces which belongs to docker-network",
			},
			{
				Name:  "add",
				Usage: "Add new network namespace",
				Action: func(c *cli.Context) {
					if _, err := namespace.New(c.Args().First()); err != nil {
						log.Fatal(err)
					}
				},
			},
			{
				Name:  "del",
				Usage: "Delete network namespace",
				Action: func(c *cli.Context) {
					ns := &namespace.Namespace{Path: c.Args().First()}
					if err := ns.Delete(); err != nil {
						log.Fatal(err)
					}
				},
			},
			{
				Name:  "join",
				Usage: "Join endpoint to specified namespace (this can be docker-network namespace or path)",
				Action: func(c *cli.Context) {
					ns := &namespace.Namespace{Path: c.Args().First()}
					if err := ns.Join(); err != nil {
						log.Fatal(err)
					}
				},
			},
			{
				Name:  "exec",
				Usage: "Execute command in namespace",
				Action: func(c *cli.Context) {
					ns := &namespace.Namespace{Path: c.Args().First()}
					tail := c.Args().Tail()
					if len(tail) == 0 {
						log.Fatal("Not enough arguments to call exec")
					}
					cmd := exec.Command(tail[0], tail[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Stdin = os.Stdin
					if err := ns.Exec(cmd); err != nil {
						log.Fatal(err)
					}
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-network"
	app.Usage = "Create and manage networks"
	app.Action = func(c *cli.Context) {
		println(app.Usage)
	}
	app.Commands = []cli.Command{
		networkCommand,
		epCommand,
		nsCommand,
	}
	app.Run(os.Args)
}
