package lottery

import (
	"math/rand"
	"time"
)

type Weighter interface {
	Weight() int32
}

type Lottery interface {
	Draw([]Weighter) int
}

type lottery struct {
	r *rand.Rand
}

func NewDefaultLottery() Lottery {
	return &lottery{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (l lottery) Draw(weighters []Weighter) int {
	if len(weighters) == 0 {
		return -1
	}

	totalWeight := int32(0)
	for _, weighter := range weighters {
		totalWeight += weighter.Weight()
	}

	lot := l.r.Int31n(totalWeight)

	tmp := int32(0)
	for i, weighter := range weighters {
		tmp += weighter.Weight()
		if lot < tmp {
			return i
		}
	}
	return -1
}
