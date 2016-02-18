# UniformRandom
Uniform Random Number generator using Linear Congruential formula

# Install Instructions

First, [download the Golang compiler](https://golang.org/dl/), at least version 1.5.
Next follow the [installation instructions](https://golang.org/doc/install). Note, that to successfully compile this source, you need to establish a Go workspace. This is a directory that contains the subdirectories `src`, `bin`, and `pkg`.
Then, set the environment variable GOVENDOREXPERIMENT to 1. (You should also set the GOPATH variable as per the installation instructions.

    GOVENDOREXPERIMENT=1
    GOPATH=$HOME/go-workspace # done during the installation process

Now, recall that this directory needs to be located in the Go Path at a specific location so that the compiler can find the dependency packages. One option (not preferred) is to move this directory to `$GOPATH/github.com/RobbieMcKinstry/UniformRandom`.

Instead, you should use the built in Go tool, `go get`. Run `go get github.com/RobbieMcKinstry/UniformRandom`. This will successfully pull down a copy of this repo into the correct location in the Go path. `cd` into that directory.

`cd $GOPATH/github.com/RobbieMcKinstry/UniformRandom`.

Then, you can successfully compile and execute the source. `go run main.go`, or alternatively, `go build main.go`

# Usage

To compile the binary, `cd` into this directory and run `go build main.go`. This will create an executable called `main`.

Run `./main`.

You should see the following:

    NAME:
       main - A new cli application

    USAGE:
       main [global options] command [command options] [arguments...]
       
    VERSION:
       0.0.0
       
    COMMANDS:
       kolmogorov, k	run the Kolmogov-Smirnov Test for uniformity
       runs, r		runs Test for independence
       autocorrelation, a	autocorrelation test for independence
       chi, c		Chi-Square Frequency Test for uniformity
       help, h		Shows a list of commands or help for one command
       
    GLOBAL OPTIONS:
       -n "10000"		number of elements to generate. Defaults to 10000. NOTE: Kolgorov only uses the first 100
       -f "output.txt"	File to which the output is written, defaults to output.txt
       --help, -h		show help
       --version, -v	print the version

This should be self-explanatory. The binary uses a Git-like subcommand pattern. As you can see, the commands are `kolmogorov`, `runs`, `autocorrelation`, `chi`, and `help`. To see the subcommands, run the command with no other arguments.

    `$ ./main runs`

     NAME:
       main runs - runs Test for independence

    USAGE:
       main runs command [command options] [arguments...]

    COMMANDS:
       stdlib	uses the Golang standard library generator, found in 'golang.org/pkg/math/rand'
       lipschultz	uses the generator seeded as a = 101427, c = 321, m = 1 << 16
       randu	uses the generator seeded as a = 65539, c = 0, m = 1 << 31
       help, h	Shows a list of commands or help for one command
       
    OPTIONS:
       --help, -h	show help
 
This is true for all commands except for `help`, which does not have a subcommand.

Now, you can specify which generator you wish to use. After specifying the generator, the program will output the test statistic.

Lastly, note that you can provide options to adjust the output file, and adjust the size of the dataset, by passing them as command line options prior to specifying a subcommand. 

Here's an example command that runs a runs test on a dataset with a sample size of 50,000, and outputs the results to `experiment.txt`.

    $ ./main -n 50000 -f experiment.txt runs stdlib
