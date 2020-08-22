package main

func consumer(done chan<- interface{}, ch <-chan Object, result *Result, batchSize int) {
	objects := []Object{}
	go func() {
		defer close(done)
		for {
			select {
			case obj, ok := <-ch:
				if !ok {
					if len(objects) > 0 {
						result.Add(objects)
					}
					return
				}
				objects = append(objects, obj)
				if batchSize == len(objects) {
					result.Add(objects)
					objects = []Object{}
				}
			default:
			}
		}
	}()
}
