package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// fakeEnv returns a getenv function backed by the given map, so detection
// tests are fully isolated from the real environment.
func fakeEnv(vars map[string]string) func(string) string {
	return func(key string) string {
		return vars[key]
	}
}

func TestDetectLevelFromEnvironment(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want Level
	}{
		{name: "empty environment falls back to basic colors", env: map[string]string{}, want: LevelBasic},
		{name: "NO_COLOR disables colors", env: map[string]string{"NO_COLOR": "1"}, want: LevelNone},
		{name: "NO_COLOR wins over COLORTERM", env: map[string]string{"NO_COLOR": "1", "COLORTERM": "truecolor"}, want: LevelNone},
		{name: "CLICOLOR=0 disables colors", env: map[string]string{"CLICOLOR": "0"}, want: LevelNone},
		{name: "TERM=dumb disables colors", env: map[string]string{"TERM": "dumb"}, want: LevelNone},
		{name: "COLORTERM=truecolor", env: map[string]string{"COLORTERM": "truecolor"}, want: LevelTrueColor},
		{name: "COLORTERM=24bit", env: map[string]string{"COLORTERM": "24bit"}, want: LevelTrueColor},
		{name: "explicit COLORTERM wins over TERM=dumb", env: map[string]string{"COLORTERM": "truecolor", "TERM": "dumb"}, want: LevelTrueColor},
		{name: "TERM=xterm-direct", env: map[string]string{"TERM": "xterm-direct"}, want: LevelTrueColor},
		{name: "TERM=xterm-256color", env: map[string]string{"TERM": "xterm-256color"}, want: Level256},
		{name: "TERM=xterm", env: map[string]string{"TERM": "xterm"}, want: LevelBasic},
		{name: "Windows Terminal", env: map[string]string{"WT_SESSION": "some-guid"}, want: LevelTrueColor},
		{name: "ConEmu with ANSI enabled", env: map[string]string{"ConEmuANSI": "ON"}, want: LevelTrueColor},
		{name: "JetBrains terminal", env: map[string]string{"TERMINAL_EMULATOR": "JetBrains-JediTerm"}, want: LevelTrueColor},
		{name: "iTerm", env: map[string]string{"TERM_PROGRAM": "iTerm.app"}, want: LevelTrueColor},
		{name: "ANSICON marks a translated legacy console", env: map[string]string{"ANSICON": "199x50 (199x300)"}, want: LevelBasic},
		{name: "FORCE_COLOR=0 disables colors", env: map[string]string{"FORCE_COLOR": "0", "COLORTERM": "truecolor"}, want: LevelNone},
		{name: "FORCE_COLOR=1 forces at least basic colors", env: map[string]string{"FORCE_COLOR": "1"}, want: LevelBasic},
		{name: "FORCE_COLOR=1 keeps richer detected colors", env: map[string]string{"FORCE_COLOR": "1", "COLORTERM": "truecolor"}, want: LevelTrueColor},
		{name: "FORCE_COLOR=2 forces 256 colors", env: map[string]string{"FORCE_COLOR": "2"}, want: Level256},
		{name: "FORCE_COLOR=3 forces true color", env: map[string]string{"FORCE_COLOR": "3"}, want: LevelTrueColor},
		{name: "FORCE_COLOR wins over NO_COLOR", env: map[string]string{"FORCE_COLOR": "3", "NO_COLOR": "1"}, want: LevelTrueColor},
		{name: "CLICOLOR_FORCE wins over TERM=dumb", env: map[string]string{"CLICOLOR_FORCE": "1", "TERM": "dumb"}, want: LevelBasic},
		{name: "CLICOLOR_FORCE=0 does not force", env: map[string]string{"CLICOLOR_FORCE": "0", "TERM": "xterm"}, want: LevelBasic},
		{name: "CLICOLOR_FORCE keeps richer detected colors", env: map[string]string{"CLICOLOR_FORCE": "1", "TERM": "xterm-256color"}, want: Level256},
		{name: "CLICOLOR=1 does not disable colors", env: map[string]string{"CLICOLOR": "1", "TERM": "xterm"}, want: LevelBasic},
		{name: "FORCE_COLOR=false disables colors", env: map[string]string{"FORCE_COLOR": "false", "COLORTERM": "truecolor"}, want: LevelNone},
		{name: "FORCE_COLOR=none disables colors", env: map[string]string{"FORCE_COLOR": "none", "COLORTERM": "truecolor"}, want: LevelNone},
		{name: "FORCE_COLOR=true forces at least basic colors", env: map[string]string{"FORCE_COLOR": "true", "TERM": "dumb"}, want: LevelBasic},
		{name: "FORCE_COLOR=2 caps a truecolor environment", env: map[string]string{"FORCE_COLOR": "2", "COLORTERM": "truecolor"}, want: Level256},
		{name: "COLORTERM is case-insensitive", env: map[string]string{"COLORTERM": "TrueColor"}, want: LevelTrueColor},
		{name: "TERM containing truecolor", env: map[string]string{"TERM": "xterm-truecolor"}, want: LevelTrueColor},
		{name: "TERM=screen-256color", env: map[string]string{"TERM": "screen-256color"}, want: Level256},
		{name: "vscode terminal", env: map[string]string{"TERM_PROGRAM": "vscode"}, want: LevelTrueColor},
		{name: "ghostty terminal", env: map[string]string{"TERM_PROGRAM": "ghostty"}, want: LevelTrueColor},
		{name: "Apple Terminal", env: map[string]string{"TERM_PROGRAM": "Apple_Terminal"}, want: LevelTrueColor},
		{name: "Unknown TERM_PROGRAM falls back to basic colors", env: map[string]string{"TERM_PROGRAM": "SomeTerm"}, want: LevelBasic},
		{name: "256color TERM wins over unknown TERM_PROGRAM", env: map[string]string{"TERM_PROGRAM": "SomeTerm", "TERM": "xterm-256color"}, want: Level256},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, detectLevel(fakeEnv(test.env), LevelNone, false))
		})
	}
}

