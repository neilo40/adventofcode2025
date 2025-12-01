package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DownloadInput() {
	fInfo, _ := os.Stat("input.txt")
	if fInfo != nil {
		return
	}

	p, _ := os.Getwd()
	d := filepath.Base(p)
	dNum, _ := strconv.Atoi(strings.TrimPrefix(d, "day"))
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", dNum), nil)
	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", os.Getenv("AOC_SESSION")
	req.AddCookie(cookie)
	req.Header.Set("User-Agent", "neilo40's golang client")
	c := &http.Client{Timeout: 60 * time.Second}
	r, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	bodyBytes, _ := io.ReadAll(r.Body)
	os.WriteFile("input.txt", bodyBytes, 0644)
}
