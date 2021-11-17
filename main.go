package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	input = ""

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "input",
			Aliases:     []string{"i"},
			Usage: 		"enter input in terminal",
			Destination: &input,
		},
	}

)

func main(){
	app := cli.NewApp()
	app.Commands = cli.Commands{
		&cli.Command{
			Name: "run",
			Action: start,
			Flags: flags,
			Usage: "run script",
		},
	}
	app.Run(os.Args)
}

func start(c *cli.Context) error {
	pos := 0
	mapPos := 0
	mapPoss := make(map[int]int, 0)
	d := make(map[int]int64, 0)
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			if _, ok := d[pos]; ok {
				d[pos] += 1
			} else {
				d[pos] = 1
			}
		case '-':
			d[pos] -= 1
		case '<':
			pos -= 1
		case '>':
			pos += 1
		case '[':
			mapPos ++
			mapPoss[mapPos] = i
		case ']':
			if d[pos] == 0 {
				mapPos --
				continue
			} else {
				i = mapPoss[mapPos]
			}
		case ',':
		case '.':
			fmt.Print(string(d[pos]))
		}
	}
	return nil
}