func TestDetectLevelUsesConsoleVerdict(t *testing.T) {
	// A modern Windows console (virtual terminal processing enabled) renders
	// true color even without any descriptive environment variables.
	assert.Equal(t, LevelTrueColor, detectLevel(fakeEnv(nil), LevelTrueColor, true))

	// A legacy Windows console (e.g. cmd.exe on Windows 7/8) cannot render
	// ANSI sequences at all.
	assert.Equal(t, LevelNone, detectLevel(fakeEnv(nil), LevelNone, true))

	// Environment variables win over the console verdict: ANSICON translates
	// ANSI sequences even for consoles that cannot render them natively.
	legacyWithANSICON := map[string]string{"ANSICON": "1"}
	assert.Equal(t, LevelBasic, detectLevel(fakeEnv(legacyWithANSICON), LevelNone, true))

	// Opt-outs win over a capable console.
	noColor := map[string]string{"NO_COLOR": "1"}
	assert.Equal(t, LevelNone, detectLevel(fakeEnv(noColor), LevelTrueColor, true))
}

func TestDetectLevelMatchesConsoleSupport(t *testing.T) {
	// DetectLevel with a scrubbed environment must reflect the platform
	// console verdict: the verdict itself where one exists (Windows with an
	// attached console), the basic fallback everywhere else.
	for _, name := range []string{"NO_COLOR", "FORCE_COLOR", "CLICOLOR", "CLICOLOR_FORCE", "COLORTERM", "TERM", "WT_SESSION", "ConEmuANSI", "TERMINAL_EMULATOR", "TERM_PROGRAM", "ANSICON"} {
		t.Setenv(name, "")
	}

	want := LevelBasic
	if console, known := consoleSupport(); known {
		want = console
	}

	assert.Equal(t, want, DetectLevel())
}
