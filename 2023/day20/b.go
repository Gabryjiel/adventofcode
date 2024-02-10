package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()
	content, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	modules := parse(content)
	modules = addInputsToModules(modules)
	result := solve(modules)

	fmt.Println("Result:", result)
	fmt.Println("Time:", float64(time.Since(timeStart).Microseconds()) / 1000, "ms")
}

func openFile(name string) ([]string, error) {
	file, err := os.Open(name)	
	defer file.Close()

	content := make([]string, 0)

	if (err != nil) {
		return nil, fmt.Errorf("File %s not found", name)
	}
	
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		content = append(content, sc.Text())
	}

	return content, nil
}

type Module struct {
	name string
	prefix string 
	outputs []string
	state bool
	inputs map[string]bool
}

func parse(content []string) ([]Module) {
	modules := make([]Module, len(content))

	for rowIndex, row := range content {
		arrowIndex := strings.Index(row, "->") - 1
		prefix := string(row[0])
		name := row[1:arrowIndex]

		if prefix != "%" && prefix != "&" {
			name = row[:arrowIndex]
		}

		outputsStr := row[arrowIndex + 4:]
		outputs := strings.Split(outputsStr, ", ")

		modules[rowIndex] = Module{
			name,
			prefix,
			outputs,
			false,
			make(map[string]bool),
		}
	}

	return modules
}

func addInputsToModules(modules []Module) []Module {
	for index, module := range modules {
		if module.prefix != "&" {
			continue
		}

		inputs := make([]string, 0)

		for _, searchModule := range modules {
			isIn := false

			for _, output := range searchModule.outputs {
				if output == module.name {
					isIn = true
					break
				}
			}

			if isIn {
				inputs = append(inputs, searchModule.name)
			}
		}

		for _, input := range inputs {
			modules[index].inputs[input] = false
		}
	}

	return modules
}

type Pulse struct {
	source string
	state bool
	target string
}

func solve(modules []Module) int{
	searchModuleNames := filterModuleNamesBasedOnOutput(modules, "zg")
	fmt.Println(searchModuleNames)
	numsA := make([]int, len(searchModuleNames))
	numsB := make([]int, len(searchModuleNames))
	isCycle := make([]bool, len(searchModuleNames))

	for i := 0; i < 10_000; i++ {
		if isEveryTrue(isCycle) {
			break
		}

		for index, name := range searchModuleNames {
			modIdx := getModuleIndexByName(name, modules)
			if modules[modIdx].state == true {
				fmt.Println("IN")
				if numsA[index] == 0 {
					numsA[index] = i
				} else {
					isCycle[index] = true
					numsB[index] = i
				}
			}
		}


		pulses := make([]Pulse, 1)
		pulses[0] = Pulse{
			source: "button",
			state: false,
			target: "broadcaster",
		}


		for len(pulses) > 0 {
			head := pulses[0]
			pulses = pulses[1:]


			newPulses := processPulse(head, modules)
			pulses = append(pulses, newPulses...)
		}
	}

	fmt.Println(numsA)
	fmt.Println(numsB)

	return -1
}

func isEveryTrue(array []bool) bool {
	for _, element := range array {
		if element == false {
			return false
		}
	}

	return true
}

func filterModuleNamesBasedOnOutput(modules []Module, name string) []string {
	result := make([]string, 0)

	for _, module := range modules {
		for _, output := range module.outputs {
			isIn := false

			if output == name {
				isIn = true
			}

			if isIn {
				result = append(result, module.name)
			}
		}
	}

	return result
}

func processPulse(pulse Pulse, modules []Module) []Pulse {
	resultPulses := make([]Pulse, 0)
	targetModuleIndex := getModuleIndexByName(pulse.target, modules)

	if targetModuleIndex == -1 {
		return resultPulses
	}

	targetModule := &modules[targetModuleIndex]

	if targetModule.prefix == "%" {
		if pulse.state == false {
			stateToSend := true

			if targetModule.state {
				stateToSend = false
			}

			targetModule.state = !targetModule.state
			resultPulses = createPulses(targetModule.outputs, pulse.target, stateToSend)
		}
	} else if targetModule.prefix == "&" {
		_, ok := targetModule.inputs[pulse.source]

		if ok == false {
			targetModule.inputs[pulse.source] = false
		} else {
			targetModule.inputs[pulse.source] = pulse.state
		}

		stateToSend := !areAllInputsHigh(*targetModule)
		resultPulses = createPulses(targetModule.outputs, pulse.target, stateToSend)
	} else if targetModule.prefix == "b" {
		resultPulses = createPulses(targetModule.outputs, pulse.target, false)
	} else {
		fmt.Println("Unknown prefix", targetModule.prefix)
	}

	return resultPulses
}

func createPulses(outputs []string, source string, state bool) []Pulse {
	result := make([]Pulse, len(outputs))

	for index, output := range outputs {
		result[index] = Pulse{
			source: source,
			state: state,
			target: output,
		}
	}

	return result
}

func getModuleIndexByName(name string, modules []Module) int {
	return slices.IndexFunc(modules, func (module Module) bool {
		return module.name == name
	})
}

func areAllInputsHigh(module Module) bool {
	result := true 

	for _, val := range module.inputs {
		if val == false {
			result = false
		}
	}

	return result
}
