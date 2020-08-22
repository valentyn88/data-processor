package main

import (
	"reflect"
	"sync"
	"testing"
)

func TestConsumerProducer(t *testing.T) {
	batchSize := 2

	doneProducer := make(chan interface{})
	defer close(doneProducer)

	randomObjects := []Object{
		{ID: "5", Seq: 5, Data: "Data for ID 5"},
		{ID: "2", Seq: 2, Data: "Data for ID 2"},
		{ID: "4", Seq: 4, Data: "Data for ID 4"},
		{ID: "1", Seq: 1, Data: "Data for ID 1"},
		{ID: "3", Seq: 3, Data: "Data for ID 3"},
		{ID: "0", Seq: 0, Data: "Data for ID 0"},
	}

	objectsCh := producer(doneProducer, randomObjects, batchSize)

	done := make(chan interface{})
	result := Result{data: Objects{}, mtx: &sync.Mutex{}}
	consumer(done, objectsCh, &result, batchSize)

	<-done

	expected := Objects{
		{ID: "0", Seq: 0, Data: "Data for ID 0"},
		{ID: "1", Seq: 1, Data: "Data for ID 1"},
		{ID: "2", Seq: 2, Data: "Data for ID 2"},
		{ID: "3", Seq: 3, Data: "Data for ID 3"},
		{ID: "4", Seq: 4, Data: "Data for ID 4"},
		{ID: "5", Seq: 5, Data: "Data for ID 5"},
	}

	if !reflect.DeepEqual(expected, result.data) {
		t.Errorf("expected %v\n and got %v\n are no equal", expected, result.data)
	}
}
