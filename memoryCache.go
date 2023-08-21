package inmemorycache

import "errors"

type Memory struct {
	memoryValue map[string]interface{}
}

func (m *Memory) Set(key string, value interface{}) {
	m.memoryValue[key] = value
}

func (m *Memory) Get(key string) (interface{}, error) {
	if value, inMap := m.memoryValue[key]; inMap {
		return value, nil
	}
	return nil, errors.New("key not found")
}

func (m *Memory) Delete(key string) {
	delete(m.memoryValue, key)
}
