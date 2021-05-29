package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Output(1, "## Generating PUtils docs")

	goDocOutputBytes, err := exec.Command("go", "doc", "-all", "./putils").Output()
	if err != nil {
		log.Panic(err)
	}

	goDocOutput := string(goDocOutputBytes)
	goDocOutput = strings.Join(strings.Split(goDocOutput, "FUNCTIONS")[1:], "")

	goDocOutputLines := strings.Split(goDocOutput, "\n")
	var goDocOutputFiltered []string
	for _, line := range goDocOutputLines {
		if strings.HasPrefix(line, "func") {
			goDocOutputFiltered = append(goDocOutputFiltered, line)
		}
	}
	goDocOutput = strings.Join(goDocOutputFiltered, "\n")

	goDocOutput = fmt.Sprintf("```go\n%s\n```\n", goDocOutput)
	goDocOutput = `# PUtils - PTerm Utils

This package contains some utility functions, to get you started with PTerm even faster!  

## Util Functions

` + goDocOutput

	ioutil.WriteFile("./putils/README.md", []byte(goDocOutput), 0600)

	log.Output(1, "## Generating Examples")
	files, err := ioutil.ReadDir("./_examples/")
	if err != nil {
		log.Panic(err)
	}

	var readmeExamples string

	for _, f := range files {
		processFile(f)
	}

	for _, f := range files {
		exampleCode, err := ioutil.ReadFile("./_examples/" + f.Name() + "/main.go")
		if err != nil {
			log.Panic(err)
		}

		readmeExamples += "### " + f.Name() + "\n\n"
		readmeExamples += "![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/" + f.Name() + "/animation.svg)\n\n"
		readmeExamples += "<details>\n\n<summary>SHOW SOURCE</summary>\n\n"
		readmeExamples += "```go\n"
		readmeExamples += string(exampleCode) + "\n"
		readmeExamples += "```\n\n"
		readmeExamples += "</details>\n\n"
	}

	log.Output(3, "### Appending examples to root README.md")

	readmeContent, err := ioutil.ReadFile("./README.md")
	if err != nil {
		log.Panic(err)
	}

	var newReadmeContent string

	log.Output(3, "### Counting unit tests...")

	unittestTimeout := make(chan string, 1)

	go func() {
		cmd := exec.Command("bash", "-c", "go test -v -p 1 ./...")
		json, _ := cmd.CombinedOutput()
		unitTestCount := fmt.Sprint(strings.Count(string(json), "RUN"))
		log.Output(4, "### Unit test count: "+unitTestCount)
		unittestTimeout <- unitTestCount
	}()

	log.Output(4, "#### Replacing strings in readme")

	newReadmeContent = string(readmeContent)

	select {
	case res := <-unittestTimeout:
		newReadmeContent = writeBetween("unittestcount", newReadmeContent, `<img src="https://img.shields.io/badge/Unit_Tests-`+res+`-magenta?style=flat-square" alt="Forks">`)
		newReadmeContent = writeBetween("unittestcount2", newReadmeContent, "**`"+res+"`**")
	case <-time.After(time.Second * 10):
		log.Output(4, "Timeout in counting unit tests!")
	}

	newReadmeContent = writeBetween("examples", newReadmeContent, "\n"+readmeExamples+"\n")

	log.Output(4, "### Writing readme")
	err = ioutil.WriteFile("./README.md", []byte(newReadmeContent), 0600)
	if err != nil {
		log.Panic(err)
	}

	log.Output(4, "### Writing readme to pterm.sh")
	err = ioutil.WriteFile("./docs/README.md", []byte(newReadmeContent), 0600)
	if err != nil {
		log.Panic(err)
	}
}

