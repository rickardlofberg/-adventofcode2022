package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/rickardlofberg/adventofcode2024/pkg/util"
	"github.com/urfave/cli/v2"
)

func main() {
	var day int
	var part int

	app := &cli.App{
		Name:  "Advent of code 2024",
		Usage: "Solutions to the problems posed by advent of code 2021!",
		Action: func(c *cli.Context) error {
			if day < 1 || day > 25 {
				return errors.New("day need to be set to between 1 and 25")
			}
			if part < 1 || part > 2 {
				return errors.New("part can only be set to 1 or 2")
			}

			answer, err := util.SolveProblem(day, part)
			if err != nil {
				return err
			}
			fmt.Printf("The answer to day %d, part %d is: %d \n", day, part, answer)

			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "day",
				Usage:       "Which day to solve (1-25)",
				Required:    true,
				Aliases:     []string{"d"},
				Destination: &day,
			},
			&cli.IntFlag{
				Name:        "part",
				Usage:       "Which part to solve (1 or 2)",
				Required:    true,
				Aliases:     []string{"p"},
				Destination: &part,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
