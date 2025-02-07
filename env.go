package env

import (
	"log"
	"os"
	"strings"
)

type Env struct {
	Items []Item
}

type Item struct {
	Key   string
	Value string
}

func New(filename ...string) *Env {
	var rows []Item
	var f string

	if len(filename) == 0 {
		f = ".env"
	} else {
		f = filename[0]
	}

	file, err := os.ReadFile(f)
	if err != nil {
		log.Print(err)
		return nil
	}
	fileString := string(file)

	lines := strings.Split(fileString, "\n")

	for _, line := range lines {
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			continue
		}
		key := pair[0]
		value := pair[1]
		row := Item{Key: key, Value: strings.Trim(value, "\"")}
		rows = append(rows, row)
	}

	e := &Env{Items: rows}
	e.checkDuplicate()
	return e
}

// If there are multiple keys with different values, use the last value
func (e *Env) checkDuplicate() {
	var keys []string
	for _, item := range e.Items {
		keys = append(keys, item.Key)
	}
	for i, key := range keys {
		for j := range len(keys) {
			if key == keys[j] {
				//fmt.Println("dup")
				e.Items[i].Value = e.Items[j].Value
			}
		}
	}
}

func (e *Env) Get(key string) string {
	if e == nil {
		return ""
	}
	for _, j := range e.Items {
		if j.Key == key {
			return j.Value
		} else {
			continue
		}
	}
	return ""
}
