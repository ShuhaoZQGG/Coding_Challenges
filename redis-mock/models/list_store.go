package models

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

type ListStore struct {
	state sync.RWMutex
	data  map[string]list.List
}

func NewListStore() *ListStore {
	return &ListStore{
		data: make(map[string]list.List),
	}
}

func (store *ListStore) Lget(key string) (list.List, bool) {
	store.state.RLock()
	defer store.state.RUnlock()
	value, ok := store.data[key]
	return value, ok
}

// /	<description>
// /	Implement Lrange function to return all the elements within the range of start and stop position
// /	</description>
// /	<params>
// /	key string: key of data in ListStore
// / start int: start position of the range
// / end   int: end position of the range
// / 0 be the first element of the list, 1 being the next element and so on
// /-1 be the last element of the list, -2 being the penultimate and so on
// /	</params>
// / <example>
// / "key": ["value1", "value2", "value3"]
// / key: "key", start: 0, stop: -1
// / return
// / (1) "value1"\r\n
// / (2) "value2"\r\n
// / (3) "value3"\r\n
// / </example>
func (store *ListStore) Lrange(key string, start int, stop int) (output []string, err error) {
	value, ok := store.Lget(key)
	if !ok {
		return nil, errors.New(fmt.Sprintf("key %s does not exist", key))
	}
	length := value.Len()

	// -1 => the last element = len + (-1) = len - 1
	if start < 0 {
		start = length + start
	}

	if stop < 0 {
		stop = length + stop
	}

	// check if start is still less than 0 (client input is problematic)
	if start < 0 {
		start = 0
	}

	if stop < start {
		start, stop = stop, start
	}

	if stop >= length {
		stop = length - 1
	}

	i := 0
	for e := value.Front(); e != nil; e = e.Next() {
		if i >= start && i <= stop {
			output = append(output, e.Value.(string))
		}

		if i > stop {
			break
		}

		i++
	}

	return output, nil
}

func (store *ListStore) Lpush(key string, values []string) {
	value, ok := store.Lget(key)
	if !ok {
		value = *list.New().Init()
	}

	for _, val := range values {
		value.PushFront(val)
	}

	store.state.Lock()
	defer store.state.Unlock()

	store.data[key] = value
}
