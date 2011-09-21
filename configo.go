package main

import (
	"strings"
	"strconv"
	"os"
	"io/ioutil"
)


type Configo struct {
	path string
	conf map[string]string
}

func NewConfigo(path string) *Configo {
	c := new(Configo)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return c
	}

	c.path = path
	c.conf = make(map[string]string, 100)

	return c
}

func (c Configo) Load() os.Error {
	file, _ := os.Open(c.path)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(contents), "\n") {
		trim_line := strings.TrimSpace(line)
		if !strings.HasPrefix(trim_line, "#") {
			sps := strings.SplitN(trim_line, "=", 2)
			c.conf[strings.TrimSpace(sps[0])] = strings.TrimSpace(sps[1])
		}
	}
	return nil
}

func (c Configo) getString(k string, deflt string) string {
	return c.conf[k]
}
func (c Configo) getInt(k string, deflt int) int {
	i, err := strconv.Atoi(c.conf[k])
	if err != nil {
		return deflt
	}
	return i
}
func (c Configo) getBool(k string, deflt bool) bool {
	i, err := strconv.Atob(c.conf[k])
	if err != nil {
		return deflt
	}
	return i
}
