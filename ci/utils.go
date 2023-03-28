package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pterm/pterm"
)

func do(title string, currentLevel int, f func(currentLevel int)) {
	pterm.DefaultSection.WithBottomPadding(0).WithLevel(currentLevel).Println(title)
	f(currentLevel + 1)
}

func writeBetween(name string, original string, insertText string) string {
	beforeRegex := regexp.MustCompile(`(?ms).*<!-- ` + name + `:start -->`)
	afterRegex := regexp.MustCompile(`(?ms)<!-- ` + name + `:end -->.*`)
	before := beforeRegex.FindAllString(original, 1)[0]
	after := afterRegex.FindAllString(original, 1)[0]

	ret := before
	ret += insertText
	ret += after

	return ret
}

func execute(command string) {
	cmd := exec.Command("bash", "-c", command)
	pterm.Info.Println("Running: " + cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Panicf("Fail Running [%s] with output [%s] and status [%s]", cmd.String(), output, err)
	}
	pterm.FgGray.Println(strings.Repeat("-", 80))
	pterm.Println(string(output))
	pterm.FgGray.Println(strings.Repeat("-", 80))
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getLinesFromFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
