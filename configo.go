package configo

import (
	"strings"
	"strconv"
	"os"
  "errors"
	"io/ioutil"
	"reflect"
  )

const kSeparator = "="

type Configo struct {
	Path string
	Conf map[string]string
}

type result struct {
	cur_result string
	cur_err    error
	deflt      string
}

func (r result) AsBool() (bool, error) {
	i, err := strconv.ParseBool(r.cur_result)
	if err != nil {
		d, err := strconv.ParseBool(r.deflt)
		if err != nil {
			return false, err
		}
		return d, nil
	}
	return i, nil
}
func (r result) AsString() (string, error) {
	if r.cur_err != nil {
		if r.deflt != "" {
			return r.deflt, nil
		}
		return "", r.cur_err
	}
	return r.cur_result, r.cur_err
}
func (r result) AsInt() (int, error) {
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

func (c *Configo) Get(k string) *result {
	cur := new(result)
	val, ok := c.Conf[k]
	if !ok {
		cur.cur_err = errors.New("No value for key")
		return cur
	}

	cur.cur_result = val
	return cur
}

func (r result) Default(deflt string) result {
	r.deflt = deflt
	return r
}

func NewConfigo(p string) *Configo {
	c := new(Configo)
	f, err := os.Open(p)
	defer f.Close()
	if err != nil {
		return c
	}

	c.Path = p
	c.Conf = make(map[string]string, 100)

	return c
}

func (c *Configo) Hydrate(st interface{}) *Configo {
	valueOfSt := reflect.ValueOf(st)
	c.Load()
	for k, _ := range c.Conf {
		field := valueOfSt.Elem().FieldByName(k)

		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Int {
				val, _ := c.Get(k).AsInt()
				field.SetInt(int64(val))
			} else if field.Kind() == reflect.Bool {
				val, _ := c.Get(k).AsBool()
				field.SetBool(val)
			} else if field.Kind() == reflect.String {
				val, _ := c.Get(k).AsString()
				field.SetString(val)
			}
		}
	}
	return c
}

func (c *Configo) Load() error {
	file, _ := os.Open(c.Path)
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
			c.Conf[strings.TrimSpace(sps[0])] = strings.TrimSpace(sps[1])
		}
	}
	return nil
}
