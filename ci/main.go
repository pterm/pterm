package main

import (
	"bufio"
	"fmt"
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

	putilsTemplateBytes, err := os.ReadFile("./putils/README.template.md")
	if err != nil {
		log.Panic(err)
	}

	putilsReadme := string(putilsTemplateBytes)
	putilsReadme += "\n## Util Functions\n\n"

	putilsReadme += fmt.Sprintf("```go\n%s\n```\n", strings.Join(goDocOutputFiltered, "\n"))

	os.WriteFile("./putils/README.md", []byte(putilsReadme), 0600)
	os.WriteFile("./docs/docs/putils.md", []byte(putilsReadme), 0600)

	log.Output(1, "## Generating Examples")
	files, _ := os.ReadDir("./_examples/")
	var readmeExamples string
	for _, section := range files {
		var sectionExamples string
		log.Output(2, "Section: "+section.Name())
		examples, _ := os.ReadDir("./_examples/" + section.Name())

		for _, example := range examples {
			processFile(section.Name() + "/" + example.Name())
			log.Output(2, "## Generating readme for example: "+example.Name())
			exampleCode, err := os.ReadFile("./_examples/" + section.Name() + "/" + example.Name() + "/main.go")
			if err != nil {
				log.Panic(err)
			}

			sectionExamples += "### " + section.Name() + "/" + example.Name() + "\n\n"
			sectionExamples += "![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/" + section.Name() + "/" + example.Name() + "/animation.svg)\n\n"
			sectionExamples += "<details>\n\n<summary>SHOW SOURCE</summary>\n\n"
			sectionExamples += "```go\n"
			sectionExamples += string(exampleCode) + "\n"
			sectionExamples += "```\n\n"
			sectionExamples += "</details>\n\n"

			readmeExamples += sectionExamples
		}
		os.WriteFile("./_examples/"+section.Name()+"/README.md", []byte(sectionExamples), 0600)
	}

	log.Output(3, "### Generating examples README")
	examplesReadme, err := os.ReadFile("./_examples/README.md")
	examplesReadmeContent := string(examplesReadme)
	examplesReadmeContent = writeBetween("examples", examplesReadmeContent, "\n"+readmeExamples+"\n")
	os.WriteFile("./_examples/README.md", []byte(examplesReadmeContent), 0600)

	log.Output(3, "### Appending examples to root README.md")

	readmeContent, err := os.ReadFile("./README.md")
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
	err = os.WriteFile("./README.md", []byte(newReadmeContent), 0600)
	if err != nil {
		log.Panic(err)
	}

	log.Output(4, "### Writing readme to pterm.sh")
	err = os.WriteFile("./docs/README.md", []byte(newReadmeContent), 0600)
	if err != nil {
		log.Panic(err)
	}
}

func processFile(dir string) {
	log.Output(3, "### ['"+dir+"'] Generating animations for example")
	animationDataPath := "./_examples/" + dir + "/animation_data.json"
	animationSvgPath := "./_examples/" + dir + "/animation.svg"
	exampleCode, err := os.ReadFile("./_examples/" + dir + "/main.go")
	if err != nil {
		log.Panic(err)
	}

	if fileExists(animationDataPath) {
		log.Output(4, "#### ['"+dir+"']  animation_data.json already exists. Removing it.")
		err = os.Remove(animationDataPath)
		if err != nil {
			log.Panic(err)
		}
	}
	if fileExists(animationSvgPath) {
		log.Output(4, "#### ['"+dir+"']  animation.svg already exists. Removing it.")
		err := os.Remove(animationSvgPath)
		if err != nil {
			log.Panic(err)
		}
	}

	log.Output(4, "#### ['"+dir+"'] Running asciinema")
	execute(`asciinema rec ` + animationDataPath + ` -c "go run ./_examples/` + dir + `"`)

	log.Output(4, "#### ['"+dir+"']  Adding sleep to end of animation_data.json")
	animationDataLines := getLinesFromFile(animationDataPath)
	animationDataLastLine := animationDataLines[len(animationDataLines)-1]
	re := regexp.MustCompile(`\[\d[^,]*`).FindAllString(animationDataLastLine, 1)[0]
	lastTime, _ := strconv.ParseFloat(strings.ReplaceAll(re, "[", ""), 64)
	sleepString := `[` + strconv.FormatFloat(lastTime+5, 'f', 6, 64) + `, "o", "\nRestarting animation...\n"]`
	animationDataFile, err := os.OpenFile(animationDataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Panicf("[%s] %s", dir, err.Error())
	}
	defer animationDataFile.Close()
	_, err = animationDataFile.WriteString(sleepString)
	if err != nil {
		log.Panicf("[%s] %s", dir, err.Error())
	}

	log.Output(4, "#### ['"+dir+"']  Generating SVG")
	execute(`svg-term --in ` + animationDataPath + ` --out ` + animationSvgPath + ` --no-cursor --window true --no-optimize --profile "./ci/terminal-theme.txt" --term "iterm2"`)

	log.Output(4, "#### ['"+dir+"']  Overwriting SVG font")

	svgContent, err := os.ReadFile(animationSvgPath)
	if err != nil {
		log.Panicf("[%s] %s", dir, err.Error())
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
	os.WriteFile(animationSvgPath, svgContent, 0600)

	log.Output(4, "#### ['"+dir+"']  Generating README.md")
	readmeString := "# " + dir + "\n\n![Animation](animation.svg)\n\n"
	readmeString += "```go\n"
	readmeString += string(exampleCode)
	readmeString += "\n```\n"
	err = os.WriteFile("./_examples/"+dir+"/README.md", []byte(readmeString), 0600)
	if err != nil {
		log.Panic(err)
	}

	log.Output(4, "#### ['"+dir+"']  Adding example to global example list")

	log.Output(4, "#### ['"+dir+"']  Cleaning files")
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
