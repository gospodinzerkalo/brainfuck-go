package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	inp = ""
	path = ""
	inputPath = ""

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
		},&cli.StringFlag{
			Name:        "inputfile",
			Aliases:     []string{"if"},
			Usage: 		"enter path to input file",
			Destination: &inputPath,
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
	var res string
	if path != "" {
		res = withFile(path, inputPath)
		fmt.Println(res)
		return nil
	}
	res = parse(inp, "")
	fmt.Println(res)
	return nil
}

func parse(input string, inputF string) string {
	jumpMap(input)
	pos := 0
	res := ""
	inputPos := 0
	d := make(map[int]uint8, 0)
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
			if pos == 0 {
				panic("negative memory index")
			}
			pos -= 1
		case '>':
			pos += 1
		case '[':
			if d[pos] == 0 {
				i = jump[i]
				continue
			}
		case ']':
			if d[pos] == 0 {
				continue
			} else {
				i = jump[i]
			}
		case ',':
			d[pos] = inputF[inputPos]
			inputPos++
		case '.':
			//fmt.Print(string(d[pos]))
			res += string(d[pos])
		}
	}
	return res
}

func withFile(filePath, inputPath string) string {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var inputDat []byte
	if inputPath != "" {
		inputDat, err = os.ReadFile(inputPath)
		if err != nil {
			panic(err)
		}
	}
	s := ""
	for _, i := range string(dat) {
		switch i {
		case '>','<','.',',','[',']','+','-':
			s += string(i)
		}
	}
	return parse(s, string(inputDat))
}

func jumpMap(s string) {
	var  mapPos = 0
	mapPoss := map[int]int{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			mapPos++
			mapPoss[mapPos] = i
		case ']':
			jump[mapPoss[mapPos]] = i
			jump[i] = mapPoss[mapPos]
			mapPos--
		}
	}
}