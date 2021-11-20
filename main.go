package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	inp = ""
	path = ""

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "input",
			Aliases:     []string{"i"},
			Usage: 		"enter input in terminal",
			Destination: &inp,
		},	&cli.StringFlag{
			Name:        "path",
			Aliases:     []string{"p"},
			Usage: 		"enter path to file",
			Destination: &path,
		},
	}

	jump = map[int]int{}

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
	if path != "" {
		withFile(path)
		return nil
	}
	parse(inp)
	return nil
}

func parse(input string) {
	jumpMap(input)
	pos := 0
	mapPos := 0
	mapPoss := make(map[int]int, 0)
	d := make(map[int]int, 0)
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			if _, ok := d[pos]; ok {
				d[pos] += 1
			} else {
				d[pos] = 1
			}
		case '-':
			if d[pos] > 0 {
				d[pos] -= 1
			}
		case '<':
			pos -= 1
		case '>':
			pos += 1
		case '[':
			mapPos++
			mapPoss[mapPos] = i + 1
		case ']':
			if d[pos] == 0 {
				if mapPos > 0 {
					mapPos--
				}
			} else {
				i = mapPoss[mapPos]
			}
		case ',':
		case '.':
			fmt.Print(string(d[pos]))
		}
	}
}

func withFile(filePath string) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	s := ""
	for _, i := range string(dat) {
		switch i {
		case '>','<','.',',','[',']','+','-':
			s += string(i)
		}
	}
	fmt.Println(s)
	parse(s)
}

func jumpMap(s string) {
	var  mapPos = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			mapPos++
		case ']':
			mapPos--
			jump[mapPos] = i
			jump[i] = mapPos
		}
	}
	fmt.Println(jump)
}