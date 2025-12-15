package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc/utils"
)

func parseBinaryString(s string) int64 {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Machine struct {
	Indicator int64
	Buttons   []int64
	Joltage   []int
}

func (m *Machine) ParseIndicators(indicators []string) {
	s := indicators[0]
	s = strings.ReplaceAll(s, ".", "0")
	s = strings.ReplaceAll(s, "#", "1")
	m.Indicator = parseBinaryString(s)
}

func (m Machine) String() string {
	bs := ""
	for i, b := range m.Buttons {
		bs += fmt.Sprintf("(%04b)", b)
		if i < len(m.Buttons)-1 {
			bs += " "
		}
	}
	return fmt.Sprintf("[%08b] %s {%v}", m.Indicator, bs, m.Joltage)
}

func (m *Machine) ParseButtons(buttons [][]string, totalIndicators int) {
	buttonConfigs := []int64{}
	for _, b := range buttons {
		b := b[1:][0]
		numbers := strings.Split(b, ",")
		var buttonConfig int64
		for _, n := range numbers {
			singleButton := ""
			for i := range totalIndicators {
				value := toInt(n)
				if value == i {
					singleButton += "1"
				} else {
					singleButton += "0"
				}
			}
			singleButtonInt := parseBinaryString(singleButton)
			buttonConfig = buttonConfig | singleButtonInt
		}
		buttonConfigs = append(buttonConfigs, buttonConfig)
	}
	m.Buttons = buttonConfigs
}

func (m *Machine) PressButton(index int) {
	m.Indicator = m.Indicator ^ m.Buttons[index]
}

func (m *Machine) ParseJoltages(joltage []string) {
	joltages := strings.Split(joltage[0], ",")
	joltageList := []int{}
	for _, j := range joltages {
		joltageList = append(joltageList, toInt(j))
	}
	m.Joltage = joltageList
}

func GetMachines() []Machine {
	machines := []Machine{}
	re := regexp.MustCompile(`[\[\(\{](.[^\s]*)[\]\)\}]`)
	lines := utils.GetLines(10)
	for _, line := range lines {
		m := Machine{}
		submatches := re.FindAllStringSubmatch(line, -1)
		indicators := submatches[0][1:]
		m.ParseIndicators(indicators)
		m.ParseButtons(submatches[1:len(submatches)-1], len(indicators[0]))
		m.ParseJoltages(submatches[len(submatches)-1][1:])
		machines = append(machines, m)
	}
	return machines
}

func main() {
	machines := GetMachines()
	for _, m := range machines {
		fmt.Println(m)
	}
}
