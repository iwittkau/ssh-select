package metric

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

// Character is a metric that counts the characters
type Character struct {
	count int64
}

// InitMetricFile initializes an empty metric file
func InitMetricFile() error {
	c := Character{count: 0}
	return c.Persist()
}

// Load reads the last metric state
func Load() (char *Character, err error) {

	usr, err := user.Current()
	if err != nil {
		return char, err
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s%s.sshs-metric", usr.HomeDir, string(os.PathSeparator)))
	if err != nil {
		return char, err
	}

	var count int64

	if string(data) == "" {
		count = 0
	} else {
		_, err = fmt.Sscanf(string(data), "%d", &count)
		if err != nil {
			return char, err
		}
	}

	char = &Character{count}

	return char, nil
}

// Count returns the current count
func (c *Character) Count() int64 {
	return c.count
}

// Add addes characters to the count ignoring the provided ignored chars
func (c *Character) Add(characters, ignore string) {
	c.count = c.count + int64(len(characters)) - int64(len(ignore))
}

// Persist saves the metric to disk
func (c *Character) Persist() error {

	usr, err := user.Current()
	if err != nil {
		return err
	}

	data := fmt.Sprintf("%d", c.count)

	err = ioutil.WriteFile(fmt.Sprintf("%s%s.sshs-metric", usr.HomeDir, string(os.PathSeparator)), []byte(data), os.ModePerm)

	return err
}
