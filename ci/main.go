package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm"
)

type Example struct {
	name    string
	content string
}

type Examples struct {
	mu       sync.Mutex
	examples []Example
}

func (e *Examples) Add(name, content string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.examples = append(e.examples, Example{
		name:    name,
		content: content,
	})
}

func (e *Examples) GetSorted() []Example {
	e.mu.Lock()
	defer e.mu.Unlock()
	sort.Slice(e.examples, func(i, j int) bool {
		return e.examples[i].name < e.examples[j].name
	})
	return e.examples
}

func (e *Examples) String() string {
	var output string
	for _, example := range e.GetSorted() {
		output += example.content
	}
	return output
}

func main() {
	pterm.Info.Prefix = pterm.Prefix{
		Text:  "LOG",
		Style: pterm.NewStyle(pterm.FgGray),
	}
	pterm.Info.MessageStyle = pterm.NewStyle(pterm.FgDefault)

	do("Running PTerm CI System", 1, func(currentLevel int) {
		do("Generating PUtils Docs", currentLevel, func(currentLevel int) {
			pterm.Info.Println("Getting docs from 'go doc'...")
			goDocOutputBytes, err := exec.Command("go", "doc", "-all", "./putils").Output()
			if err != nil {
				log.Panic(err)
			}
			goDocOutput := string(goDocOutputBytes)
			goDocOutput = strings.Join(strings.Split(goDocOutput, "FUNCTIONS")[1:], "")
			goDocOutputLines := strings.Split(goDocOutput, "\n")

			pterm.Info.Println("Parsing docs...")
			var goDocOutputFiltered []string
			for _, line := range goDocOutputLines {
				if strings.HasPrefix(line, "func") {
					goDocOutputFiltered = append(goDocOutputFiltered, line)
				}
			}

			pterm.Info.Println("Reading README Template")
			putilsTemplateBytes, err := os.ReadFile("./putils/README.template.md")
			if err != nil {
				log.Panic(err)
			}

			pterm.Info.Println("Generating README.md")
			putilsReadme := string(putilsTemplateBytes)
			putilsReadme += "\n## Util Functions\n\n"

			putilsReadme += fmt.Sprintf("```go\n%s\n```\n", strings.Join(goDocOutputFiltered, "\n"))

			pterm.Info.Println("Writing './putils/README.md'")
			os.WriteFile("./putils/README.md", []byte(putilsReadme), 0600)
			pterm.Info.Println("Writing './docs/docs/putils.md'")
			os.WriteFile("./docs/docs/putils.md", []byte(putilsReadme), 0600)
		})

		var allPrinters []string
		do("Geneating printers Table", currentLevel, func(currentLevel int) {
			// get features located in "_examples/*"
			dirs, _ := os.ReadDir("./_examples/")

			for _, dir := range dirs {
				if dir.Name() == "README.md" {
					continue
				}

				// Exclude demo folder
				if dir.Name() == "demo" {
					continue
				}

				allPrinters = append(allPrinters, dir.Name())
			}

			// generate table
			tableContent := "| Feature | Feature | Feature | Feature | Feature |\n| :-------: | :-------: | :-------: | :-------: | :-------: |\n"
			for i, feature := range allPrinters {
				// the table should contain 5 columns. Each cell is a feature.
				// Make multiple rows, if there are more than 4 features.
				// A link to the examples should be included in every cell.
				// Format: "[Example](https://github.com/pterm/pterm/tree/master/_examples/FEATURE)"
				if i%5 == 0 {
					tableContent += "| "
				}
				name := strings.ToUpper(string(feature[0])) + feature[1:]
				name = strings.ReplaceAll(name, "_", " ")
				tableContent += fmt.Sprintf("%s <br/> [(Examples)](https://github.com/pterm/pterm/tree/master/_examples/%s) |", name, feature)
				if (i+1)%5 == 0 {
					tableContent += "\n"
				}

				// fill left over cells with empty strings if in the last row
				if i == len(allPrinters)-1 {
					tableContent += strings.Repeat(" | ", 5-(i+1)%5)
				}
			}

			// read readme
			readme, _ := os.ReadFile("./README.md")

			// replace table in readme
			readmeString := string(readme)
			readmeString = writeBetween("printers", readmeString, "\n"+tableContent+"\n")

			// write readme
			os.WriteFile("./README.md", []byte(readmeString), 0600)
		})

		do("Write printers to website", currentLevel, func(currentLevel int) {
			pterm.Info.Println("Writing printers to website")
			websiteIndex, _ := os.ReadFile("./docs/index.html")
			websiteIndexString := string(websiteIndex)

			// Write as li elements, which contain a link to the example. (https://github.com/pterm/pterm/tree/master/_examples/{name})
			var links []string
			for _, printer := range allPrinters {
				links = append(links, fmt.Sprintf(`<li><a href="https://github.com/pterm/pterm/tree/master/_examples/%s">%s</a></li>`, printer, printer))
			}

			// Replace placeholder with li elements.
			websiteIndexString = writeBetween("printers", websiteIndexString, "\n"+strings.Join(links, "\n")+"\n")

			os.WriteFile("./docs/index.html", []byte(websiteIndexString), 0600)
		})

		var readmeExamples string
		do("Generating Examples", currentLevel, func(currentLevel int) {
			files, _ := os.ReadDir("./_examples/")

			var sectionExamples sync.Map
			var wg sync.WaitGroup

			for _, section := range files {
				if section.Name() == "README.md" {
					continue
				}

				wg.Add(1)
				go func(section os.DirEntry) {
					defer wg.Done()

					sectionContent := generateSectionContent(section)
					sectionExamples.Store(section.Name(), sectionContent)
				}(section)
			}

			pterm.Info.Println("Waiting for all examples to be generated...")
			wg.Wait()

			var keys []string
			sectionExamples.Range(func(key, value any) bool {
				keys = append(keys, key.(string))
				return true
			})

			sort.Strings(keys)

			for _, key := range keys {
				value, _ := sectionExamples.Load(key)
				readmeExamples += value.(string)
			}

			pterm.Info.Println("Writing '/_examples/README.md'")
			examplesReadme, _ := os.ReadFile("./_examples/README.md")
			examplesReadmeContent := string(examplesReadme)
			examplesReadmeContent = writeBetween("examples", examplesReadmeContent, "\n"+readmeExamples+"\n")
			os.WriteFile("./_examples/README.md", []byte(examplesReadmeContent), 0600)
		})

		var newReadmeContent string
		readmeContent, err := os.ReadFile("./README.md")
		if err != nil {
			log.Panic(err)
		}

		do("Counting Unit Tests", currentLevel, func(currentLevel int) {
			unittestTimeout := make(chan string, 1)

			go func() {
				cmd := exec.Command("bash", "-c", "go test -v -p 1 ./...")
				json, _ := cmd.CombinedOutput()
				unitTestCount := fmt.Sprint(strings.Count(string(json), "RUN"))
				pterm.Info.Println("Unit test count:", unitTestCount)
				unittestTimeout <- unitTestCount
			}()

			pterm.Info.Println("Replacing Strings in README")

			newReadmeContent = string(readmeContent)

			select {
			case res := <-unittestTimeout:
				newReadmeContent = writeBetween("unittestcount", newReadmeContent, `<img src="https://img.shields.io/badge/Unit_Tests-`+res+`-magenta?style=flat-square" alt="Forks">`)
				newReadmeContent = writeBetween("unittestcount2", newReadmeContent, "**`"+res+"`**")
			case <-time.After(time.Second * 10):
				pterm.Info.Println("!!! Timeout in counting unit tests")
			}

			newReadmeContent = writeBetween("examples", newReadmeContent, "\n"+readmeExamples+"\n")
		})

		do("Writing README", currentLevel, func(currentLevel int) {
			pterm.Info.Println("Appending examples to root README.md")

			pterm.Info.Println("Writing README")
			err = os.WriteFile("./README.md", []byte(newReadmeContent), 0600)
			if err != nil {
				log.Panic(err)
			}
		})
	})
}

