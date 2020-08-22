package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	l := log.New(os.Stderr, "", 0)

	objects, err := randomObjects("testdata/random_objects.json")
	if err != nil {
		l.Fatal(err)
	}

	doneProducer := make(chan interface{})
	defer close(doneProducer)

	batchSize := 100
	objectsCh := producer(doneProducer, objects, batchSize)

	done := make(chan interface{})
	result := Result{data: Objects{}, mtx: &sync.Mutex{}}
	consumer(done, objectsCh, &result, batchSize)

	<-done

	resultFilePath, err := writeResult2File(result.data)
	if err != nil {
		l.Fatal(err)
	}

	fmt.Printf("You can check result here: %s\n", resultFilePath)
}
