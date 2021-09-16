package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {}

// =================================
// I/O util
// =================================

const (
	bufSize = 1000000
)

var (
	s *bufio.Scanner
	w *bufio.Writer
)

func init() {
	s = bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, bufSize), bufSize)
	w = bufio.NewWriterSize(os.Stdout, bufSize)
}

func splitSpace() {
	s.Split(bufio.ScanWords)
}

func next() string {
	s.Scan()
	return s.Text()
}

func nextInt() int {
	return stringToInt(next())
}

func write(result ...interface{}) {
	fmt.Fprintln(w, result...)
	w.Flush()
}

// =================================
// Util
// =================================

func stringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
