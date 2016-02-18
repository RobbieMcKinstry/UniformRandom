package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/codegangsta/cli"

	"github.com/RobbieMcKinstry/UniformRandom/RandGenerator"
	"github.com/RobbieMcKinstry/UniformRandom/StatTests"
)

func main() {

	app := cli.NewApp()

	var (
		size       int
		outputFile string
		gen        RandGenerator.Generator
		ds         *RandGenerator.Dataset
	)

	const (
		SEED = 122949823
	)

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "n",
			Value:       10000,
			Usage:       "number of elements to generate. Defaults to 10000. NOTE: Kolgorov only uses the first 100",
			Destination: &size,
		}, cli.StringFlag{
			Name:        "f",
			Value:       "output.txt",
			Usage:       "File to which the output is written, defaults to output.txt",
			Destination: &outputFile,
		},
	}

	// Specifies how to interact with the command line tool.
	// See https://github.com/codegangsta/cli for explanation
	app.Commands = []cli.Command{
		{
			Name:    "kolmogorov",
			Aliases: []string{"k"},
			Usage:   "run the Kolmogov-Smirnov Test for uniformity",
			Action: func(c *cli.Context) {
				println("Please specify a subcommand. See -h for a list of commands.")
				cli.ShowSubcommandHelp(c)
			},
			Subcommands: []cli.Command{
				{
					Name:  "stdlib",
					Usage: "uses the Golang standard library generator, found in 'golang.org/pkg/math/rand'",
					Action: func(c *cli.Context) {
						gen = rand.New(rand.NewSource(SEED))
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("Standard Library Kolmogorov Test Statistic:\t%f\n", StatTests.KolmogorovSmirnov(ds))
					},
				}, {
					Name:  "lipschultz",
					Usage: "uses the generator seeded as a = 101427, c = 321, m = 1 << 16",
					Action: func(c *cli.Context) {
						gen = RandGenerator.LipschultzGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("Lipschultz Kolmogorov Test Statistic:\t%f\n", StatTests.KolmogorovSmirnov(ds))
					},
				}, {
					Name:  "randu",
					Usage: "uses the generator seeded as a = 65539, c = 0, m = 1 << 31",
					Action: func(c *cli.Context) {
						gen = RandGenerator.NewRANDUGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("RANDU Kolmogorov Test Statistic:\t%f\n", StatTests.KolmogorovSmirnov(ds))
					},
				},
			},
		}, {
			Name:    "runs",
			Aliases: []string{"r"},
			Usage:   "runs Test for independence",
			Action: func(c *cli.Context) {
				cli.ShowSubcommandHelp(c)
			},
			Subcommands: []cli.Command{
				{
					Name:  "stdlib",
					Usage: "uses the Golang standard library generator, found in 'golang.org/pkg/math/rand'",
					Action: func(c *cli.Context) {
						gen = rand.New(rand.NewSource(SEED))
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("Standard Library Run Test Statistic:\t%f\n", StatTests.Run(ds, StatTests.DIFFERENCE))
					},
				}, {
					Name:  "lipschultz",
					Usage: "uses the generator seeded as a = 101427, c = 321, m = 1 << 16",
					Action: func(c *cli.Context) {
						gen = RandGenerator.LipschultzGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("Lipschultz Run Test Statistic:\t%f\n", StatTests.Run(ds, StatTests.DIFFERENCE))
					},
				}, {
					Name:  "randu",
					Usage: "uses the generator seeded as a = 65539, c = 0, m = 1 << 31",
					Action: func(c *cli.Context) {
						gen = RandGenerator.NewRANDUGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("RANDU Run Test Statistic:\t%f\n", StatTests.Run(ds, StatTests.DIFFERENCE))
					},
				},
			},
		}, {
			Name:    "autocorrelation",
			Aliases: []string{"a"},
			Usage:   "autocorrelation test for independence",
			Action: func(c *cli.Context) {
				cli.ShowSubcommandHelp(c)
			},
			Subcommands: []cli.Command{
				{
					Name:  "stdlib",
					Usage: "uses the Golang standard library generator, found in 'golang.org/pkg/math/rand'",
					Action: func(c *cli.Context) {
						gen = rand.New(rand.NewSource(SEED))
						ds = RandGenerator.NewDataset(gen, size)
						mu, variance := StatTests.Autocorrelation(ds, 1, 3)
						fmt.Printf("Standard Library Autocorrelation Test Statistic:\t%f\n", mu/variance)

					},
				}, {
					Name:  "lipschultz",
					Usage: "uses the generator seeded as a = 101427, c = 321, m = 1 << 16",
					Action: func(c *cli.Context) {
						gen = RandGenerator.LipschultzGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						mu, variance := StatTests.Autocorrelation(ds, 1, 3)
						fmt.Printf("Lipschultz Autocorrelation Test Statistic:\t%f\n", mu/variance)
					},
				}, {
					Name:  "randu",
					Usage: "uses the generator seeded as a = 65539, c = 0, m = 1 << 31",
					Action: func(c *cli.Context) {
						gen = RandGenerator.NewRANDUGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						mu, variance := StatTests.Autocorrelation(ds, 1, 3)
						fmt.Printf("RANDU Autocorrelation Test Statistic:\t%f\n", mu/variance)
					},
				},
			},
		}, {
			Name:    "chi",
			Aliases: []string{"c"},
			Usage:   "Chi-Square Frequency Test for uniformity",
			Action: func(c *cli.Context) {
				cli.ShowSubcommandHelp(c)
			},
			Subcommands: []cli.Command{
				{
					Name:  "stdlib",
					Usage: "uses the Golang standard library generator, found in 'golang.org/pkg/math/rand'",
					Action: func(c *cli.Context) {
						gen = rand.New(rand.NewSource(SEED))
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("Standard Library Chi Square Test Statistic:\t%f\n", StatTests.ChiSquared(ds))
					},
				}, {
					Name:  "lipschultz",
					Usage: "uses the generator seeded as a = 101427, c = 321, m = 1 << 16",
					Action: func(c *cli.Context) {
						gen = RandGenerator.LipschultzGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("Lipschultz Chi Square Test Statistic:\t%f\n", StatTests.ChiSquared(ds))
					},
				}, {
					Name:  "randu",
					Usage: "uses the generator seeded as a = 65539, c = 0, m = 1 << 31",
					Action: func(c *cli.Context) {
						gen = RandGenerator.NewRANDUGen()
						gen.Seed(SEED)
						ds = RandGenerator.NewDataset(gen, size)
						fmt.Printf("RANDU Chi Square Test Statistic:\t%f\n", StatTests.ChiSquared(ds))

					},
				},
			},
		},
	}

	app.Run(os.Args)

	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}

	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := 0; i < ds.Len(); i++ {
		val := ds.Get(i)
		w.WriteString(fmt.Sprintf("%f\n", val))
	}
}
