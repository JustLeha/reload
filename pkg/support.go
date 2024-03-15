package support

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для чтения файла
func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	fileContent := string(data)
	return fileContent, nil
}

// Функция для запись в файл
func WriteToFile(fileName string, data string) error {
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Функция для замены шестнадцатеричных чисел на десятичные
func HexToDecimal(hex string) (string, error) {
	decimal := new(big.Int)
	decimal, err := decimal.SetString(hex, 16)
	if !err {
		return "", fmt.Errorf("Невозможно преобразовать шестнадцатеричное число %s в десятичное", hex)
	}
	return decimal.String(), nil
}

// Функция для замены двоичных чисел на десятичные
func BinaryToDecimal(binary string) (int64, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

// ReplaceNewLines заменяет все вхождения подстроки `search` в строке `text` на подстроку `replacement`.
// Функция принимает три аргумента:
//   - text: исходная строка, в которой производится поиск и замена.
//   - search: строка, которую необходимо заменить.
//   - replacement: строка, на которую необходимо заменить вхождения `search`.
func ReplaceNewLines(text string, search string, replacement string) string {
	// regexp.MustCompile() компилирует регулярное выражение для поиска подстроки `search`.
	re := regexp.MustCompile(search)
	// re.FindAllStringIndex() находит индексы всех вхождений `search` в строке `text`.
	matches := re.FindAllStringIndex(text, -1)
	// replacedText - это объект strings.Builder для накопления результирующей строки после замен.
	var replacedText strings.Builder
	// startID используется для отслеживания начала следующего вхождения `search`.
	startID := 0
	// Обходим все найденные вхождения.
	for _, match := range matches {
		// Выделяем подстроку между предыдущим и текущим вхождением `search`.
		subStr := text[startID:match[0]]
		// Заменяем в этой подстроке все вхождения `search` на строку `replacement` и добавляем результат в replacedText.
		replacedText.WriteString(strings.ReplaceAll(subStr, search, replacement))
		// Добавляем `replacement` столько раз, сколько символов `search` в текущем вхождении.
		replacedText.WriteString(strings.Repeat(replacement, match[1]-match[0]))
		// Обновляем startID до конца текущего вхождения `search`.
		startID = match[1]
	}
	// Добавляем оставшуюся часть строки после последнего вхождения `search`.
	replacedText.WriteString(strings.ReplaceAll(text[startID:], search, replacement))
	// Возвращаем строковое представление результирующей строки из объекта strings.Builder.
	return replacedText.String()
}

// ReplaceAwithAn Функция для замены a на an
func ReplaceAwithAn(text string) string {
	re := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH]\w+)`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}
		wordA := matches[1]
		word := matches[2]
		if wordA == "A" {
			correctWord := "An" + " " + word
			return correctWord
		} else {
			correctWord := "an" + " " + word
			return correctWord
		}

	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

// Функция для применения команд к тексту
func ApplyCommands(content string) string {
	// Регулярное выражение для поиска команд в тексте
	re := regexp.MustCompile(`((\w+)([\s[:punct:]\s]*)[^\d[:punct:]]?)((|[А-Яа-я()\[\][:punct:]]+)(\s*)\((cap|up|low|hex|bin)(?:, (-?\d+))?\))`)
	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем команду и ее аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 7 {
			// Если не удалось извлечь команду, возвращаем исходную строку
			return match
		}
		wordWithSymbol := matches[1]
		word := matches[2]
		symbol := matches[3]
		command := matches[7]
		argStr := matches[8]
		// Если аргумент указан, преобразуем его в число
		var arg int
		if argStr != "" {
			arg, _ = strconv.Atoi(argStr)
			if arg <= 0 {
				return word
			}
			arg--
		}
		commandWithArg := ""
		if arg > 0 {
			commandWithArg = "(" + command + "," + " " + strconv.Itoa(arg) + ")"
		}
		// Применяем команду к слову в зависимости от ее типа
		switch command {
		case "up":
			if arg > 0 && argStr != "" {
				if symbol == "'" && len(symbol) == 1 {
					return commandWithArg + " " + strings.ToUpper(wordWithSymbol)
				}
				return commandWithArg + " " + strings.ToUpper(word) + symbol
			}
			if symbol == "'" && len(symbol) == 1 {
				return commandWithArg + " " + strings.ToUpper(wordWithSymbol)
			}
			return strings.ToUpper(word) + symbol
		case "low":
			if arg > 0 && argStr != "" {
				if symbol == "'" && len(symbol) == 1 {
					return commandWithArg + " " + strings.ToLower(wordWithSymbol)
				}
				return commandWithArg + " " + strings.ToLower(word) + symbol
			}
			if symbol == "'" && len(symbol) == 1 {
				return commandWithArg + " " + strings.ToLower(wordWithSymbol)
			}
			return strings.ToLower(word) + symbol
		case "cap":
			if arg > 0 {
				return commandWithArg + " " + strings.Title(strings.ToLower(word)) + symbol
				if symbol == "'" && len(symbol) == 1 {
					return commandWithArg + " " + strings.Title(strings.ToLower(wordWithSymbol))
				}
			}
			if symbol == "'" && len(symbol) == 1 {
				return commandWithArg + " " + strings.Title(strings.ToLower(wordWithSymbol))
			}
			return strings.Title(strings.ToLower(word)) + symbol
		case "hex":
			hexNum, err := HexToDecimal(word)
			if err != nil {
				return word
			}
			return hexNum + symbol
		case "bin":
			decimalNum, err := BinaryToDecimal(word)
			if err != nil {
				return word
			}
			return strconv.FormatInt(decimalNum, 10)
		default:
			return match + symbol
		}
	}
	// Применяем функцию замены ко всем совпадениям
	result := re.ReplaceAllStringFunc(content, replacer)
	if re.MatchString(result) {
		result = ApplyCommands(result)
	}
	return result
}

// DeleteCommand удаляет команды из текста и удаляет лишние пробелы.
func DeleteCommand(text string) string {
	re := regexp.MustCompile(`\((cap|up|low|CAP|UP|LOW|Cap|Up|Low|hex|bin)(?:,\s*(\d+))?\)`)
	data := re.ReplaceAllString(text, "")
	data = strings.ReplaceAll(data, "(hex\n", "")
	data = strings.ReplaceAll(data, "(bin\n", "")
	data = strings.ReplaceAll(data, "(hex", "")
	data = strings.ReplaceAll(data, "(bin", "")
	content := regexp.MustCompile(`(?m)[^\S\n]+`).ReplaceAllString(data, " ")
	return content
}

// FixQuotest исправляет кавычки в тексте, добавляя пробелы перед и после кавычек, если они отсутствуют.
func FixQuotes(text string) string {
	re := regexp.MustCompile(`([^\w]+)' *([^']+?) *'`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}
		symbol := matches[1]
		word := matches[2]
		newWord := symbol + "'" + word + "'"
		return newWord
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