func processFile(f os.FileInfo) {
	log.Output(3, "### ['"+f.Name()+"'] Generating animations for example")
	animationDataPath := "./_examples/" + f.Name() + "/animation_data.json"
	animationSvgPath := "./_examples/" + f.Name() + "/animation.svg"
	exampleCode, err := ioutil.ReadFile("./_examples/" + f.Name() + "/main.go")
	if err != nil {
		log.Panic(err)
	}

	if fileExists(animationDataPath) {
		log.Output(4, "#### ['"+f.Name()+"']  animation_data.json already exists. Removing it.")
		err = os.Remove(animationDataPath)
		if err != nil {
			log.Panic(err)
		}
	}
	if fileExists(animationSvgPath) {
		log.Output(4, "#### ['"+f.Name()+"']  animation.svg already exists. Removing it.")
		err := os.Remove(animationSvgPath)
		if err != nil {
			log.Panic(err)
		}
	}

	log.Output(4, "#### ['"+f.Name()+"'] Running asciinema")
	execute(`asciinema rec ` + animationDataPath + ` -c "go run ./_examples/` + f.Name() + `"`)

	log.Output(4, "#### ['"+f.Name()+"']  Adding sleep to end of animation_data.json")
	animationDataLines := getLinesFromFile(animationDataPath)
	animationDataLastLine := animationDataLines[len(animationDataLines)-1]
	re := regexp.MustCompile(`\[\d[^,]*`).FindAllString(animationDataLastLine, 1)[0]
	lastTime, _ := strconv.ParseFloat(strings.ReplaceAll(re, "[", ""), 10)
	sleepString := `[` + strconv.FormatFloat(lastTime+5, 'f', 6, 64) + `, "o", "\nRestarting animation...\n"]`
	animationDataFile, err := os.OpenFile(animationDataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Panicf("[%s] %s", f.Name(), err.Error())
	}
	defer animationDataFile.Close()
	_, err = animationDataFile.WriteString(sleepString)
	if err != nil {
		log.Panicf("[%s] %s", f.Name(), err.Error())
	}

	log.Output(4, "#### ['"+f.Name()+"']  Generating SVG")
	execute(`svg-term --in ` + animationDataPath + ` --out ` + animationSvgPath + ` --no-cursor --window true --no-optimize --profile "./ci/terminal-theme.txt" --term "iterm2"`)

	log.Output(4, "#### ['"+f.Name()+"']  Overwriting SVG font")

	svgContent, err := ioutil.ReadFile(animationSvgPath)
	if err != nil {
		log.Panicf("[%s] %s", f.Name(), err.Error())
	}

	svgContent = []byte(strings.ReplaceAll(string(svgContent), `font-family:`, `font-family:'JetBrainsMono',`))
	svgContent = []byte(strings.ReplaceAll(string(svgContent), `font-family="`, `font-family="'JetBrainsMono',`))

	svgContent = []byte(strings.Replace(string(svgContent), "<style>", `<style>.e, .g, .f {
	font-family:
			'Courier New'
			Roboto,
			Helvetica,
			Arial,
			sans-serif,
			'Apple Color Emoji',
			'Segoe UI Emoji' !important;
}`, 1))

	os.Remove(animationSvgPath)
	ioutil.WriteFile(animationSvgPath, svgContent, 0600)

	log.Output(4, "#### ['"+f.Name()+"']  Generating README.md")
	readmeString := "# " + f.Name() + "\n\n![Animation](animation.svg)\n\n"
	readmeString += "```go\n"
	readmeString += string(exampleCode)
	readmeString += "\n```\n"
	err = ioutil.WriteFile("./_examples/"+f.Name()+"/README.md", []byte(readmeString), 0600)
	if err != nil {
		log.Panic(err)
	}

	log.Output(4, "#### ['"+f.Name()+"']  Adding example to global example list")

	log.Output(4, "#### ['"+f.Name()+"']  Cleaning files")
	os.Remove(animationDataPath)

	// wg.Done()
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
	log.Output(1, "Running: "+cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Panicf("Fail Running [%s] with output [%s] and status [%s]", cmd.String(), output, err)
	}
	log.Output(1, fmt.Sprintf("Finish Running [%s] with output [%s]", cmd.String(), output))
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
