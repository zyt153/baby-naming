package naming

import (
    "math/rand"
    "time"
	"sort"
	"fmt"
)

var maleNamesMap = map[int] []string {3: []string{"Tim", "Bob"}, 4: []string{"Mike", "Nick"}, 5: []string{"Peter", "Harry", "Denny"}, 6: []string{"Andrew", "Newman"}}
var femaleNamesMap = map[int] []string {3: []string{"Ann", "Ida"}, 4: []string{"Ella", "Jane", "Vivi"}, 5: []string{"Jenny", "Kelly"}, 6: []string{"Simona"}}


func CreateBabyName(male bool, minLen int) (string, error) {
    rand.Seed(time.Now().Unix())

    if male {
		names, err := FindByLength(maleNamesMap, minLen)
		if err == nil {
			index := rand.Intn(len(names))
			return names[index], nil
		} else {
			return "", err
		}
	} else {
		names, err := FindByLength(femaleNamesMap, minLen)
		if err == nil {
			index := rand.Intn(len(names))
			return names[index], nil
		} else {
			return "", err
		}
	}
}

func FindByLength(namesMap map[int][]string, minLen int) ([]string, error) {
	rand.Seed(time.Now().Unix())
	var keys []int
	for key := range namesMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	if minLen > keys[0] {
		return []string{}, fmt.Errorf("No matching name found with min length %v", minLen)
	} else {
		index := rand.Intn(keys[0] - minLen + 1)  + minLen
		return namesMap[index], nil
	}
}