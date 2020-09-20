package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	fmt.Println("## Generating Examples")
	files, err := ioutil.ReadDir("./_examples/")
	if err != nil {
		log.Fatal(err)
	}

	var readmeExamples string

	for _, f := range files {
		fmt.Println("### Generating animations for example '" + f.Name() + "'")
		animationDataPath := "./_examples/" + f.Name() + "/animation_data.json"
		animationSvgPath := "./_examples/" + f.Name() + "/animation.svg"
		exampleCode, err := ioutil.ReadFile("./_examples/" + f.Name() + "/main.go")
		if err != nil {
			log.Fatal(err)
		}

		if fileExists(animationDataPath) {
			fmt.Println("#### animation_data.json already exists. Removing it.")
			err := os.Remove(animationDataPath)
			if err != nil {
				log.Fatal(err)
			}
		}
		if fileExists(animationDataPath) {
			fmt.Println("#### animation.svg already exists. Removing it.")
			err := os.Remove(animationSvgPath)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("#### Running asciinema")
		execute(`asciinema rec ` + animationDataPath + ` -c "go run ./_examples/` + f.Name() + `"`)

		fmt.Println("#### Adding sleep to end of animation_data.json")
		sleepString := `[5, "o", "\r\nrestarting...\r\n"]`
		animationDataFile, err := os.OpenFile(animationDataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Println(err)
		}
		defer animationDataFile.Close()
		if _, err := animationDataFile.WriteString(sleepString); err != nil {
			log.Println(err)
		}

		fmt.Println("#### Generating SVG")
		execute(`svg-term --in ` + animationDataPath + ` --out ` + animationSvgPath + ` --window true --no-optimize`)

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

		readmeExamples += "### " + f.Name() + "\n\n"
		readmeExamples += "![Animation](https://raw.githubusercontent.com/pterm/pterm/master/_examples/" + f.Name() + "/animation.svg)\n\n"
		readmeExamples += "```go\n"
		readmeExamples += string(exampleCode) + "\n"
		readmeExamples += "```"
		readmeExamples += "\n\n"

		fmt.Println("#### Cleaning files")
		os.Remove(animationDataPath)

	}

	fmt.Println("### Appending examples to root README.md")

	readmeContent, err := ioutil.ReadFile("./README.md")
	if err != nil {
		log.Fatal(err)
	}

	beforeRegex := regexp.MustCompile(`(?ms).*<!-- examples:start -->`)
	afterRegex := regexp.MustCompile(`(?ms)<!-- examples:end -->.*`)

	before := beforeRegex.FindAllString(string(readmeContent), 1)[0]
	after := afterRegex.FindAllString(string(readmeContent), 1)[0]

	var newReadmeContent string

	newReadmeContent += before + "\n"
	newReadmeContent += readmeExamples
	newReadmeContent += after + "\n"

	err = ioutil.WriteFile("./README.md", []byte(newReadmeContent), 0600)
	if err != nil {
		log.Fatal(err)
	}

}

func execute(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Println("Running: " + cmd.String())
	err := cmd.Run()
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
