package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	if cache.entries == nil {
		t.Errorf("expected entries map to be initialized, but got nil")
	}
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "https://empty.com",
			val: []byte{},
		},
		{
			key: "https://nil.com",
			val: nil,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestAddGetOverwrite(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key          string
		val          []byte
		overwriteVal []byte
	}{
		{
			key:          "https://example.com",
			val:          []byte("testdata"),
			overwriteVal: []byte("overwritten"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			cache.Add(c.key, c.overwriteVal)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.overwriteVal) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestAddGetNotOverwrite(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key          string
		val          []byte
		overwriteVal []byte
	}{
		{
			key:          "https://example.com",
			val:          []byte("testdata"),
			overwriteVal: []byte("overwritten"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			cache.Add(c.key, c.overwriteVal)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) == string(c.val) {
				t.Errorf("expected value to have been overwriten")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestReapLoopMultiple(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))
	cache.Add("https://example.com/path?offset=0&limit=20", []byte("testdatawithpathandquery"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://example.com/path?offset=0&limit=20")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	_, ok = cache.Get("https://notexample.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
	_, ok = cache.Get("https://example.com/path?offset=0&limit=20")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
