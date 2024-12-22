package day22

import (
	"container/heap"
	"fmt"
	"os"
	"strings"

	"github.com/winchest3r/aoc2024/utils"
)

type Pair struct {
	Secret int
	Diff   int
}

type Price struct {
	Value int
	Seq   []int
}

type PriceHeap []Price

func (h PriceHeap) Len() int {
	return len(h)
}

func (h PriceHeap) Less(i, j int) bool {
	return h[i].Value > h[j].Value
}

func (h PriceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PriceHeap) Push(p any) {
	*h = append(*h, p.(Price))
}

func (h *PriceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CalculateNextSecret(secret int) int {
	first := secret * 64
	secret = (secret ^ first) % 16777216
	second := secret / 32
	secret = (secret ^ second) % 16777216
	third := secret * 2048
	secret = (secret ^ third) % 16777216
	return secret
}

func CalculateNSecretsWithDiff(init, n int) []Pair {
	result := make([]Pair, n)
	secret := CalculateNextSecret(init)
	result[0] = Pair{secret, secret%10 - init%10}
	for i := 1; i < n; i++ {
		secret = CalculateNextSecret(secret)
		result[i] = Pair{secret, secret%10 - result[i-1].Secret%10}
	}
	return result
}

func CalculateNBestChanges(secret, secN, n int) []Price {
	pq := &PriceHeap{}
	heap.Init(pq)
	data := CalculateNSecretsWithDiff(secret, secN)
	diffs := []int{data[0].Diff, data[0].Diff, data[1].Diff, data[2].Diff}
	for i := 3; i < len(data); i++ {
		diffs = diffs[1:]
		diffs = append(diffs, data[i].Diff)
		price := data[i].Secret % 10
		heap.Push(pq, Price{price, diffs})
	}
	result := []Price{}
	for i := 0; i < n; i++ {
		price := heap.Pop(pq).(Price)
		result = append(result, price)
	}
	return result
}

func CalculateNSecret(secret, n int) int {
	result := secret
	for i := 0; i < n; i++ {
		result = CalculateNextSecret(result)
	}
	return result
}

func ReadInput(fname string) []int {
	file, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return utils.FieldsToInt(string(file))
}

func SolvePartOne(fname string) {
	buyers := ReadInput(fname)
	result := 0
	for _, b := range buyers {
		result += CalculateNSecret(b, 2000)
	}
	fmt.Println(result)
}

type Buyer struct {
	Id    int
	Price int
}

func GetAllSeqPrices(buyers []int, secN, bestN int) map[string][]Buyer {
	result := make(map[string][]Buyer)
	for id, secret := range buyers {
		prices := CalculateNBestChanges(secret, secN, bestN)
		for _, p := range prices {
			diffStr := strings.Join(utils.CastToString(p.Seq), "")
			_, ok := result[diffStr]
			if !ok {
				result[diffStr] = make([]Buyer, 0)
			}
			result[diffStr] = append(result[diffStr], Buyer{id, p.Value})
		}
	}
	return result
}

func GetBestPrice(buyers []int, secN int) int {
	allPrices := GetAllSeqPrices(buyers, secN, secN-3)
	result := 0
	for _, prices := range allPrices {
		best := make(map[int]int)
		for _, p := range prices {
			best[p.Id] = max(best[p.Id], p.Price)
		}
		sum := 0
		for _, val := range best {
			sum += val
		}
		result = max(result, sum)
	}
	return result
}

func SolvePartTwo(fname string) {
	buyers := ReadInput(fname)
	fmt.Println(GetBestPrice(buyers, 2000))
}
