package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var (
	k = flag.String("k", ":", "It is replacement key header.") //key select
	f = flag.Bool("f", false, "It is file option flag.")       // file mode
	e = flag.Bool("e", false, "exec result par line.")         // exec mode
	g = flag.Bool("g", false, "async option.")                 // async mode
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage of this:
		-k  : It is replacement key header.
		-f  : It is file option flag.
		-e  : It exec option. result par line.
		-g  : It process in parallel.`)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments. Please 3 args")
		os.Exit(1)
	}

	os.Exit(run(args[0], args[1]))
}

// getData returns embedded data.
func getData(name string) (string, error) {
	c, ioe := ioutil.ReadFile(name)
	if ioe != nil {
		return "", ioe
	}
	return string(c), nil
}

// output outputs str on os.Stdout
func output(str string) {
	fmt.Fprintln(os.Stdout, str)
}

// outputDo does command
func outputDo(str string) {
	cmds := strings.Split(str, " ")
	var err error
	var res []byte
	if len(cmds) == 1 {
		res, err = exec.Command(str).Output()
	} else {
		res, err = exec.Command(cmds[0], cmds[1:]...).Output()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	output(string(res))
}

// replaceTemplate replaces template.
func replaceTemplate(templ string, r []string) string {
	n := len(r)
	for {
		key := *k + strconv.Itoa(n)
		templ = strings.Replace(templ, key, r[n-1], -1)
		if n-1 == 0 {
			break
		}
		n--
	}
	return templ
}

func run(template string, csvName string) int {
	var templ string
	var err error
	if *f {
		templ, err = getData(template)
	} else {
		templ = template
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%s", err)
		return 1
	}

	fp, err := os.Open(csvName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "No Such File:"+csvName)
		return 1
	}
	defer fp.Close()

	reader := csv.NewReader(fp)

	var wg sync.WaitGroup

	// function
	var doFunc func(str string)
	if *e {
		doFunc = outputDo
	} else {
		doFunc = output
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}

		out := replaceTemplate(templ, row)

		if *g {
			wg.Add(1)
			go func(out string) {
				defer wg.Done()
				doFunc(out)
			}(out)
		} else {
			doFunc(out)
		}
	}
	wg.Wait()

	return 0
}
