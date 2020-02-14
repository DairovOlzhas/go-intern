package telnet

import "github.com/urfare/cli"

var {
	flags = []cli.Flag{
		clt.StringFlag(){
			Name: "config",
			Aliases: []string{"c"}
			Destiantion: &cofigF(config)
		}

	}

}

dunc main(){
	app := cli.NewApp()

	app.Flags = flags
	app.Run(op.Args)
	app.Commands = cli.commands{
		&cli.Commands{
			name: "+"
		}
	}
}
