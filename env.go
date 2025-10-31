package env

import (
	"os"
	"strings"
)

// Env represents a collection of Items from .env files
type Env struct {
	Items []Item
}

// Item represents key-value pairs
type Item struct {
	Key   string
	Value string
}

// New reads environment variables from a file and returns an Env instance.
// If no filename is provided, it defaults to ".env".
// Returns an error if the file cannot be read.
func New(filename ...string) (*Env, error) {
	var rows []Item
	var f string

	if len(filename) == 0 {
		f = ".env"
	} else {
		f = filename[0]
	}

	file, err := os.ReadFile(f)
	if err != nil {
		//log.Print(err)
		return nil, err
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
	return e, nil
}

// checkDuplicate ensures that if there are multiple keys with different values,
// the last value is used.
func (e *Env) checkDuplicate() {
	var keys []string
	for _, item := range e.Items {
		keys = append(keys, item.Key)
	}
	for i, key := range keys {
		for j := range len(keys) {
			if key == keys[j] {
				e.Items[i].Value = e.Items[j].Value
			}
		}
	}
}

// Get returns the value for the given key from the environment.
// Returns an empty string if the key is not found or Env is nil.
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
