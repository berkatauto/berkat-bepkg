package berkatbepkg

import (
	"math/rand"
	"time"
)

func randomUserId(deploynumber int, toID RandomNumber) {
	source := rand.NewSource(time.Now().UnixNano())
	rand_source := rand.New(source)
	for i := 0; i < 5; i++ {
		rand_num := rand_source.Int()
		deploynumber = rand_num
	}
	toID = RandomNumber{random: deploynumber}
}

func randomArticleID(deploynumber int, toID RandomNumber) {
	source := rand.NewSource(time.Now().UnixNano())
	rand_source := rand.New(source)
	for i := 0; i < 5; i++ {
		rand_num := rand_source.Int()
		deploynumber = rand_num
	}
	toID = RandomNumber{random: deploynumber}
}
