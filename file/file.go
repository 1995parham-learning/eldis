package file

import (
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
		data, err := ioutil.ReadFile("./foo/" + fileNames[i])
		if err != nil {
			log.Fatal(err)
		}

		textData := string(data)

		lines := strings.Split(textData, "\n")

		for _, line := range lines {
			keyValue := strings.Split(line, " ")
			if keyValue[0] == key {
				return keyValue[1]
			}
		}

	}

	return ""
}

func binarySearch(key string, lines []string) string {
	low := 0
	high := len(lines) - 1

	for {
		middle := (low + high) / 2

		kv := strings.Split(lines[middle], " ")
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