package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	objects, err := randomObjects("testdata/random_objects.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("You can check result here: %s\n", resultFilePath)
}
