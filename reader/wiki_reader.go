package reader

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var wikiSpaceDelimiter = "_"
var selectNamesRegexp = regexp.MustCompile(`("([^"]|"")*")`)
var selectNonLettersRegexp = regexp.MustCompile(`[^a-zA-Z0-9]+`)

type WikiTitlesReader struct {
	file   *os.File
	reader *bufio.Reader
}

func Open(path string) (reader WikiTitlesReader, err error) {

	file, err := os.OpenFile(path, 0, 0)
	if err != nil {
		return
	}

	reader.file = file
	reader.reader = bufio.NewReader(file)

	// skip first row (column titles)
	_, err = reader.readLine()
	return
}

func (wtr WikiTitlesReader) NextTitle() (string, error) {
	line, err := wtr.readLine()
	if err != nil {
		return "", nil
	}
	// row example: `0	"Blind_Lemon"_Jefferson\n`
	// remove new line character `\n` and first column `0`
	return strings.Split(strings.TrimSuffix(line, "\n"), "\t")[1], nil
}

// title example: `"Blind_Lemon"_Jefferson`
// each title consists with names, and regular words
// each name separated by "" pair. In example case, `Blind_Lemon` is a name.
// current method should return ["blind lemon", "blind", "lemon", "jefferson"] keywords
func (wtr WikiTitlesReader) NextTitleWithKeywords() (string, []string, error) {

	var keywords []string
	title, err := wtr.NextTitle()
	if err != nil {
		return "", keywords, err
	}

	// select names (surrounded by "" text)
	names := selectNamesRegexp.FindAllString(title, -1)
	for _, name := range names {
		keywords = append(keywords, nameKeywords(name)...)
	}

	// select each words separately
	words := strings.Split(title, wikiSpaceDelimiter)
	for _, word := range words {
		keywords = append(keywords, selectNonLettersRegexp.ReplaceAllString(word, ""))
	}

	var filteredKeywords []string
	for _, keyword := range keywords {
		if len(keyword) == 0 || keyword == "" {
			continue
		}
		filteredKeywords = append(filteredKeywords, keyword)
	}
	return title, filteredKeywords, nil
}

func (wtr WikiTitlesReader) readLine() (string, error) {
	return wtr.reader.ReadString('\n')
}

func (wtr WikiTitlesReader) Close() error {
	return wtr.file.Close()
}

func nameKeywords(name string) []string {
	var keywords []string
	name = strings.Replace(name, wikiSpaceDelimiter, " ", -1)
	keywords = append(keywords, name)
	keywords = append(keywords, selectNonLettersRegexp.ReplaceAllString(name, ""))
	return keywords
}
