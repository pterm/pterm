package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm"
)

type Examples struct {
	sync.Mutex
	m map[string]string
}

func (e *Examples) Add(name, content string) {
	e.Lock()
	defer e.Unlock()
	e.m[name] = content
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

		var readmeExamples string
		do("Generating Examples", currentLevel, func(currentLevel int) {
			files, _ := os.ReadDir("./_examples/")
			for _, section := range files {
				if section.Name() == "README.md" {
					continue
				}
				var sectionExamples string
				do("Section: "+section.Name(), currentLevel, func(currentLevel int) {
					examples, _ := os.ReadDir("./_examples/" + section.Name())

					for _, example := range examples {
						if example.Name() == "README.md" {
							continue
						}
						dir := section.Name() + "/" + example.Name()
						do("Generating Animations for Example: "+dir, currentLevel, func(currentLevel int) {
							exampleRenderStart := time.Now()
							animationDataPath := "./_examples/" + dir + "/animation_data.json"
							animationSvgPath := "./_examples/" + dir + "/animation.svg"
							exampleCode, err := os.ReadFile("./_examples/" + dir + "/main.go")
							if err != nil {
								log.Panic(err)
							}

							if fileExists(animationDataPath) {
								pterm.Info.Println("'animation_data.json' already exists. Removing it.")
								err = os.Remove(animationDataPath)
								if err != nil {
									log.Panic(err)
								}
							}
							if fileExists(animationSvgPath) {
								pterm.Info.Println("'animation.svg' already exists. Removing it.")
								err := os.Remove(animationSvgPath)
								if err != nil {
									log.Panic(err)
								}
							}

							pterm.Info.Println("Running asciinema")
							execute(`asciinema rec ` + animationDataPath + ` -c "go run ./_examples/` + dir + `"`)

							pterm.Info.Println("Adding sleep to end of 'animation_data.json'")
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

							pterm.Info.Println("['" + dir + "']  Generating SVG")
							execute(`svg-term --in ` + animationDataPath + ` --out ` + animationSvgPath + ` --no-cursor --window true --no-optimize --profile "./ci/terminal-theme.txt" --term "iterm2"`)

							pterm.Info.Println("Overwriting SVG font")

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

							pterm.Info.Println("Generating README")
							readmeString := "# " + dir + "\n\n![Animation](animation.svg)\n\n"
							readmeString += "```go\n"
							readmeString += string(exampleCode)
							readmeString += "\n```\n"
							err = os.WriteFile("./_examples/"+dir+"/README.md", []byte(readmeString), 0600)
							if err != nil {
								log.Panic(err)
							}

							pterm.Info.Println("Cleaning files")
							os.Remove(animationDataPath)

							// wg.Done()
							pterm.Info.Println("Generating README for: " + example.Name())
							exampleCode, err = os.ReadFile("./_examples/" + section.Name() + "/" + example.Name() + "/main.go")
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
							pterm.Info.Println("This example took: " + time.Since(exampleRenderStart).String())
						})

					}
					os.WriteFile("./_examples/"+section.Name()+"/README.md", []byte(sectionExamples), 0600)
				})

				pterm.Info.Println("Writing '/_examples/README.md'")
				examplesReadme, _ := os.ReadFile("./_examples/README.md")
				examplesReadmeContent := string(examplesReadme)
				examplesReadmeContent = writeBetween("examples", examplesReadmeContent, "\n"+readmeExamples+"\n")
				os.WriteFile("./_examples/README.md", []byte(examplesReadmeContent), 0600)
			}
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

			pterm.Info.Println("Writing README for https://pterm.sh")
			err = os.WriteFile("./docs/README.md", []byte(newReadmeContent), 0600)
			if err != nil {
				log.Panic(err)
			}

		})
	})
}