// FormatPunctuation функция для добавления пробела после запятой, если его нет
func FormatPunctuation(text string) string {
	// Регулярное выражение для добавления пробела после запятой, если его нет
	re := regexp.MustCompile(`(\w+)(\s*)([:,;.!?]+)(?:\s*)(\w*)`)
	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 4 {
			// Если не удалось извлечь аргументы, возвращаем исходную строку
			return match
		}
		word1 := matches[1]
		space := matches[2]
		symbol := matches[3]
		word2 := matches[4]
		if word2 != "" {
			correctWord := word1 + symbol + " " + word2
			return correctWord
		} else if word1 != "" && space == " " && symbol != "" && word2 == "" {
			correctWord := word1 + symbol
			return correctWord
		} else {
			correctWord := match
			return correctWord
		}
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

// SpaceAfterCharter добавляет пробел после знаков препинания в тексте.
func SpaceAfterCharter(text string) string {
	re := regexp.MustCompile(`(\:+|\,+)`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 1 {
			return match
		}
		symbol := matches[1]
		correctWord := symbol + " "
		return correctWord
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

// DeleteSpace удаляет лишние пробелы из текста.
func DeleteSpace(text string) string {
	re := regexp.MustCompile(`( +)`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 1 {
			return match
		}
		correctLetter := " "
		return correctLetter
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

// DeleteSpaceStartAndFinish удаляет пробелы в начале и в конце текста.
func DeleteSpaceStartAndFinish(text string) string {
	re := regexp.MustCompile(`(^\s+|\s+$)`)
	replacer := func(match string) string {
		// Заменяем найденные пробелы на пустую строку
		return ""
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

// CorrectNewLine исправляет символы новой строки в тексте.
func CorrectNewLine(text string) string {
	re := regexp.MustCompile(`(\ *\n+\ *)`)
	replacer := func(match string) string {
		return "\n"
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

//func CorrectNewLine(text string) string {
//	re := regexp.MustCompile(`(\ *\n+\ *)`)
//	result := re.ReplaceAllStringFunc(text, "\n")
//	return result
//}

// ReplaceSymbolToNewLine заменяет повторяющиеся символы новой строки (\n) на одиночный символ новой строки.
// Функция принимает один аргумент:
//   - text: исходная строка, в которой производится замена.
func ReplaceSymbolToNewLine(text string) string {
	// Создаем регулярное выражение для поиска последовательностей символов и символов новой строки.
	re := regexp.MustCompile(`([\w[:punct:]]*)(\n+)`)
	// replacer - это функция, которая вызывается для каждого совпадения с регулярным выражением и заменяет его на новое значение.
	replacer := func(match string) string {
		// re.FindStringSubmatch() находит все подсовпадения в строке, соответствующие регулярному выражению.
		matches := re.FindStringSubmatch(match)
		// Если совпадение не найдено или не содержит две группы, возвращаем исходное совпадение без изменений.
		if len(matches) < 2 {
			return match
		}
		// Извлекаем первую и вторую группы из совпадения.
		group1 := matches[1] // Группа символов и символов пунктуации.
		group2 := matches[2] // Группа символов новой строки.
		// Создаем новую строку для замены повторяющихся символов новой строки на одиночный символ.
		newGroup2 := ""
		if len(group2) > 1 {
			// Если длина второй группы больше 1, значит, есть повторяющиеся символы новой строки.
			// Добавляем в новую строку одиночный символ новой строки столько раз, сколько раз повторяется символ.
			for i := 0; i < len(group2)/2; i++ {
				newGroup2 += "\n"
			}
		}
		// Формируем корректное слово, состоящее из символов первой группы и новой строки.
		correctWord := group1 + newGroup2
		return correctWord
	}
	// re.ReplaceAllStringFunc() заменяет все совпадения с регулярным
	// выражением в тексте на значения, возвращаемые функцией replacer.
	result := re.ReplaceAllStringFunc(text, replacer)
	// Возвращаем обработанный текст с замененными символами новой строки.
	return result
}

// CorrectComma исправляет пробелы перед знаками препинания в тексте.
// Функция принимает один аргумент:
//   - text: исходная строка, в которой производится исправление.
func CorrectComma(text string) string {
	// regexp.MustCompile() компилирует регулярное выражение для поиска
	// пробелов перед знаками препинания (. , : ! ? % $ @).
	re := regexp.MustCompile(`( +)(\,|\.|\:|\!|\?|\%|\$|\@)`)

	// replacer - это функция, которая вызывается для каждого совпадения
	// с регулярным выражением и заменяет его на новое значение.
	replacer := func(match string) string {
		// re.FindStringSubmatch() находит все подсовпадения в строке,
		// соответствующие регулярному выражению.
		matches := re.FindStringSubmatch(match)

		// Если совпадение не найдено или не содержит две группы,
		// возвращаем исходное совпадение без изменений.
		if len(matches) < 2 {
			return match
		}

		// Извлекаем символ знака препинания из второй группы совпадения.
		correctWord := matches[2]
		// Возвращаем только символ знака препинания без пробелов.
		return correctWord
	}
	// re.ReplaceAllStringFunc() заменяет все совпадения с регулярным
	// выражением в тексте на значения, возвращаемые функцией replacer.
	result := re.ReplaceAllStringFunc(text, replacer)

	// Возвращаем обработанный текст без пробелов перед знаками препинания.
	return result
}
