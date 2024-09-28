package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/martinlindhe/subtitles"
	"github.com/schollz/progressbar/v3"
)

const eol = "\n"

func main() {
	var filePath, srcLang, destLang string

	flag.StringVar(&srcLang, "src", "de", "source language")
	flag.StringVar(&destLang, "dst", "en", "destination language")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: vtt-translate [flags] <file>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filePath = flag.Arg(0)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("file does not exist")
	}

	srcFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	rawData, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err)
	}
	res, err := subtitles.NewFromVTT(string(rawData))
	if err != nil {
		panic(err)
	}

	fileName := filepath.Base(filePath)

	destFile, err := os.Create("translated_" + fileName)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)
	writer.WriteString("WEBVTT\n\n")

	bar := progressbar.Default(int64(len(res.Captions)))

	for _, cap := range res.Captions {
		srcLine := strings.Join(cap.Text, "\n")
		line := renderLine(cap)
		trans, _ := Translate(srcLine, srcLang, destLang)
		if trans == "" {
			line += eol
		} else {
			trans = strings.Trim(trans, "\n")
			trans = strings.Trim(trans, "\r")
			line += trans + eol + eol
		}
		writer.WriteString(line)
		bar.Add(1)
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("saved to translated_" + fileName)
}

func renderLine(cap subtitles.Caption) string {
	res := TimeVTT(cap.Start) + " --> " + TimeVTT(cap.End) + eol
	for _, line := range cap.Text {
		res += line + eol
	}
	return res
}

// TimeVTT renders a timestamp for use in WebVTT
func TimeVTT(t time.Time) string {
	if t.Hour() == 0 {
		return t.Format("04:05.000")
	}
	return t.Format("15:04:05.000")
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	var text []string
	var result []interface{}

	encodedSource := url.QueryEscape(source)
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" +
		sourceLang + "&tl=" + targetLang + "&dt=t&q=" + encodedSource

	r, err := http.Get(url)
	if err != nil {
		return "", errors.New("error getting translate.googleapis.com")
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", errors.New("error reading response body")
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "", errors.New("error 400 (bad request)")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", errors.New("error unmarshal data")
	}

	if len(result) <= 0 {
		return "", errors.New("no translated data in response")
	}

	inner := result[0]
	for _, slice := range inner.([]interface{}) {
		for _, translatedText := range slice.([]interface{}) {
			text = append(text, fmt.Sprintf("%v", translatedText))
			break
		}
	}
	cText := strings.Join(text, "")
	return cText, nil
}
