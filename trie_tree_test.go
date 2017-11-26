package trie

import (
	"testing"
)

func TestAddEntry(t *testing.T) {
	tr := NewTrieTree()
	tr.Add("test")
	exist := tr.Has("test")
	if !exist {
		t.Error("'test' should exist")
	}

	exist = tr.Has("tes")
	if exist {
		t.Error("'tes' should not exist")
	}

	tr.Add("tes")
	exist = tr.Has("test")
	if !exist {
		t.Error("'test' should exist")
	}

	exist = tr.Has("tes")
	if !exist {
		t.Error("'tes' should exist now")
	}
}

func TestDump(t *testing.T) {
	tr := NewTrieTree()
	tr.Add("test")
	tr.Add("teacher")

	tr.Dump()
}

func TestDelete(t *testing.T) {
	tr := NewTrieTree()
	tr.Add("test")
	tr.Add("teacher")

	tr.Delete("teacher")
	tr.Delete("test")

	tr.Dump()
}
