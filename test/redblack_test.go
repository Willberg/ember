package test

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"testing"
)

type FoodRatings struct {
	foodRating  map[string]int
	foodCuisine map[string]string
	cuisineMap  map[string]*redblacktree.Tree
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodRating, foodCuisine, cuisineMap := map[string]int{}, map[string]string{}, map[string]*redblacktree.Tree{}
	var comparator func(a, b interface{}) int
	comparator = func(a, b interface{}) int {
		f1, f2 := a.(string), b.(string)
		r1, r2 := foodRating[f1], foodRating[f2]
		if r1 != r2 {
			return r2 - r1
		}
		m, n := len(f1), len(f2)
		for i := 0; i < m && i < n; i++ {
			if f1[i] != f2[i] {
				return int(f1[i]) - int(f2[i])
			}
		}
		return n - m
	}
	for i, f := range foods {
		foodRating[f] = ratings[i]
		foodCuisine[f] = cuisines[i]
		tree, ok := cuisineMap[cuisines[i]]
		if !ok {
			tree = redblacktree.NewWith(comparator)
			cuisineMap[cuisines[i]] = tree
		}
		tree.Put(f, true)
	}
	return FoodRatings{foodRating: foodRating, foodCuisine: foodCuisine, cuisineMap: cuisineMap}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine, _ := this.foodCuisine[food]
	tree, _ := this.cuisineMap[cuisine]
	tree.Remove(food)
	this.foodRating[food] = newRating
	tree.Put(food, true)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.cuisineMap[cuisine].Left().Key.(string)
}

func TestMyComparator(t *testing.T) {
	f := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	fmt.Println(f.HighestRated("japanese"))
}
