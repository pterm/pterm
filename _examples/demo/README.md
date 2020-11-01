# demo

![Animation](animation.svg)

```go
package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

// Change this to time.Millisecond*200 to speed up the demo.
// Useful when debugging.
const second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	introScreen()
	clear()
	pseudoApplicationHeader()
	time.Sleep(second)
	installingPseudoList()
	time.Sleep(second * 2)
	pterm.DefaultSection.WithLevel(2).Println("Program Install Report")
	installedProgramsSize()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("Tree Printer")
	installedTree()
	time.Sleep(second * 4)
	pterm.DefaultSection.Println("TrueColor Support")
	fadeText()
	time.Sleep(second)
	pterm.DefaultSection.Println("Bullet List Printer")
	listPrinter()
}

func installedTree() {
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Go"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
	}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 2, Text: s})
		}
		if s == "pseudo-chrome" {
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Tabs"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Extensions"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "Refined GitHub"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "GitHub Dark Theme"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 3, Text: "pseudo-Bookmarks"})
			leveledList = append(leveledList, pterm.LeveledListItem{Level: 4, Text: "PTerm"})
		}
	}

	pterm.DefaultTree.WithRoot(pterm.NewTreeFromLeveledList(leveledList)).Render()
}

func installingPseudoList() {
	pterm.DefaultSection.Println("Installing pseudo programs")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
	for i := 0; i < p.Total; i++ {
		p.Title = "Installing " + pseudoProgramList[i]
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(second / 2)
	}
	p.Stop()
}

func listPrinter() {
	pterm.NewBulletListFromString(`Good bye
 Have a nice day!`, " ").Render()
}

func fadeText() {
	from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
	to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

	str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)"
	strs := strings.Split(str, "")
	var fadeInfo string // String which will be used to print info.
	// For loop over the range of the string length.
	for i := 0; i < len(str); i++ {
		// Append faded letter to info string.
		fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
	}
	pterm.Info.Println(fadeInfo)
}

func installedProgramsSize() {
	d := pterm.TableData{{"Program Name", "Status", "Size"}}
	for _, s := range pseudoProgramList {
		if s != "pseudo-minecraft" {
			d = append(d, []string{s, pterm.LightGreen("pass"), strconv.Itoa(randomInt(7, 200)) + "mb"})
		} else {
			d = append(d, []string{pterm.LightRed(s), pterm.LightRed("fail"), "0mb"})
		}
	}
	pterm.DefaultTable.WithHasHeader().WithData(d).Render()
}

func pseudoApplicationHeader() *pterm.TextPrinter {
	return pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"Pseudo Application created with PTerm")
}

func introScreen() {
	pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Render()

	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Println(
		"PTDP - PTerm Demo Program")

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
	time.Sleep(second)
	for i := 14; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(second)
	}
	introSpinner.Stop()
}

func clear() {
	print("\033[H\033[2J")
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

```
