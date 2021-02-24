package main

import (
	"io/ioutil"
	"testing"
)

func testTempFile(t *testing.T) string {
	t.Helper()
	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal("err %s", err)
	}
	tf.Close()
	return tf.Name()
}

/**
 * サブテストを指定して実行
 * go test -V sample -run TestAB/A
 */
func TestAB(t *testing.T) {
	t.Run("A", func(t * testing.T) { t.Error("error") })
	t.Run("B", func(t * testing.T) { t.Error("error") })
}

var flagTests = []struct {
	in string
	out string
} {
	{"%a", "[%a]"}, {"%-a", "[%-a]"}, {"%+a", "[%+a]"},
	{"%#a", "[%#a]"}, {"% a", "[% a]"},
}
