package listex

import . "container/list"

func Contains(l *List, v interface{}) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return true
		}
	}
	return false
}

func Del(l *List, v interface{}) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			l.Remove(e)
		}
	}
}
