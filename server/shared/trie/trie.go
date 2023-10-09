package trie

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var (
	// EntryNotPresentErr indicates that the Trie doesn't contain the given entry
	EntryNotPresentErr = errors.New("trie: entry not present")
	// IsFullErr indicates that the Trie can't store more text entries
	IsFullErr = errors.New("trie: is full")
	// IsEmptyErr indicates that the Trie has no entries
	IsEmptyErr = errors.New("trie: is empty")
)

// Trie is a data structure that store text entries for quick search
type Trie struct {
	root  *node
	size  uint
	count uint
	sem   chan bool
}

// NewTrie returns a Trie struct
func NewTrie(size uint, semSize uint) Trie {
	return Trie{
		root:  newNode(""),
		size:  size,
		count: 0,
		sem:   make(chan bool, semSize),
	}
}

// AddEntry adds an entry to the Trie
func (t *Trie) AddEntry(entry string) error {
	if t.size == t.count {
		return IsFullErr
	}

	current := t.root

	for _, ch := range entry {
		chr := string(ch)

		if !current.hasChild(chr) {
			current.addChild(chr)
		}

		current = current.getChild(chr)
	}

	if current != t.root {
		current.end = true
		t.count++
	}

	return nil
}

// RemoveEntry removes the given entry from the Trie.
//
// If the Trie doesn't contain the given entry returns an error
func (t *Trie) RemoveEntry(entry string) error {
	if t.count == 0 {
		return IsEmptyErr
	}

	return t.removeEntry(t.root, entry, 0)
}

// removeEntry is a recursive implementation of RemoveEntry
func (t *Trie) removeEntry(root *node, entry string, i uint) error {
	if i == uint(len(entry)) {
		root.end = false
		t.count--

		return nil
	}

	var (
		chr   = string(entry[i])
		child = root.getChild(chr)
	)

	if child == nil {
		return EntryNotPresentErr
	}

	err := t.removeEntry(child, entry, i+1)
	if err != nil {
		return err
	}

	if !child.hasChildren() && !child.end {
		root.removeChild(chr)
		child = nil
	}

	return nil
}

// FindEntries returns all entries in the Trie that matches the given prefix.
//
// If the Trie doesn't contain entries matching returns an error
func (t *Trie) FindEntries(prefix string, limit int) ([]string, error) {
	if t.count == 0 {
		return []string{}, IsEmptyErr
	}

	prefixNode, err := t.findEndNodeOf(prefix)
	if err != nil {
		return []string{}, err
	}

	return t.findEntries(prefixNode, prefix, limit), nil
}

// findEndNodeOf finds the last node in the Trie related to the given entry.
//
// If the Trie doesn't contain entries matching returns an error.
func (t *Trie) findEndNodeOf(entry string) (*node, error) {
	return t.recursivelyFindEndNodeOf(t.root, entry, 0)
}

// recursivelyFindEndNodeOf is a recursive implementation of findEndNodeOf
func (t *Trie) recursivelyFindEndNodeOf(root *node, entry string, i uint) (*node, error) {
	if i == uint(len(entry)) {
		return root, nil
	}

	var (
		chr   = string(entry[i])
		child = root.getChild(chr)
	)

	if child == nil {
		return nil, EntryNotPresentErr
	}

	return t.recursivelyFindEndNodeOf(child, entry, i+1)
}

// findEntries searches depth first (recursively) the Trie from the given node to retrieve all entries.
//
// This is probably overengineered using goroutines, but is awesome
func (t *Trie) findEntries(root *node, prefix string, limit int) []string {
	var (
		entries     = make([]string, 0, limit)
		out         = make(chan string, limit)
		wg          = new(sync.WaitGroup)
		ctx, cancel = context.WithCancel(context.TODO())
	)

	wg.Add(1)

	go t.recursivelyFindEntries(ctx, wg, out, root, prefix)

	go func(wg *sync.WaitGroup, out chan string) {
		wg.Wait()
		close(out)
	}(wg, out)

	for entry := range out {
		if len(entries) == limit {
			cancel() // entries limit has been reached. Context is canceled to prevent more work to be done.
			continue
		}

		entries = append(entries, entry)
	}

	cancel() // in case all work has been finished before reaching entries limit. Context is cancelled for cleaning duties

	return entries
}

// recursivelyFindEntries is a recursive implementation of findEntries
func (t *Trie) recursivelyFindEntries(ctx context.Context, wg *sync.WaitGroup, out chan<- string, root *node, entry string) {
	defer wg.Done()

	select {
	case <-ctx.Done():
	default:
		t.sem <- true
		defer func(sem <-chan bool) {
			<-sem
		}(t.sem)

		if root.end {
			out <- entry
		}

		if !root.hasChildren() {
			return
		}

		wg.Add(len(root.children))

		for _, child := range root.children {
			go t.recursivelyFindEntries(ctx, wg, out, child, fmt.Sprintf("%s%s", entry, child.value))
		}
	}
}