func generateSectionContent(section os.DirEntry) string {
	var sectionExamples string
	examples, _ := os.ReadDir("./_examples/" + section.Name())

	var exampleMap sync.Map
	var wg sync.WaitGroup

	for _, example := range examples {
		if example.Name() == "README.md" {
			continue
		}

		wg.Add(1)
		go func(section os.DirEntry, example os.DirEntry) {
			defer wg.Done()

			dir := section.Name() + "/" + example.Name()
			content := generateExampleContent(dir, section, example)

			pterm.Info.Println("Storing to map: " + section.Name() + "/" + example.Name() + "...")
			exampleMap.Store(example.Name(), content)
			pterm.Info.Println("Stored to map: " + section.Name() + "/" + example.Name())
		}(section, example)
	}

	wg.Wait()

	var keys []string
	exampleMap.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		return true
	})

	sort.Slice(keys, func(i, j int) bool {
		if keys[i] == "demo" {
			return true
		}
		if keys[j] == "demo" {
			return false
		}
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		value, _ := exampleMap.Load(key)
		sectionExamples += value.(string)
	}

	os.WriteFile("./_examples/"+section.Name()+"/README.md", []byte(sectionExamples), 0600)

	return sectionExamples
}
func generateExampleContent(dir string, section os.DirEntry, example os.DirEntry) string {
	var content string

	exampleRenderStart := time.Now()
	animationDataPath := "./_examples/" + dir + "/animation_data.json"
	animationSvgPath := "./_examples/" + dir + "/animation.svg"
	exampleCode, err := os.ReadFile("./_examples/" + dir + "/main.go")
	if err != nil {
		log.Panic(err)
	}

	if fileExists(animationDataPath) {
		pterm.Info.Println("[" + dir + "] 'animation_data.json' already exists. Removing it.")
		err = os.Remove(animationDataPath)
		if err != nil {
			log.Panic(err)
		}
	}
	if fileExists(animationSvgPath) {
		pterm.Info.Println("[" + dir + "] 'animation.svg' already exists. Removing it.")
		err := os.Remove(animationSvgPath)
		if err != nil {
			log.Panic(err)
		}
	}

	pterm.Info.Println("[" + dir + "] Running asciinema")
	execute(`go build -o ./_examples/` + dir + `/bundle ./_examples/` + dir)
	execute(`asciinema rec ` + animationDataPath + ` -c "./_examples/` + dir + `/bundle && sleep 1"`)
	os.Remove("./_examples/" + dir + "/bundle")

	pterm.Info.Println("[" + dir + "] Adding sleep to end of 'animation_data.json'")
	addSleepToEndOfAnimationData(dir, animationDataPath)

	pterm.Info.Println("[" + dir + "] Generating SVG")
	generateSVG(animationDataPath, animationSvgPath)

	pterm.Info.Println("[" + dir + "] Overwriting SVG font")
	overwriteSVGFont(dir, animationSvgPath)

	pterm.Info.Println("[" + dir + "] Generating README")
	readmeString := generateExampleReadme(dir, exampleCode)
	err = os.WriteFile("./_examples/"+dir+"/README.md", []byte(readmeString), 0600)
	if err != nil {
		log.Panic(err)
	}

	pterm.Info.Println("[" + dir + "] Cleaning files")
	os.Remove(animationDataPath)

	content += "### " + section.Name() + "/" + example.Name() + "\n\n"
	content += "![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/" + section.Name() + "/" + example.Name() + "/animation.svg)\n\n"
	content += "<details>\n\n<summary>SHOW SOURCE</summary>\n\n"
	content += "```go\n"
	content += string(exampleCode) + "\n"
	content += "```\n\n"
	content += "</details>\n\n"

	pterm.Info.Println("[" + dir + "] This example took: " + time.Since(exampleRenderStart).String())

	return content
}

func generateSVG(animationDataPath, animationSvgPath string) {
	noCursorFlag := "--no-cursor"
	if strings.Contains(animationDataPath, "interactive") {
		noCursorFlag = ""
	}
	execute(`svg-term --in ` + animationDataPath + ` --out ` + animationSvgPath + " " + noCursorFlag + ` --window true --no-optimize --profile "./ci/terminal-theme.txt" --term "iterm2"`)
}

func addSleepToEndOfAnimationData(dir, animationDataPath string) {
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
}

func overwriteSVGFont(dir, animationSvgPath string) {
	svgContent, err := os.ReadFile(animationSvgPath)
	if err != nil {
		log.Panicf("[%s] %s", dir, err.Error())
	}

	svgContent = []byte(strings.ReplaceAll(string(svgContent), `font-family:`, `font-family:'Courier New',`))
	svgContent = []byte(strings.ReplaceAll(string(svgContent), `font-family="`, `font-family="'Courier New',`))

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
}

func generateExampleReadme(dir string, exampleCode []byte) string {
	readmeString := "# " + dir + "\n\n![Animation](animation.svg)\n\n"
	readmeString += "```go\n"
	readmeString += string(exampleCode)
	readmeString += "\n```\n"

	return readmeString
}
