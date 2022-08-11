package common

import "sync"

var (
	Count = 0
	lock  = sync.Mutex{}
)

func I() {
	lock.Lock()
	defer lock.Unlock()
	Count++
}

func D() {
	lock.Lock()
	defer lock.Unlock()
	Count--
}
