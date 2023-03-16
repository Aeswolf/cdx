package funcs_test

import (
	"testing"

	"github.com/Aeswolf/cdx/cd"
	"github.com/Aeswolf/cdx/funcs"
)

type testStruct struct {
	data []string
	outcome,
	expected string
}

// testing the gepath func
func TestGePath(t *testing.T) {
	testData := []testStruct{
		{data: []string{"Documents", "GoWorkSpace"}, expected: "/home/aeswolf/Documents/GoWorkSpace"},
		{data: []string{"Documents", "cdx/funcs"}, expected: "/home/aeswolf/Documents/GoWorkSpace/src/github.com/Aeswolf/cdx/funcs"},
		{data: []string{"Documents", "cdx"}, expected: "/home/aeswolf/Documents/GoWorkSpace/src/github.com/Aeswolf/cdx"},
		{data: []string{"Documents", "AI"}, expected: "/home/aeswolf/Documents/PythonWS/src/github.com/Aeswolf/AI"},
		{data: []string{"Documents", "machine-learning"}, expected: "/home/aeswolf/Documents/PythonWS/src/github.com/Aeswolf/AI/machine-learning"},
		{data: []string{"Documents", "funcs/all.go"}, expected: "/home/aeswolf/Documents/GoWorkSpace/src/github.com/Aeswolf/cdx/funcs/all.go"},
		{data: []string{"Documents", "cdx/main.go"}, expected: "/home/aeswolf/Documents/GoWorkSpace/src/github.com/Aeswolf/cdx/main.go"},
	}

	for _, sample := range testData {
		cdx := cd.NewCdx(sample.data[0])
		cdx.SetDest(sample.data[1])
		sample.outcome = funcs.GetPath(cdx)

		if sample.outcome != sample.expected {
			t.Errorf("FAILED!!!\nSrc : %s\nDest : %s\nOutcome : %s\nExpected : %s\n", cdx.GetSrc(), cdx.GetDest(), sample.outcome, sample.expected)
		} else {
			t.Logf("PASSED!!!\nSrc : %s\nDest : %s\nOutcome : %s\nExpected : %s\n", cdx.GetSrc(), cdx.GetDest(), sample.outcome, sample.expected)
		}

	}
}
