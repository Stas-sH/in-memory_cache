package inmemorycache

import (
	"errors"
	"time"
)

type Memory struct {
	memoryValue map[string]interface{}
	keys        []string
	ttl         []time.Duration
	timeCreat   []time.Time
}

func NewMemory() Memory {

	x := Memory{
		memoryValue: make(map[string]interface{}),
		keys:        make([]string, 0),
		ttl:         make([]time.Duration, 0),
		timeCreat:   make([]time.Time, 0),
	}
	return x
}

func (m *Memory) Set(key string, value interface{}, timeToLive time.Duration) error {
	if _, inMap := m.memoryValue[key]; inMap {
		return errors.New("key is working yet")
	}
	m.keys = append(m.keys, key)
	m.ttl = append(m.ttl, timeToLive)
	m.timeCreat = append(m.timeCreat, time.Now())

	m.memoryValue[key] = value
	return nil
}

func (m *Memory) Get(key string) (interface{}, error) {

	if value, inMap := m.memoryValue[key]; inMap {
		indxElem := m.serchIndexKey(key)

		if time.Since(m.timeCreat[indxElem]) > (m.ttl[indxElem]) {
			m.Delete(key)
			return nil, errors.New("time is ower, it is deleting now")
		}
		return value, nil
	}
	return nil, errors.New("key not found")
}

func (m *Memory) Delete(key string) {
	delete(m.memoryValue, key)
}

func (m *Memory) serchIndexKey(key string) int {
	var indexKey int
	for i := 0; i < len(m.keys); i++ {
		if m.keys[i] == key {
			indexKey = i
		}
	}
	return indexKey
}
