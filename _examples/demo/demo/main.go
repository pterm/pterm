package main

import (
	"flag"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

// Speed the demo up, by setting this flag.
// Usefull for debugging.
// Example:
//
//	go run main.go -speedup
var speedup = flag.Bool("speedup", false, "Speed up the demo")
var skipIntro = flag.Bool("skip-intro", false, "Skips the intro")
var second = time.Second

var pseudoProgramList = strings.Split("pseudo-excel pseudo-photoshop pseudo-chrome pseudo-outlook pseudo-explorer "+
	"pseudo-git pseudo-vsc pseudo-intellij pseudo-minecraft pseudo-scoop pseudo-chocolatey", " ")

func main() {
	setup() // Setup the demo (flags etc.)

	// Show intro
	if !*skipIntro {
		introScreen()
		clear()
	}

	showcase("Structured Logging", 5, func() {
		logger := pterm.DefaultLogger.
			WithLevel(pterm.LogLevelTrace)

		logger.Trace("Doing not so important stuff", logger.Args("priority", "super low"))

		time.Sleep(time.Second * 3)

		interstingStuff := map[string]any{
			"when were crayons invented":  "1903",
			"what is the meaning of life": 42,
			"is this interesting":         true,
		}
		logger.Debug("This might be interesting", logger.ArgsFromMap(interstingStuff))
		time.Sleep(time.Second * 3)

		logger.Info("That was actually interesting", logger.Args("such", "wow"))
		time.Sleep(time.Second * 3)
		logger.Warn("Oh no, I see an error coming to us!", logger.Args("speed", 88, "measures", "mph"))
		time.Sleep(time.Second * 3)
		logger.Error("Damn, here it is!", logger.Args("error", "something went wrong"))
		time.Sleep(time.Second * 3)
		logger.Info("But what's really cool is, that you can print very long logs, and PTerm will automatically wrap them for you! Say goodbye to text, that has weird line breaks!", logger.Args("very", "long"))
	})

	showcase("Progress bar", 2, func() {
		pb, _ := pterm.DefaultProgressbar.WithTotal(len(pseudoProgramList)).WithTitle("Installing stuff").Start()
		for i := 0; i < pb.Total; i++ {
			pb.UpdateTitle("Installing " + pseudoProgramList[i])
			if pseudoProgramList[i] == "pseudo-minecraft" {
				pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
			} else {
				pterm.Success.Println("Installing " + pseudoProgramList[i])
			}
			pb.Increment()
			time.Sleep(second / 2)
		}
		pb.Stop()
	})

	showcase("Spinner", 2, func() {
		list := pseudoProgramList[7:]
		spinner, _ := pterm.DefaultSpinner.Start("Installing stuff")
		for i := 0; i < len(list); i++ {
			spinner.UpdateText("Installing " + list[i])
			if list[i] == "pseudo-minecraft" {
				pterm.Warning.Println("Could not install pseudo-minecraft\nThe company policy forbids games.")
			} else {
				pterm.Success.Println("Installing " + list[i])
			}
			time.Sleep(second)
		}
		spinner.Success()
	})

	showcase("Live Output", 2, func() {
		pterm.Info.Println("You can use an Area to display changing output:")
		pterm.Println()
		area, _ := pterm.DefaultArea.WithCenter().Start() // Start the Area printer, with the Center option.
		for i := 0; i < 10; i++ {
			str, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(time.Now().Format("15:04:05"))).Srender() // Save current time in str.
			area.Update(str)                                                                                              // Update Area contents.
			time.Sleep(time.Second)
		}
		area.Stop()
	})

	showcase("Tables", 4, func() {
		for i := 0; i < 3; i++ {
			pterm.Println()
		}
		td := [][]string{
			{"Library", "Description"},
			{"PTerm", "Make beautiful CLIs"},
			{"Testza", "Programmer friendly test framework"},
			{"Cursor", "Move the cursor around the terminal"},
		}
		table, _ := pterm.DefaultTable.WithHasHeader().WithData(td).Srender()
		boxedTable, _ := pterm.DefaultTable.WithHasHeader().WithData(td).WithBoxed().Srender()
		pterm.DefaultCenter.Println(table)
		pterm.DefaultCenter.Println(boxedTable)
	})

	showcase("TrueColor Support", 7, func() {
		from := pterm.NewRGB(0, 255, 255) // This RGB value is used as the gradients start point.
		to := pterm.NewRGB(255, 0, 255)   // This RGB value is used as the gradients first point.

		str := "If your terminal has TrueColor support, you can use RGB colors!\nYou can even fade them :)\n\nLorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."
		strs := strings.Split(str, "")
		var fadeInfo string // String which will be used to print info.
		// For loop over the range of the string length.
		for i := 0; i < len(str); i++ {
			// Append faded letter to info string.
			fadeInfo += from.Fade(0, float32(len(str)), float32(i), to).Sprint(strs[i])
		}
		pterm.DefaultCenter.WithCenterEachLineSeparately().Println(fadeInfo)
	})

	showcase("Fully Customizable", 2, func() {
		for i := 0; i < 4; i++ {
			pterm.Println()
		}
		text := "All printers are fully customizable!"
		area := pterm.DefaultArea.WithCenter()
		area.Update(pterm.DefaultBox.Sprintln(text))
		time.Sleep(second)
		area.Update(pterm.DefaultBox.WithTopPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleTopLeft().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleTopCenter().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleTopRight().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleBottomRight().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleBottomCenter().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithTitle("Some title!").WithTitleBottomLeft().Sprintln(text))
		time.Sleep(second / 3)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithBoxStyle(pterm.NewStyle(pterm.FgCyan)).Sprintln(text))
		time.Sleep(second / 5)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithBoxStyle(pterm.NewStyle(pterm.FgRed)).Sprintln(text))
		time.Sleep(second / 5)
		area.Update(pterm.DefaultBox.WithTopPadding(1).WithBottomPadding(1).WithLeftPadding(1).WithRightPadding(1).WithBoxStyle(pterm.NewStyle(pterm.FgGreen)).Sprintln(text))
		time.Sleep(second / 5)
		area.Update(pterm.DefaultBox.WithTopPadding(1).
			WithBottomPadding(1).
			WithLeftPadding(1).
			WithRightPadding(1).
			WithHorizontalString("═").
			WithVerticalString("║").
			WithBottomLeftCornerString("╗").
			WithBottomRightCornerString("╔").
			WithTopLeftCornerString("╝").
			WithTopRightCornerString("╚").
			Sprintln(text))
		area.Stop()
	})

	showcase("Themes", 2, func() {
		pterm.Info.Println("You can change the color theme of PTerm easily to fit your needs!\nThis is the default one:")
		time.Sleep(second / 2)
		// Print every value of the default theme with its own style.
		v := reflect.ValueOf(pterm.ThemeDefault)
		typeOfS := v.Type()

		if typeOfS == reflect.TypeOf(pterm.Theme{}) {
			for i := 0; i < v.NumField(); i++ {
				field, ok := v.Field(i).Interface().(pterm.Style)
				if ok {
					field.Println(typeOfS.Field(i).Name)
				}
				time.Sleep(time.Millisecond * 250)
			}
		}
	})

	showcase("And much more!", 3, func() {
		for i := 0; i < 4; i++ {
			pterm.Println()
		}
		box := pterm.DefaultBox.
			WithBottomPadding(1).
			WithTopPadding(1).
			WithLeftPadding(3).
			WithRightPadding(3).
			Sprintf("Have fun exploring %s!", pterm.Cyan("PTerm"))
		pterm.DefaultCenter.Println(box)
	})
}

func setup() {
	flag.Parse()
	if *speedup {
		second = time.Millisecond * 200
	}
}

func introScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.NewStyle(pterm.FgLightCyan)),
		putils.LettersFromStringWithStyle("Term", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("PTDP - PTerm Demo Program"))

	pterm.Info.Println("This animation was generated with the latest version of PTerm!" +
		"\nPTerm works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want, you can customize everything :)" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("./_examples/demo") + " directory." +
		"\n" +
		"\nThis demo was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
	introSpinner, _ := pterm.DefaultSpinner.WithShowTimer(false).WithRemoveWhenDone(true).Start("Waiting for 15 seconds...")
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

func showcase(title string, seconds int, content func()) {
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithFullWidth().Println(title)
	pterm.Println()
	time.Sleep(second / 2)
	content()
	time.Sleep(second * time.Duration(seconds))
	print("\033[H\033[2J")
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
