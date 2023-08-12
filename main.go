
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func shuffle(data []*Sentence) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func printEvaluation(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sentences", "Seconds", "Accuracy"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

var commandTrain = cli.Command{
	Name:  "train",
	Usage: "Train a parsing model by easy-first algorithm",
	Description: `
Train a parsing model by easy-first algorithm.
`,
	Action: doTrain,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "train-filename"},
		cli.StringFlag{Name: "dev-filename"},
		cli.StringFlag{Name: "model-filename"},
		cli.IntFlag{Name: "max-iter", Value: 10},
	},
}

var commandEval = cli.Command{
	Name:  "eval",
	Usage: "Evaluate a parsing model by easy-first algorithm",
	Description: `
Evaluate a parsing model by easy-first algorithm.
`,
	Action: doEval,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "test-filename"},
		cli.StringFlag{Name: "model-filename"},
	},
}