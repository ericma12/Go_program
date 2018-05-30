package main

import (
	"fmt"
	"math/rand"
	"time"
)

const testNum = 50
const maxSize = 20000000
const findMutx = 1 //99999999

type lst []int

func (l lst) has(xx int) bool {
	for _, i := range l {
		if i == xx {
			return true
		}

	}
	return false
}

type mp map[int]struct{}

func (m mp) has(xxx int) bool {
	_, ok := m[xxx]
	return ok
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for size := 1; size < testNum; size++ {
		l := make(lst, 0, size)
		m := make(mp, size)
		for i := 0; i < size; i++ {
			e := rand.Int()
			l = append(l, e)
			m[e] = struct{}{}
		}

		found := 0

		start := time.Now()
		for x := 0; x < maxSize; x++ {
			xx := rand.Intn(size * findMutx)
			if l.has(xx) {
				found++
			}
		}
		sliceDuration := time.Now().Sub(start)

		start = time.Now()
		for x := 0; x < maxSize; x++ {
			xx := rand.Intn(size * findMutx)
			if m.has(xx) {
				found++
			}
		}
		mapDuration := time.Now().Sub(start)

		if found > 0 {
			rand.Int()
		}

		fmt.Printf("%d\t%v\n", size, sliceDuration-mapDuration)
	}

}
