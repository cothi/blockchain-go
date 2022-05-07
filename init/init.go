package unit

import "sync"

type Component struct {
	count int
}

var instance *Component
var once sync.Once

func GetInstance() *Component {
	once.Do(func() {
		instance = &Component{count: 0}
	})
	return instance
}

func (s *Component) add() {
	s.count++
}

func (s *Component) getCount() int {
	return s.count
}
