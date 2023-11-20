package berkatbepkg

import (
	"math/rand"
	"time"
)

func randomnumber() {
	source := rand.NewSource(time.Now().UnixNano())
	rand_source := rand.New(source)
	for i := 0; i < 5; i++ {
		rand_num := rand_source.Int()
	}
}
