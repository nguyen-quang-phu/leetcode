package leetcode

// @leet start

import "container/heap"

type FoodToCuisine struct {
	cuisine string
	rating  int
}

type CuisineToFood struct {
	food   string
	rating int
}

type CuisineToFoodHeap []CuisineToFood

func (h *CuisineToFoodHeap) Len() int {
	return len(*h)
}
func (h *CuisineToFoodHeap) Less(i, j int) bool {
	if (*h)[i].rating == (*h)[j].rating {
		return (*h)[i].food < (*h)[j].food
	}
	return (*h)[i].rating > (*h)[j].rating
}

func (h *CuisineToFoodHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *CuisineToFoodHeap) Push(x interface{}) {
	*h = append(*h, x.(CuisineToFood))
}

func (h *CuisineToFoodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *CuisineToFoodHeap) Peek() interface{} {
	return (*h)[0]
}

type FoodRatings struct {
	FoodToCuisineMap map[string]FoodToCuisine
	CuisineToFoodMap map[string]*CuisineToFoodHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodToCuisineMap := make(map[string]FoodToCuisine)
	cuisineToFoodMap := make(map[string]*CuisineToFoodHeap)
	n := len(foods)
	for i := 0; i < n; i++ {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]
		foodToCuisineMap[food] = FoodToCuisine{
			cuisine: cuisine,
			rating:  rating,
		}
		if _, ok := cuisineToFoodMap[cuisine]; !ok {
			cuisineToFoodMap[cuisine] = &CuisineToFoodHeap{}
		}
		*cuisineToFoodMap[cuisine] = append(*cuisineToFoodMap[cuisine], CuisineToFood{
			food:   food,
			rating: rating,
		})
	}
	for _, h := range cuisineToFoodMap {
		heap.Init(h)
	}

	return FoodRatings{
		FoodToCuisineMap: foodToCuisineMap,
		CuisineToFoodMap: cuisineToFoodMap,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	foodDetails := this.FoodToCuisineMap[food]
	foodDetails.rating = newRating
    this.FoodToCuisineMap[food] = foodDetails
	heap.Push(this.CuisineToFoodMap[foodDetails.cuisine], CuisineToFood{
		food:   food,
		rating: newRating,
	})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
    h := this.CuisineToFoodMap[cuisine]
    for len(*h) > 0 {
        top := h.Peek().(CuisineToFood)
        food := top.food
        rating := top.rating

        if this.FoodToCuisineMap[food].rating != rating {
            heap.Pop(h)
            continue
        }
        return food
    }
    return ""
}
/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
// @leet end

// Keynold
