package main

func producer(done <-chan interface{}, objects []Object, batchSize int) <-chan Object {
	ch := make(chan Object, batchSize)

	go func() {
		defer close(ch)
		for _, obj := range objects {
			select {
			case <-done:
				return
			case ch <- obj:
			}

		}
	}()

	return ch
}
