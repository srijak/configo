package configo

import (
	"strings"
	"strconv"
	"os"
	"io/ioutil"
)

const kSeparator = "="

type Configo struct {
	path string
	conf map[string]string
}

type result struct {
	cur_result string
	cur_err    os.Error
	deflt      string
}

func (r result) AsBool() (bool, os.Error) {
	i, err := strconv.Atob(r.cur_result)
	if err != nil {
		d, err := strconv.Atob(r.deflt)
		if err != nil {
			return false, err
		}
		return d, nil
	}
	return i, nil
}
func (r result) AsString() (string, os.Error) {
	if r.cur_err != nil {
		if r.deflt != "" {
			return r.deflt, nil
		}
		return "", r.cur_err
	}
	return r.cur_result, r.cur_err
}
func (r result) AsInt() (int, os.Error) {
	i, err := strconv.Atoi(r.cur_result)
	if err != nil {
		i, err := strconv.Atoi(r.deflt)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return i, nil
}


func (c Configo) Get(k string) *result {
	cur := new(result)
	val, ok := c.conf[k]
	if !ok {
		cur.cur_err = os.NewError("No value for key")
		return cur
	}

	cur.cur_result = val
	return cur
}

func (r result) Default(deflt string) result {
	r.deflt = deflt
	return r
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
		if strings.Index(trim_line, kSeparator) > 0 && !strings.HasPrefix(trim_line, "#") {
			// wish i could use the index to split the string.
			sps := strings.SplitN(trim_line, kSeparator, 2)
			c.conf[strings.TrimSpace(sps[0])] = strings.TrimSpace(sps[1])
		}
	}
	return nil
}
