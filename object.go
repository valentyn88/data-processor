package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"sync"
	"time"
)

// Object main object for processing.
type Object struct {
	ID   string `json:"id"`
	Seq  int64  `json:"seq"`
	Data string `json:"data"`
}

// Objects set of Object.
type Objects []Object

// Len returns length of the slice.
func (a Objects) Len() int {
	return len(a)
}

// Less checks whether Object.Seq is less or not.
func (a Objects) Less(i, j int) bool {
	return a[i].Seq < a[j].Seq
}

// Swap swaps elements in slice.
func (a Objects) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Result represents final result.
type Result struct {
	mtx  *sync.Mutex
	data Objects
}

// Add adds elements to the result.
func (s *Result) Add(objects []Object) {
	s.mtx.Lock()
	s.data = append(s.data, objects...)
	sort.Sort(s.data)
	s.mtx.Unlock()
}

func randomObjects(filePath string) ([]Object, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []Object{}, err
	}

	objects := []Object{}
	err = json.Unmarshal(content, &objects)
	if err != nil {
		return []Object{}, err
	}

	return objects, nil
}

func writeResult2File(objects Objects) (string, error) {
	data, err := json.Marshal(objects)
	if err != nil {
		return "", err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "\t"); err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("results/result_%d.json", time.Now().UnixNano())
	if err := ioutil.WriteFile(filePath, prettyJSON.Bytes(), 0644); err != nil {
		return "", err
	}

	return filePath, nil
}
