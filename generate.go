package homograph

import (
	"golang.org/x/net/idna"
)

type generateFn func(c rune) []string

var p = idna.New()

func Generate(input string, minimum bool) []string {
	gens := make([][]string, len(input))
	for i, c := range input {
		if minimum {
			gens[i] = generateSimple(c)
		} else {
			gens[i] = generateAll(c)
		}
	}

	results := []string{}
	setAll(gens, &results, "")

	return results
}

func Punycode(input string, minimum bool) []string {
	gens := make([][]string, len(input))
	for i, c := range input {
		if minimum {
			gens[i] = generateSimple(c)
		} else {
			gens[i] = generateAll(c)
		}
	}

	results := []string{}
	setPunycodeAll(gens, &results, "")

	return results
}

func setPunycodeAll(all [][]string, results *[]string, currentStr string) {
	if len(all) == 1 {
		for _, s := range all[0] {
			ps, err := p.ToASCII(currentStr + s)
			if err == nil {
				*results = append(*results, ps)
			} else {
				*results = append(*results, currentStr+s)
			}
		}
	} else {
		for _, s := range all[0] {
			setPunycodeAll(all[1:], results, currentStr+s)
		}
	}
}

func setAll(all [][]string, results *[]string, currentStr string) {
	if len(all) == 1 {
		for _, s := range all[0] {
			*results = append(*results, currentStr+s)
		}
	} else {
		for _, s := range all[0] {
			setAll(all[1:], results, currentStr+s)
		}
	}
}

func generateAll(c rune) []string {
	switch string(c) {
	case "!":
		return []string{"!", "ǃ", "！"}
	case "\"":
		return []string{"\"", "״", "″", "＂"}
	case "$":
		return []string{"$", "＄"}
	case "%":
		return []string{"%", "％"}
	case "&":
		return []string{"&", "＆"}
	case "'":
		return []string{"'", "＇"}
	case "(":
		return []string{"(", "﹝", "（"}
	case ")":
		return []string{")", "﹞", "）"}
	case "*":
		return []string{"*", "⁎", "＊"}
	case "+":
		return []string{"+", "＋"}
	case ",":
		return []string{",", "‚", "，"}
	case "-":
		return []string{"-", "‐", "𐀠", "－"}
	case ".":
		return []string{".", "٠", "۔", "܁", "܂", "․", "‧", "。", "．", "｡"}
	case "/":
		return []string{"/", "̸", "⁄", "∕", "╱", "⫻", "⫽", "／", "ﾉ"}
	case "0":
		return []string{"0", "O", "o", "Ο", "ο", "О", "о", "Օ", "𐐠", "Ｏ", "ｏ"}
	case "1":
		return []string{"1", "I", "ا", "１"}
	case "2":
		return []string{"2", "２"}
	case "3":
		return []string{"3", "３"}
	case "4":
		return []string{"4", "４"}
	case "5":
		return []string{"5", "５"}
	case "6":
		return []string{"6", "６"}
	case "7":
		return []string{"7", "𐐠", "７"}
	case "8":
		return []string{"8", "Ց", "８"}
	case "9":
		return []string{"9", "９"}
	case ":":
		return []string{":", "։", "܃", "܄", "∶", "꞉", "："}
	case ";":
		return []string{";", ";", "；"}
	case "<":
		return []string{"<", "‹", "＜"}
	case "=":
		return []string{"=", "𐀠", "＝"}
	case ">":
		return []string{">", "›", "＞"}
	case "?":
		return []string{"?", "？"}
	case "@":
		return []string{"@", "＠"}
	case "[":
		return []string{"[", "［"}
	case "\\":
		return []string{"\\", "＼"}
	case "]":
		return []string{"]", "］"}
	case "^":
		return []string{"^", "＾"}
	case "_":
		return []string{"_", "＿"}
	case "`":
		return []string{"`", "｀"}
	case "a":
		return []string{"A", "a", "À", "Á", "Â", "Ã", "Ä", "Å", "à", "á", "â", "ã", "ä", "å", "ɑ", "Α", "α", "а", "Ꭺ", "Ａ", "ａ"}
	case "b":
		return []string{"B", "b", "ß", "ʙ", "Β", "β", "В", "Ь", "Ᏼ", "ᛒ", "Ｂ", "ｂ"}
	case "c":
		return []string{"C", "c", "ϲ", "Ϲ", "С", "с", "Ꮯ", "Ⅽ", "ⅽ", "Ｃ", "ｃ"}
	case "d":
		return []string{"D", "d", "Ď", "ď", "Đ", "đ", "ԁ", "ժ", "Ꭰ", "ḍ", "Ⅾ", "ⅾ", "Ｄ", "ｄ"}
	case "e":
		return []string{"E", "e", "È", "É", "Ê", "Ë", "é", "ê", "ë", "Ē", "ē", "Ĕ", "ĕ", "Ė", "ė", "Ę", "Ě", "ě", "Ε", "Е", "е", "Ꭼ", "Ｅ", "ｅ"}
	case "f":
		return []string{"F", "f", "Ϝ", "Ｆ", "ｆ"}
	case "g":
		return []string{"G", "g", "ɡ", "ɢ", "Ԍ", "ն", "Ꮐ", "Ｇ", "ｇ"}
	case "h":
		return []string{"H", "h", "ʜ", "Η", "Н", "һ", "Ꮋ", "Ｈ", "ｈ"}
	case "i":
		return []string{"I", "i", "l", "ɩ", "Ι", "І", "і", "ا", "Ꭵ", "ᛁ", "Ⅰ", "ⅰ", "𐐠", "Ｉ", "ｉ"}
	case "j":
		return []string{"J", "j", "ϳ", "Ј", "ј", "յ", "Ꭻ", "Ｊ", "ｊ"}
	case "k":
		return []string{"K", "k", "Κ", "κ", "К", "Ꮶ", "ᛕ", "K", "Ｋ", "ｋ"}
	case "l":
		return []string{"L", "l", "ʟ", "ι", "ا", "Ꮮ", "Ⅼ", "ⅼ", "Ｌ", "ｌ"}
	case "m":
		return []string{"M", "m", "Μ", "Ϻ", "М", "Ꮇ", "ᛖ", "Ⅿ", "ⅿ", "Ｍ", "ｍ"}
	case "n":
		return []string{"N", "n", "ɴ", "Ν", "Ｎ", "ｎ"}
	case "o":
		return []string{"0", "O", "o", "Ο", "ο", "О", "о", "Օ", "Ｏ", "ｏ"}
	case "p":
		return []string{"P", "p", "Ρ", "ρ", "Р", "р", "Ꮲ", "Ｐ", "ｐ"}
	case "q":
		return []string{"Q", "q", "Ⴍ", "Ⴓ", "Ｑ", "ｑ"}
	case "r":
		return []string{"R", "r", "ʀ", "Ի", "Ꮢ", "ᚱ", "Ｒ", "ｒ"}
	case "s":
		return []string{"S", "s", "Ѕ", "ѕ", "Տ", "Ⴝ", "Ꮪ", "𐐠", "Ｓ", "ｓ"}
	case "t":
		return []string{"T", "t", "Τ", "τ", "Т", "Ꭲ", "Ｔ", "ｔ"}
	case "u":
		return []string{"U", "u", "μ", "υ", "Ա", "Ս", "⋃", "Ｕ", "ｕ"}
	case "v":
		return []string{"V", "v", "ν", "Ѵ", "ѵ", "Ꮩ", "Ⅴ", "ⅴ", "Ｖ", "ｖ"}
	case "w":
		return []string{"W", "w", "ѡ", "Ꮃ", "Ｗ", "ｗ"}
	case "x":
		return []string{"X", "x", "Χ", "χ", "Х", "х", "Ⅹ", "ⅹ", "Ｘ", "ｘ"}
	case "y":
		return []string{"Y", "y", "ʏ", "Υ", "γ", "у", "Ү", "Ｙ", "ｙ"}
	case "z":
		return []string{"Z", "z", "Ζ", "Ꮓ", "Ｚ", "ｚ"}
	case "{":
		return []string{"{", "｛"}
	case "|":
		return []string{"|", "ǀ", "ا", "｜"}
	case "}":
		return []string{"}", "｝"}
	case "~":
		return []string{"~", "⁓", "～"}
	case "ß":
		return []string{"ß", "ӧ"}
	case "ä":
		return []string{"ä", "Ä", "Ӓ"}
	case "ö":
		return []string{"ö", "Ö", "Ӧ"}

	}

	return []string{string(c)}
}

