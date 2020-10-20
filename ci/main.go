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
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("## Generating Examples")
	files, err := ioutil.ReadDir("./_examples/")
	if err != nil {
		log.Fatal(err)
	}

	var readmeExamples string

	for _, f := range files {
		wg.Add(1)
		go processFile(f)
	}

	wg.Wait()

	for _, f := range files {
		exampleCode, err := ioutil.ReadFile("./_examples/" + f.Name() + "/main.go")
		if err != nil {
			log.Fatal(err)
		}

		readmeExamples += "### " + f.Name() + "\n\n"
		readmeExamples += "![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/" + f.Name() + "/animation.svg)\n\n"
		readmeExamples += "<details>\n\n<summary>SHOW SOURCE</summary>\n\n"
		readmeExamples += "```go\n"
		readmeExamples += string(exampleCode) + "\n"
		readmeExamples += "```\n\n"
		readmeExamples += "</details>\n\n"
	}

	fmt.Println("### Appending examples to root README.md")

	readmeContent, err := ioutil.ReadFile("./README.md")
	if err != nil {
		log.Fatal(err)
	}

	var newReadmeContent string
	var unitTestCountBytes []byte

	cmd := exec.Command("bash", "-c", "go test -v ./... | grep -c RUN")
	unitTestCountBytes, err = cmd.Output()
	if err != nil {
		log.Println(err)
	}

	unitTestCount := strings.ReplaceAll(string(unitTestCountBytes), "\n", "")

	beforeUnitTestBadgeRegex := regexp.MustCompile(`(?ms).*<!-- unittestcount:start -->`)
	afterUnitTestBadgeRegex := regexp.MustCompile(`(?ms)<!-- unittestcount:end -->.*`)

	beforeUnitTestBadge := beforeUnitTestBadgeRegex.FindAllString(string(readmeContent), 1)[0]
	afterUnitTestBadge := afterUnitTestBadgeRegex.FindAllString(string(readmeContent), 1)[0]

	newReadmeContent = beforeUnitTestBadge + "\n"
	newReadmeContent += `<img src="https://img.shields.io/badge/Unit_Tests-` + unitTestCount + `-brightgreen?style=flat-square" alt="Forks">`
	newReadmeContent += afterUnitTestBadge + "\n"

	beforeExamplesRegex := regexp.MustCompile(`(?ms).*<!-- examples:start -->`)
	afterExamplesRegex := regexp.MustCompile(`(?ms)<!-- examples:end -->.*`)

	beforeExamples := beforeExamplesRegex.FindAllString(newReadmeContent, 1)[0]
	afterExamples := afterExamplesRegex.FindAllString(newReadmeContent, 1)[0]

	newReadmeContent = beforeExamples + "\n"
	newReadmeContent += readmeExamples
	newReadmeContent += afterExamples + "\n"

	err = ioutil.WriteFile("./README.md", []byte(newReadmeContent), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(f os.FileInfo) {
	fmt.Println("### Generating animations for example '" + f.Name() + "'")
	animationDataPath := "./_examples/" + f.Name() + "/animation_data.json"
	animationSvgPath := "./_examples/" + f.Name() + "/animation.svg"
	exampleCode, err := ioutil.ReadFile("./_examples/" + f.Name() + "/main.go")
	if err != nil {
		log.Fatal(err)
	}

	if fileExists(animationDataPath) {
		fmt.Println("#### animation_data.json already exists. Removing it.")
		err = os.Remove(animationDataPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	if fileExists(animationSvgPath) {
		fmt.Println("#### animation.svg already exists. Removing it.")
		err := os.Remove(animationSvgPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("#### Running asciinema")
	execute(`asciinema rec ` + animationDataPath + ` -c "go run ./_examples/` + f.Name() + `"`)

	fmt.Println("#### Adding sleep to end of animation_data.json")
	animationDataLines := getLinesFromFile(animationDataPath)
	animationDataLastLine := animationDataLines[len(animationDataLines)-1]
	re := regexp.MustCompile(`\[\d[^,]*`).FindAllString(animationDataLastLine, 1)[0]
	lastTime, _ := strconv.ParseFloat(strings.ReplaceAll(re, "[", ""), 10)
	sleepString := `[` + strconv.FormatFloat(lastTime+5, 'f', 6, 64) + `, "o", "\r\nrestarting...\r\n"]`
	animationDataFile, err := os.OpenFile(animationDataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	defer animationDataFile.Close()
	_, err = animationDataFile.WriteString(sleepString)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("#### Generating SVG")
	execute(`svg-term --in ` + animationDataPath + ` --out ` + animationSvgPath + ` --no-cursor --window true --no-optimize --profile "./ci/terminal-theme.txt" --term "iterm2"`)

	fmt.Println("#### Generating README.md")
	readmeString := "# " + f.Name() + "\n\n![Animation](animation.svg)\n\n"
	readmeString += "```go\n"
	readmeString += string(exampleCode)
	readmeString += "\n```\n"
	err = ioutil.WriteFile("./_examples/"+f.Name()+"/README.md", []byte(readmeString), 0600)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("#### Adding example to global example list")

	fmt.Println("#### Cleaning files")
	os.Remove(animationDataPath)

	wg.Done()
}

func execute(command string) {
	cmd := exec.Command("bash", "-c", command)
	fmt.Println("Running: " + cmd.String())
	err := cmd.Run()
	cmd.Stderr = os.Stderr
	if err != nil {
		log.Fatal(err)
	}
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
