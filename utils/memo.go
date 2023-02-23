package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type MemoDaySt struct {
	description string
	ymd         time.Time
}

func LineByLine(file string) ([]string, error) {

	var err error
	var res []string

	f, err := os.Open(file)
	if err != nil {
		return res, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		res = append(res, line)
	}
	return res, nil
}

func GetMemoData(filename string) []MemoDaySt {
	data, _ := LineByLine(filename)
	var memodays []MemoDaySt
	today := time.Now()
	for _, s := range data {
		description, day_raw := strings.Split(s, "：")[0], strings.Split(s, "：")[1]
		description = strings.Trim(description, "\n")
		description = strings.TrimSpace(description)
		day_raw = strings.Trim(day_raw, "\n")
		day_raw = strings.TrimSpace(day_raw)
		ymd, _ := time.Parse("2006-01-02 15:04:05", day_raw+" 00:00:00")

		if today.Sub(ymd).Seconds() >= 0 { // 舍弃过期的时间
			continue
		}

		memodays = append(memodays, MemoDaySt{
			description: description,
			ymd:         ymd,
		})
	}
	return memodays
}
