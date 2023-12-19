package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type MachinePart struct {
	x, m, a, s int
}

type Condition struct {
	variable string
	operator string
	value int
	exit string
}

type Workflow struct {
	name string
	conditions []Condition 
	lastRule string
}

func main() {
	timeStart := time.Now()
	content, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	workflows, machineParts := parse(content)
	result := solve(workflows, machineParts)

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

func atoi(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		return -1
	}

	return value
}

func parse(content []string) ([]Workflow, []MachinePart) {
	workflows := make([]Workflow, 0)
	machineParts := make([]MachinePart, 0)

	for _, row := range content {
		if row == "" {
			continue
		}

		if row[0] == '{' {
			specStr := row[1:len(row)-1]
			specs := strings.Split(specStr, ",")

			x := atoi(specs[0][2:])
			m := atoi(specs[1][2:])
			a := atoi(specs[2][2:])
			s := atoi(specs[3][2:])

			machineParts = append(machineParts, MachinePart{ x, m, a, s})
		} else {
			braceIndex := strings.Index(row, "{")
			name := row[0:braceIndex]
			specStr := row[braceIndex + 1:len(row) - 1]
			specs := strings.Split(specStr, ",")

			conditions := make([]Condition, len(specs) - 1)

			for index, spec := range specs[0:len(specs) - 1] {
				semicolonIndex := strings.Index(spec, ":")

				conditions[index] = Condition{
					variable: string(spec[0]),
					operator: string(spec[1]),
					value: atoi(spec[2:semicolonIndex]),
					exit: spec[semicolonIndex + 1:],
				}
			}

			workflows = append(workflows, Workflow{
				name: name,
				conditions: conditions,
				lastRule: specs[len(specs) - 1],
			})
		}

	}

	return workflows, machineParts
}

func getWorkflowByName(name string, workflows []Workflow) Workflow {
	workflowIndex := slices.IndexFunc(workflows, func (workflow Workflow) bool {
		return workflow.name == name
	})

	return workflows[workflowIndex]
}

func solve(workflows []Workflow, machineParts []MachinePart) (result int) {
	inWorkflow := getWorkflowByName("in", workflows) 

	accepted := make([]MachinePart, 0)

	for _, part := range machineParts {
		nextWorkflowName := checkWorkflow(part, inWorkflow)

		for ;; {
			curWorkflow := getWorkflowByName(nextWorkflowName, workflows)
			nextWorkflowName = checkWorkflow(part, curWorkflow)
			
			if nextWorkflowName == "A" {
				accepted = append(accepted, part)
				break
			} else if nextWorkflowName == "R" {
				break
			}
		}
	}

	result = calculateResult(accepted)

	return
}

func calculateResult(parts []MachinePart) (result int) {
	for _, part := range parts {
		result += part.x + part.m +part.a + part.s
	}

	return
}

func checkWorkflow(part MachinePart, workflow Workflow) string {
	for _, condition := range workflow.conditions {
		isSated := checkCondition(part, condition)

		if isSated {
			return condition.exit
		}
	}

	return workflow.lastRule
}

func checkCondition(part MachinePart, condition Condition) bool {
	partValue := part.x
	conditionValue := condition.value

	if condition.variable == "x" {
		partValue = part.x
	} else if condition.variable == "m" {
		partValue = part.m
	} else if condition.variable == "a" {
		partValue = part.a
	} else if condition.variable == "s" {
		partValue = part.s
	}

	if condition.operator == ">" {
		return partValue > conditionValue	
	} else if condition.operator == "<" {
		return partValue < conditionValue
	}

	fmt.Println("Should not happen", partValue, condition)
	return false
}
