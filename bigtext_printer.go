package pterm

import (
	"io"
	"strings"

	"github.com/gookit/color"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// Letters is a slice of Letter.
type Letters []Letter

// Letter is an object, which holds a string and a specific Style for it.
type Letter struct {
	String string
	Style  *Style
	RGB    RGB
}

// WithStyle returns a new Letter with a specific Style.
func (l Letter) WithStyle(style *Style) *Letter {
	l.Style = style
	return &l
}

// WithRGB returns a new Letter with a specific RGB color (overwrites style).
func (l Letter) WithRGB(rgb RGB) *Letter {
	l.RGB = rgb
	return &l
}

// WithString returns a new Letter with a specific String.
func (l Letter) WithString(s string) *Letter {
	l.String = s
	return &l
}

// BigTextPrinter renders big text.
// You can use this as title screen for your application.
type BigTextPrinter struct {
	// BigCharacters holds the map from a normal character to it's big version.
	BigCharacters map[string]string
	Letters       Letters
	Writer        io.Writer
}

// WithBigCharacters returns a new BigTextPrinter with specific BigCharacters.
func (p BigTextPrinter) WithBigCharacters(chars map[string]string) *BigTextPrinter {
	p.BigCharacters = chars
	return &p
}

// WithLetters returns a new BigTextPrinter with specific Letters
func (p BigTextPrinter) WithLetters(letters ...Letters) *BigTextPrinter {
	l := Letters{}
	for _, letter := range letters {
		l = append(l, letter...)
	}
	p.Letters = l
	return &p
}

// WithWriter sets the custom Writer.
func (p BigTextPrinter) WithWriter(writer io.Writer) *BigTextPrinter {
	p.Writer = writer
	return &p
}

// Srender renders the BigText as a string.
func (p BigTextPrinter) Srender() (string, error) {
	var ret string

	if RawOutput {
		for _, letter := range p.Letters {
			ret += letter.String
		}
		return ret, nil
	}

	var bigLetters Letters
	for _, l := range p.Letters {
		if val, ok := p.BigCharacters[l.String]; ok {
			bigLetters = append(bigLetters, Letter{
				String: val,
				Style:  l.Style,
				RGB:    l.RGB,
			})
		}
	}

	var maxHeight int

	for _, l := range bigLetters {
		h := strings.Count(l.String, "\n")
		if h > maxHeight {
			maxHeight = h
		}
	}

	for i := 0; i <= maxHeight; i++ {
		for _, letter := range bigLetters {
			var letterLine string
			letterLines := strings.Split(letter.String, "\n")
			maxLetterWidth := internal.GetStringMaxWidth(letter.String)
			if len(letterLines) > i {
				letterLine = letterLines[i]
			}
			letterLineLength := runewidth.StringWidth(letterLine)
			if letterLineLength < maxLetterWidth {
				letterLine += strings.Repeat(" ", maxLetterWidth-letterLineLength)
			}

			if letter.RGB != (RGB{}) && (color.IsSupportRGBColor() || internal.RunsInCi()) {
				ret += letter.RGB.Sprint(letterLine)
			} else {
				ret += letter.Style.Sprint(letterLine)
			}
		}
		ret += "\n"
	}

	return ret, nil
}

// Render prints the BigText to the terminal.
func (p BigTextPrinter) Render() error {
	s, _ := p.Srender()
	Fprintln(p.Writer, s)

	return nil
}

// DefaultBigText contains default values for BigTextPrinter.
var DefaultBigText = BigTextPrinter{
	BigCharacters: map[string]string{
		"a": ` █████  
██   ██ 
███████ 
██   ██ 
██   ██ `,
		"A": ` █████  
██   ██ 
███████ 
██   ██ 
██   ██ `,
		"b": `██████  
██   ██ 
██████  
██   ██ 
██████`,
		"B": `██████  
██   ██ 
██████  
██   ██ 
██████`,
		"c": ` ██████ 
██      
██      
██      
 ██████`,
		"C": ` ██████ 
██      
██      
██      
 ██████`,
		"d": `██████  
██   ██ 
██   ██ 
██   ██ 
██████ `,
		"D": `██████  
██   ██ 
██   ██ 
██   ██ 
██████ `,
		"e": `███████ 
██      
█████   
██      
███████`,
		"E": `███████ 
██      
█████   
██      
███████`,
		"f": `███████ 
██      
█████   
██      
██     `,
		"F": `███████ 
██      
█████   
██      
██     `,
		"g": ` ██████  
██       
██   ███ 
██    ██ 
 ██████  `,
		"G": ` ██████  
██       
██   ███ 
██    ██ 
 ██████  `,
		"h": `██   ██ 
██   ██ 
███████ 
██   ██ 
██   ██ `,
		"H": `██   ██ 
██   ██ 
███████ 
██   ██ 
██   ██ `,
		"i": `██ 
██ 
██ 
██ 
██`,
		"I": `██ 
██ 
██ 
██ 
██`,
		"j": `     ██ 
     ██ 
     ██ 
██   ██ 
 █████ `,
		"J": `     ██ 
     ██ 
     ██ 
██   ██ 
 █████ `,
		"k": `██   ██ 
██  ██  
█████   
██  ██  
██   ██`,
		"K": `██   ██ 
██  ██  
█████   
██  ██  
██   ██`,
		"l": `██      
██      
██      
██      
███████ `,
		"L": `██      
██      
██      
██      
███████ `,
		"m": `███    ███ 
████  ████ 
██ ████ ██ 
██  ██  ██ 
██      ██`,
		"M": `███    ███ 
████  ████ 
██ ████ ██ 
██  ██  ██ 
██      ██`,
		"n": `███    ██ 
████   ██ 
██ ██  ██ 
██  ██ ██ 
██   ████`,
		"N": `███    ██ 
████   ██ 
██ ██  ██ 
██  ██ ██ 
██   ████`,
		"o": ` ██████  
██    ██ 
██    ██ 
██    ██ 
 ██████  `,
		"O": ` ██████  
██    ██ 
██    ██ 
██    ██ 
 ██████  `,
		"p": `██████  
██   ██ 
██████  
██      
██     `,
		"P": `██████  
██   ██ 
██████  
██      
██     `,
		"q": ` ██████  
██    ██ 
██    ██ 
██ ▄▄ ██ 
 ██████  
    ▀▀   `,
		"Q": ` ██████  
██    ██ 
██    ██ 
██ ▄▄ ██ 
 ██████  
    ▀▀   `,
		"r": `██████  
██   ██ 
██████  
██   ██ 
██   ██`,
		"R": `██████  
██   ██ 
██████  
██   ██ 
██   ██`,
		"s": `███████ 
██      
███████ 
     ██ 
███████`,
		"S": `███████ 
██      
███████ 
     ██ 
███████`,
		"t": `████████ 
   ██    
   ██    
   ██    
   ██    `,
		"T": `████████ 
   ██    
   ██    
   ██    
   ██    `,
		"u": `██    ██ 
██    ██ 
██    ██ 
██    ██ 
 ██████ `,
		"U": `██    ██ 
██    ██ 
██    ██ 
██    ██ 
 ██████ `,
		"v": `██    ██ 
██    ██ 
██    ██ 
 ██  ██  
  ████   `,
		"V": `██    ██ 
██    ██ 
██    ██ 
 ██  ██  
  ████   `,
		"w": `██     ██ 
██     ██ 
██  █  ██ 
██ ███ ██ 
 ███ ███ `,
		"W": `██     ██ 
██     ██ 
██  █  ██ 
██ ███ ██ 
 ███ ███ `,
		"x": `██   ██ 
 ██ ██  
  ███   
 ██ ██  
██   ██ `,
		"X": `██   ██ 
 ██ ██  
  ███   
 ██ ██  
██   ██ `,
		"y": `██    ██ 
 ██  ██  
  ████   
   ██    
   ██   `,
		"Y": `██    ██ 
 ██  ██  
  ████   
   ██    
   ██   `,
		"z": `███████ 
   ███  
  ███   
 ███    
███████`,
		"Z": `███████ 
   ███  
  ███   
 ███    
███████`,
		"0": ` ██████  
██  ████ 
██ ██ ██ 
████  ██ 
 ██████ `,
		"1": ` ██ 
███ 
 ██ 
 ██ 
 ██ `,
		"2": `██████  
     ██ 
 █████  
██      
███████ `,
		"3": `██████  
     ██ 
 █████  
     ██ 
██████ `,
		"4": `██   ██ 
██   ██ 
███████ 
     ██ 
     ██ `,
		"5": `███████ 
██      
███████ 
     ██ 
███████`,
		"6": ` ██████  
██       
███████  
██    ██ 
 ██████ `,
		"7": `███████ 
     ██ 
    ██  
   ██   
   ██`,
		"8": ` █████  
██   ██ 
 █████  
██   ██ 
 █████ `,
		"9": ` █████  
██   ██ 
 ██████ 
     ██ 
 █████ `,
		" ": "    ",
		"!": `██ 
██ 
██ 
   
██ `,
		"$": `▄▄███▄▄·
██      
███████ 
     ██ 
███████ 
  ▀▀▀  `,
		"%": `██  ██ 
   ██  
  ██   
 ██    
██  ██`,
		"/": `    ██ 
   ██  
  ██   
 ██    
██   `,
		"(": ` ██ 
██  
██  
██  
 ██ `,
		")": `██  
 ██ 
 ██ 
 ██ 
██  `,
		"?": `██████  
     ██ 
  ▄███  
  ▀▀    
  ██   `,
		"[": `███ 
██  
██  
██  
███`,
		"]": `███ 
 ██ 
 ██ 
 ██ 
███ `,
		".": `   
   
   
   
██`,
		",": `   
   
   
   
▄█`,
		"-": `      
      
█████ 
      
      
     `,
		"<": `  ██ 
 ██  
██   
 ██  
  ██ `,
		">": `██   
 ██  
  ██ 
 ██  
██ `,
		"*": `      
▄ ██ ▄
 ████ 
▀ ██ ▀
     `,
		"#": ` ██  ██  
████████ 
 ██  ██  
████████ 
 ██  ██ `,
		"_": `        
        
        
        
███████ `,
		":": `   
██ 
   
   
██ `,
		"°": ` ████  
██  ██ 
 ████  
       
      `,
	},
}