func generateSimple(c rune) []string {
	switch string(c) {
	case "0":
		return []string{"0", "O", "Ο", "О"}
	case "1":
		return []string{"1", "１"}
	case "2":
		return []string{"2", "２"}
	case "3":
		return []string{"3", "３"}
	case "4":
		return []string{"4", "４"}
	case "5":
		return []string{"5", "５"}
	case "6":
		return []string{"6", "６"}
	case "7":
		return []string{"7", "７"}
	case "8":
		return []string{"8", "Ց", "８"}
	case "9":
		return []string{"9", "９"}
	case "a":
		return []string{"a", "ɑ", "α", "а", "ａ"}
	case "b":
		return []string{"b", "Ь", "ｂ"}
	case "c":
		return []string{"c", "ϲ", "с", "ⅽ", "ｃ"}
	case "d":
		return []string{"d", "ḍ", "ⅾ", "ｄ"}
	case "e":
		return []string{"e", "е", "ｅ"}
	case "f":
		return []string{"f", "ｆ"}
	case "g":
		return []string{"g", "ɡ", "ｇ"}
	case "h":
		return []string{"h", "һ", "ｈ"}
	case "i":
		return []string{"i", "l", "і", "ⅰ", "ｉ"}
	case "j":
		return []string{"j", "ϳ", "ј", "ｊ"}
	case "k":
		return []string{"k", "κ", "ᛕ", "ｋ"}
	case "l":
		return []string{"l", "ا", "ⅼ", "ｌ"}
	case "m":
		return []string{"m", "ⅿ", "ｍ"}
	case "n":
		return []string{"n", "ɴ", "ｎ"}
	case "o":
		return []string{"0", "O", "o", "Ο", "ο", "О", "о", "ｏ"}
	case "p":
		return []string{"p", "ρ", "р", "ｐ"}
	case "q":
		return []string{"q", "ｑ"}
	case "r":
		return []string{"r", "ｒ"}
	case "s":
		return []string{"s", "ѕ", "ｓ"}
	case "t":
		return []string{"t", "ｔ"}
	case "u":
		return []string{"u", "υ", "⋃", "ｕ"}
	case "v":
		return []string{"v", "ν", "ｖ"}
	case "w":
		return []string{"w", "ѡ", "ｗ"}
	case "x":
		return []string{"x", "х", "ⅹ", "ｘ"}
	case "y":
		return []string{"y", "у", "ｙ"}
	case "z":
		return []string{"z", "ｚ"}
	}

	return []string{string(c)}
}
