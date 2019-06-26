package metric_test

import (
	"testing"

	"github.com/5FeetUnder/ssh-select/metric"
)

func TestCharacter(t *testing.T) {
	err := metric.InitMetricFile()
	if err != nil {
		t.Fatal(err)
	}

	m, err := metric.Load()
	if err != nil {
		t.Fatal(err)
	}

	if m.Count() != 0 {
		t.Fatal("Count() should return 0")
	}

	m.Add("abc", "")

	if m.Count() != 3 {
		t.Fatal("Count() should return 3")
	}

	m.Add("abc", "bc")

	if m.Count() != 4 {
		t.Fatal("Count() should return 4")
	}

	err = m.Persist()
	if err != nil {
		t.Fatal(err)
	}

}
