package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RFStr(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return string(buf)
}

func GetLines(day int) []string {
	url_str := "https://adventofcode.com/2022/day/" + strconv.Itoa(day) + "/input"
	session_string := RFStr("session.txt")
	email_string := RFStr("email.txt")
	cookie := &http.Cookie{Name: "session", Value: session_string}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url_str, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.AddCookie(cookie)
	req.Header.Add("User-Agent", "https://github.com/OFFTKP/aoc-22 "+email_string)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Status code:%d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func GetLinesStr(str string) []string {
	scanner := bufio.NewScanner(strings.NewReader(str))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func abs(x int) int {
	return absDiffInt(x, 0)
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x == 0 {
		return 0
	}
	return -1
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if i > max {
			max = i
		}
	}

	return max
}
