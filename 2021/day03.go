package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day03", "Relative path to input file.")
var sample = flag.Bool("sample", false, "Use sample input")

// sloppy - clean up to avoid 3 iterations
func main() {
	flag.Parse()
	fileName := *inputFile
	if *sample {
		fileName += ".sample"
	}
	fileName += ".input"

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	zeros := make([]int, len(split[0]))
	ones := make([]int, len(split[0]))

	for _, binNum := range split {
		for pos, bit := range binNum {
			switch bit {
			case '0':
				zeros[pos] += 1
			case '1':
				ones[pos] += 1
			}
		}
	}

	gammaBits := make([]string, len(split[0]))
	epsilonBits := make([]string, len(split[0]))
	for i := range zeros {
		if zeros[i] > ones[i] {
			gammaBits[i] = "0"
			epsilonBits[i] = "1"
		} else {
			gammaBits[i] = "1"
			epsilonBits[i] = "0"
		}
	}
	gammaStr := strings.Join(gammaBits, "")
	epsilonStr := strings.Join(epsilonBits, "")

	gammaRate, err := strconv.ParseInt(gammaStr, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonRate, err := strconv.ParseInt(epsilonStr, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("gamma bits:", strings.Join(gammaBits, ""), "epsilon bits:", strings.Join(epsilonBits, ""))
	fmt.Println("gamma rate:", gammaRate, "epsilon rate:", epsilonRate, "power consumption:", gammaRate*epsilonRate)

	var mcbs, lcbs []string
	var o2candidates []string
	prevO2candidates := split
	for idx := range split[0] {
		if len(prevO2candidates) == 1 {
			break
		}
		var zeros, ones int
		for _, binNum := range prevO2candidates {
			switch binNum[idx] {
			case '0':
				zeros++
			case '1':
				ones++
			}
		}
		if zeros > ones {
			mcbs = append(mcbs, "0")
		} else {
			mcbs = append(mcbs, "1")
		}
		fmt.Println("mcbs:", mcbs)
		for _, v := range split {
			if strings.HasPrefix(v, strings.Join(mcbs, "")) {
				o2candidates = append(o2candidates, v)
			}
		}
		fmt.Println(o2candidates)
		prevO2candidates = o2candidates
		o2candidates = []string{}
	}

	prevCO2candidates := split
	var co2candidates []string
	for idx := range split[0] {
		if len(prevCO2candidates) == 1 {
			break
		}
		var zeros, ones int
		for _, binNum := range prevCO2candidates {
			switch binNum[idx] {
			case '0':
				zeros++
			case '1':
				ones++
			}
		}
		if zeros <= ones {
			lcbs = append(lcbs, "0")
		} else {
			lcbs = append(lcbs, "1")
		}
		fmt.Println("lcbs:", lcbs)
		for _, v := range split {
			if strings.HasPrefix(v, strings.Join(lcbs, "")) {
				co2candidates = append(co2candidates, v)
			}
		}
		fmt.Println(co2candidates)
		prevCO2candidates = co2candidates
		co2candidates = []string{}
	}

	fmt.Println("Oxygen generator rating:", prevO2candidates[0], "CO2 scrubber rating:", prevCO2candidates[0])

	o2Rate, err := strconv.ParseInt(prevO2candidates[0], 2, 64)
	if err != nil {
		panic(err)
	}
	co2Rate, err := strconv.ParseInt(prevCO2candidates[0], 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Oxygen generator rating:", o2Rate, "CO2 scrubber rating:", co2Rate, "Life support rating:", o2Rate*co2Rate)
}
