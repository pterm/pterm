package pterm_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestDisableDebugMessages(t *testing.T) {
	pterm.PrintDebugMessages = true
	pterm.DisableDebugMessages()
	testza.AssertFalse(t, pterm.PrintDebugMessages)
}

func TestEnableDebugMessages(t *testing.T) {
	pterm.EnableDebugMessages()
	testza.AssertTrue(t, pterm.PrintDebugMessages)
}

func TestDisableOutput(t *testing.T) {
	pterm.DisableOutput()
	testza.AssertFalse(t, pterm.Output)
}

func TestEnableOutput(t *testing.T) {
	pterm.DisableOutput()
	pterm.EnableOutput()
	testza.AssertTrue(t, pterm.Output)
}

func TestDisableStyling(t *testing.T) {
	pterm.RawOutput = false
	pterm.DisableStyling()
	testza.AssertTrue(t, pterm.RawOutput)
}

func TestEnableStyling(t *testing.T) {
	pterm.RawOutput = true
	pterm.EnableStyling()
	testza.AssertFalse(t, pterm.RawOutput)
}

func TestInterfaceImplementation(t *testing.T) {
	// If a printer doesn't fit into the slice, the printer doesn't has the right interface anymore.
	_ = []pterm.TextPrinter{&pterm.DefaultBasicText, pterm.DefaultBox, pterm.DefaultCenter, &pterm.DefaultHeader, &pterm.DefaultParagraph, &pterm.Info, &pterm.DefaultSection, pterm.FgRed, pterm.NewRGB(0, 0, 0)}
	_ = []pterm.LivePrinter{pterm.DefaultProgressbar, &pterm.DefaultSpinner}
	_ = []pterm.RenderPrinter{pterm.NewDefaultBarChart(), pterm.DefaultBigText, pterm.DefaultBulletList, pterm.DefaultPanel, pterm.DefaultTable, pterm.DefaultTree}
}

// Acceptance test

const delay = time.Millisecond

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-dops pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func demoSnapshot(t *testing.T, tick *int) {
	testza.SnapshotCreateOrValidate(t, t.Name()+strconv.Itoa(*tick), readStdout())
	*tick++
}

func TestDemo(t *testing.T) {
	tick := 0
	introScreen()
	demoSnapshot(t, &tick)
	clear()
	demoSnapshot(t, &tick)
	pseudoApplicationHeader()
	demoSnapshot(t, &tick)
	time.Sleep(delay)
	installingPseudoList()
	demoSnapshot(t, &tick)
	time.Sleep(delay * 2)
	pterm.DefaultSection.WithLevel(2).Println("Program Install Report")
	demoSnapshot(t, &tick)
	installedProgramsSize()
	demoSnapshot(t, &tick)
	time.Sleep(delay * 4)
	pterm.DefaultSection.Println("Tree Printer")
	installedTree()
	demoSnapshot(t, &tick)
	time.Sleep(delay * 4)
	pterm.DefaultSection.Println("TrueColor Support")
	fadeText()
	demoSnapshot(t, &tick)
	time.Sleep(delay)
	pterm.DefaultSection.Println("Bullet List Printer")
	listPrinter()
	demoSnapshot(t, &tick)
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
		p.UpdateTitle("Installing " + pseudoProgramList[i])
		if pseudoProgramList[i] == "pseudo-minecraft" {
			pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
		} else {
			pterm.Success.Println("Installing " + pseudoProgramList[i])
			p.Increment()
		}
		time.Sleep(delay / 2)
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
			d = append(d, []string{s, pterm.LightGreen("pass"), "50mb"})
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
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("PTDP - PTerm Demo Program"))

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: 02 Jan 2006 - 15:04:05 MST")
	pterm.Println()
	introSpinner, _ := pterm.DefaultSpinner.WithShowTimer(false).WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
	time.Sleep(delay)
	for i := 14; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " seconds...")
		} else {
			introSpinner.UpdateText("Waiting for " + strconv.Itoa(i) + " second...")
		}
		time.Sleep(delay)
	}
	introSpinner.Stop()
}

func clear() {
	print("\033[H\033[2J")
}
