package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

type Case struct {
	inputKey string
	inputVal []byte
}

func NewCases() []Case {
	return []Case{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)

	cases := NewCases()

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", actual)
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inputVal))
		}
	}

}

func TestClearCacheLoop(t *testing.T) {

	interval := time.Millisecond * 10
	c := NewCache(interval)

	cases := NewCases()

	for _, cas := range cases {
		c.Add(cas.inputKey, cas.inputVal)
	}

	time.Sleep(interval + time.Millisecond)

	for _, cas := range cases {
		_, ok := c.Get(cas.inputKey)
		if ok {
			t.Errorf("%s should not be in Cache", cas.inputKey)
		}

	}

}

func TestNotClearCacheLoop(t *testing.T) {

	interval := time.Millisecond * 10
	c := NewCache(interval)

	cases := NewCases()

	for _, cas := range cases {
		c.Add(cas.inputKey, cas.inputVal)
	}

	time.Sleep(interval / 2)

	for _, cas := range cases {
		_, ok := c.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s should be in Cache", cas.inputKey)
		}

	}

}
