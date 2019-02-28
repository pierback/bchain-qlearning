package blockchain

import "sync"

type Queue struct {
	lock   *sync.Mutex
	Values []bchainData
}

func (q *Queue) Enqueue(x bchainData) {
	for {
		q.lock.Lock()
		q.Values = append(q.Values, x)
		q.lock.Unlock()
		return
	}
}

func (q *Queue) Dequeue() *bchainData {
	for {
		if len(q.Values) > 0 {
			q.lock.Lock()
			x := q.Values[0]
			q.Values = q.Values[1:]
			q.lock.Unlock()
			return &x
		}
		return nil
	}
}
