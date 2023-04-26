package strings

import (
	StringsPackage "github.com/Eclalang/strings"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	str := "Hello, World!"
	substr := "World"
	expected := strings.Contains(str, substr)
	actual := StringsPackage.Contains(str, substr)
	if actual != expected {
		t.Errorf("StringsPackage.Contains(\"%s\", \"%s\") returned %t, expected %t", str, substr, actual, expected)
	}
}

func TestContainsAny(t *testing.T) {
	str := "Hello, World!"
	substr := "l"
	expected := strings.ContainsAny(str, substr)
	actual := StringsPackage.ContainsAny(str, substr)
	if actual != expected {
		t.Errorf("StringsPackage.ContainsAny(\"%s\", \"%s\") returned %t, expected %t", str, substr, actual, expected)
	}
}

func TestCount(t *testing.T) {
	str := "Hello, World!"
	substr := "l"
	expected := strings.Count(str, substr)
	actual := StringsPackage.Count(str, substr)
	if actual != expected {
		t.Errorf("StringsPackage.Count(\"%s\", \"%s\") returned %d, expected %d", str, substr, actual, expected)
	}
}

func TestCut(t *testing.T) {
	str := "Hello i am a bear"
	substr := "a "
	expectedBefore, expectedAfter, expectedFound := strings.Cut(str, substr)
	actual := StringsPackage.Cut(str, substr)
	if actual[0] != expectedBefore && actual[1] != expectedAfter && actual[2] != strconv.FormatBool(expectedFound) {
		t.Errorf("StringsPackage.Cut(\"%s\", \"%s\") returned %s, expected %s", str, substr, actual, [3]string{expectedBefore, expectedAfter, strconv.FormatBool(expectedFound)})
	}
}

func TestHasPrefix(t *testing.T) {
	str := "Counter the strike"
	substr := "Count"
	expected := strings.HasPrefix(str, substr)
	actual := StringsPackage.HasPrefix(str, substr)
	if actual != expected {
		t.Errorf("StringsPackage.HasPrefix(\"%s\", \"%s\") returned %t, expected %t", str, substr, actual, expected)
	}
}

func TestHasSuffix(t *testing.T) {
	str := "Counter the strike"
	substr := "strike"
	expected := strings.HasSuffix(str, substr)
	actual := StringsPackage.HasSuffix(str, substr)
	if actual != expected {
		t.Errorf("StringsPackage.HasSuffix(\"%s\", \"%s\") returned %t, expected %t", str, substr, actual, expected)
	}
}

func TestIndexOf(t *testing.T) {
	str := "Hello, World!"
	substr := "e"
	expected := strings.Index(str, substr)
	actual := StringsPackage.IndexOf(str, substr)
	if actual != expected {
		t.Errorf("StringsPackage.IndexOf(\"%s\", \"%s\") returned %d, expected %d", str, substr, actual, expected)
	}
}

func TestJoin(t *testing.T) {
	elems := []string{"Hello", "I am", "a", "doctor"}
	sep := " "
	expected := strings.Join(elems, sep)
	actual := StringsPackage.Join(elems, sep)
	if actual != expected {
		t.Errorf("StringsPackage.Join(\"%s\", \"%s\") returned %s, expected %s", elems, sep, actual, expected)
	}
}
func TestReplace(t *testing.T) {
	str := "Hello, World!"
	old := "World"
	new := "Eclalang"
	expected := strings.Replace(str, old, new, -1)
	actual := StringsPackage.Replace(str, old, new, -1)
	if actual != expected {
		t.Errorf("StringsPackage.Replace(\"%s\", \"%s\", \"%s\", -1) returned %s, expected %s", str, old, new, actual, expected)
	}
}

func TestReplaceAll(t *testing.T) {
	str := "Hello, World!"
	old := "World"
	new := "Ecla"
	expected := strings.ReplaceAll(str, old, new)
	actual := StringsPackage.ReplaceAll(str, old, new)
	if actual != expected {
		t.Errorf("StringsPackage.ReplaceAll(\"%s\", \"%s\", \"%s\") returned %s, expected %s", str, old, new, actual, expected)
	}
}

func TestSplit(t *testing.T) {
	str := "Hello, World!"
	sep := " "
	expected := strings.Split(str, sep)
	actual := StringsPackage.Split(str, sep)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("StringsPackage.Split(\"%s\", \"%s\") returned %s, expected %s", str, sep, actual, expected)
	}
}

func TestSplitAfter(t *testing.T) {
	str := "Hello, World!"
	sep := " "
	expected := strings.SplitAfter(str, sep)
	actual := StringsPackage.SplitAfter(str, sep)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("StringsPackage.SplitAfter(\"%s\", \"%s\") returned %s, expected %s", str, sep, actual, expected)
	}
}

func TestSplitAfterN(t *testing.T) {
	str := "Hello, World!"
	sep := " "
	n := 2
	expected := strings.SplitAfterN(str, sep, n)
	actual := StringsPackage.SplitAfterN(str, sep, n)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("StringsPackage.SplitAfterN(\"%s\", \"%s\", %d) returned %s, expected %s", str, sep, n, actual, expected)
	}
}

func TestSplitN(t *testing.T) {
	str := "Hello, World!"
	sep := " "
	n := 2
	expected := strings.SplitN(str, sep, n)
	actual := StringsPackage.SplitN(str, sep, n)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("StringsPackage.SplitN(\"%s\", \"%s\", %d) returned %s, expected %s", str, sep, n, actual, expected)
	}
}

func TestToLower(t *testing.T) {
	str := "HahahahaBlaBlaBla"
	expected := strings.ToLower(str)
	actual := StringsPackage.ToLower(str)
	if actual != expected {
		t.Errorf("StringsPackage.ToLower(\"%s\") returned %s, expected %s", str, actual, expected)
	}
}

func TestToUpper(t *testing.T) {
	str := "HahahahaBlaBlaBla"
	expected := strings.ToUpper(str)
	actual := StringsPackage.ToUpper(str)
	if actual != expected {
		t.Errorf("StringsPackage.ToUpper(\"%s\") returned %s, expected %s", str, actual, expected)
	}
}

func TestTrim(t *testing.T) {
	str := "HahahahaBlaBlaBla"
	cutset := "Hahahaha"
	expected := strings.Trim(str, cutset)
	actual := StringsPackage.Trim(str, cutset)
	if actual != expected {
		t.Errorf("StringsPackage.Trim(\"%s\", \"%s\") returned %s, expected %s", str, cutset, actual, expected)
	}
}
