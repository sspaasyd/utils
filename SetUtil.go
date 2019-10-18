package utils

import (
	"sync"
)

type Set struct {
	m map[string]bool
	sync.RWMutex
}

//创建Set集合
func NewSet() *Set {
	return &Set{
		m: map[string]bool{},
	}
}

//添加元素
func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

//删除元素
func (s *Set) Remove(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

//判断是否存在元素
func (s *Set) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

//Set长度
func (s *Set) Len() int {
	return len(s.List())
}

//清除Set
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

//判断Set是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

//转为切片集合
func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

//如果map集合key为int的话，支持此方法排序
/*func (s *Set) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}*/
