package naming

import (
    "math/rand"
    "time"
)

var maleNames = []string{"Mike", "Tim", "Andrew", "Jianguo"}
var femaleNames = []string{"Elisabeth", "Jenifer", "Amy", "Yifan"}

func CreateBabyName(male bool) string {
    rand.Seed(time.Now().Unix())

    if male {
		index := rand.Intn(len(maleNames))
		return maleNames[index]
	} else {
		index := rand.Intn(len(femaleNames))
		return femaleNames[index]
	}
}