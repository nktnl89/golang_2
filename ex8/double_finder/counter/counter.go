// Package counter provides structs and methods for searching files with duplicates.
package counter

import (
	"sync"
)

// FullFileInfo provides data for collecting file's name, full path and size
type FullFileInfo struct {
	Name     string
	Size     int64
	FullPath string
}

// DoubleCounters provides RWMutex and WaitGroup for sync access to allFiles map. Also contains doubles map
type DoubleCounters struct {
	mx       sync.RWMutex
	allFiles map[string][]FullFileInfo
	doubles  map[FullFileInfo]bool
	wg       sync.WaitGroup
}

// Constructor for DoubleCounters
func NewCounters() *DoubleCounters {
	return &DoubleCounters{
		allFiles: make(map[string][]FullFileInfo),
		doubles:  make(map[FullFileInfo]bool),
		wg:       sync.WaitGroup{},
	}
}

// Method Find allows to find file by name from filled before allFiles map. Works with read lock.
func (c *DoubleCounters) Find(key string) ([]FullFileInfo, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.allFiles[key]
	return val, ok
}

// Method FindAll allows to get all files from filled before allFiles map. Works with read lock.
func (c *DoubleCounters) FindAll() map[string][]FullFileInfo {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.allFiles
}

// Method Add allows to find file by name from filled before allFiles map. Works with read lock.
func (c *DoubleCounters) Add(key FullFileInfo) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.allFiles[key.Name] = append(c.allFiles[key.Name], key)
}

// Method FindDoubles allows to find duplicate files from filled before allFiles map. Works with read lock.
func (c *DoubleCounters) FindDoubles() map[FullFileInfo]bool {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for _, f := range c.allFiles {
		if len(f) > 1 {
			first := f[0]
			for i := 1; i < len(f); i++ {
				if first.Size == f[i].Size { // Name не проверяю т.к. и так совпадает
					if !c.doubles[first] {
						c.doubles[first] = true
					}
					c.doubles[f[i]] = true
				}
			}
		}
	}
	return c.doubles
}

// Method GetDoubles - getter for map with doubles
func (c *DoubleCounters) GetDoubles() map[FullFileInfo]bool {
	return c.doubles
}

// Method WaitGroupAdd - helper for adding delta to wait group
func (c *DoubleCounters) WaitGroupAdd(delta int) {
	c.wg.Add(delta)
}

// Method WaitGroupAdd - helper for done of wait group
func (c *DoubleCounters) WaitGroupDone() {
	c.wg.Done()
}

// Method WaitGroupAdd - helper for wating of wait group
func (c *DoubleCounters) WaitGroupWait() {
	c.wg.Wait()
}
