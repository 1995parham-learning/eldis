package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type File interface {
	Flush(map[string]string, []string)
	Search(string) string
}

type TextFile struct {
	Threshold int
}

func (t TextFile) Flush(m map[string]string, s []string) {
	f, err := os.Create("foo/redis" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sort.Strings(s)

	for _, k := range s {
		_, err := f.WriteString(k + " " + m[k] + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (t TextFile) Search(key string) string {
	files, err := ioutil.ReadDir("./foo")
	if err != nil {
		log.Fatal(err)
	}

	fileNames := make([]string, 0)

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	sort.Strings(fileNames)

	for i := len(fileNames) - 1; i >= 0; i-- {
		f, err := os.Open("./foo/" + fileNames[i])
		if err != nil {
			log.Fatal(err)
		}

		result := t.binarySearch(key, f)

		if result != "" {
			return result
		}

		f.Close()
	}

	return ""
}

func (t TextFile) binarySearch(key string, f *os.File) string {
	low := 0
	high := t.Threshold - 1

	for {
		middle := (low + high) / 2

		line, _, err := readLine(f, middle)
		if err != nil {
			log.Fatal(err)
		}
		kv := strings.Split(line, " ")
		k := kv[0]
		v := kv[1]

		if k == key{
			return v
		}else if k < key {
			low = middle + 1
		}else {
			high = middle - 1
		}

		if low > high {
			return ""
		}
	}
}

func readLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			return sc.Text(), lastLine, sc.Err()
		}
	}

	return line, lastLine, io.EOF
}