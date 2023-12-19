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

type MyCondition struct {
	variable, operator string
	value int
}

type NormalizedWorkflow struct {
	name string
	condition MyCondition
	leftTrue, rightFalse string
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

func solve(workflows []Workflow, machineParts []MachinePart) (result int) {
	normalizeWorkflows := getNormalizeWorkflows(workflows)

	mins := [4]int{1, 1, 1, 1}
	maxs := [4]int{4000, 4000, 4000, 4000}

	result = traverse(normalizeWorkflows, "in", mins, maxs)

	return
}

type Part [4]int

var indexMap = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

func traverse(workflows []NormalizedWorkflow, name string, mins, maxs Part) int {
	if name == "A" {
		return calculate(mins, maxs)
	} else if name == "R" {
		return 0
	}

	workflow := getNormWorkflowByName(name, workflows)
	leftMins, rightMins := mins, mins
	leftMaxs, rightMaxs := maxs, maxs

	index, _ := indexMap[workflow.condition.variable]
	if workflow.condition.operator == "<" {
		leftMaxs[index] = min(leftMaxs[index], workflow.condition.value - 1)
		rightMins[index] = max(rightMins[index], workflow.condition.value)
	} else if workflow.condition.operator == ">" {
		leftMins[index] = max(leftMins[index], workflow.condition.value + 1)
		rightMaxs[index] = min(rightMaxs[index], workflow.condition.value)
	}

	return traverse(workflows, workflow.leftTrue, leftMins, leftMaxs) + 
		traverse(workflows, workflow.rightFalse, rightMins, rightMaxs)
}

func calculate(mins, maxs Part) int {
	return (maxs[0] - mins[0] + 1) * 
		(maxs[1] - mins[1] + 1) * 
		(maxs[2] - mins[2] + 1) * 
		(maxs[3] - mins[3] + 1)
}

func getNormWorkflowByName(name string, workflows []NormalizedWorkflow) NormalizedWorkflow {
	workflowIndex := slices.IndexFunc(workflows, func (workflow NormalizedWorkflow) bool {
		return workflow.name == name
	})

	return workflows[workflowIndex]
}

func getNormalizeWorkflows(workflows []Workflow) []NormalizedWorkflow {
	result := make([]NormalizedWorkflow, 0)

	for _, workflow := range workflows {
		for index, condition := range workflow.conditions {
			newNormWorklow := NormalizedWorkflow{}
			newNormWorklow.name = workflow.name
			newNormWorklow.condition = MyCondition{
				variable: condition.variable,
				operator: condition.operator,
				value: condition.value,
			}
			newNormWorklow.leftTrue = condition.exit
			newNormWorklow.rightFalse = workflow.lastRule

			if index > 0 {
				newNormWorklow.name = fmt.Sprintf("%s%d", workflow.name, index)
			}

			if index + 1 < len(workflow.conditions) {
				newNormWorklow.rightFalse = fmt.Sprintf("%s%d", workflow.name, index + 1)
			}

			result = append(result, newNormWorklow)
		}
	}

	return result
}
