package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/radix", func(ctx iris.Context) {
		ctx.Write(parseAndSort([]byte("[" + ctx.URLParam("array") + "]")))
	})
	app.Post("/radix", func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		ctx.Write(parseAndSort(body))
	})
	app.Run(iris.Addr(":80"))
}

func parseAndSort(bytes []byte) []byte {
	var array []int
	json.Unmarshal(bytes, &array)

	b, _ := json.Marshal(map[string][]int{"result": sort(array)})

	return b
}

func sort(list []int) []int {
	max := arrayMax(list)
	base := 10

	for digit := 1; max/digit > 0; digit *= 10 {
		list = countingSortByDigit(list, base, digit)
	}

	return list
}

func countingSortByDigit(array []int, base int, digit int) []int {
	length := len(array)
	result := make([]int, length)
	counts := make([]int, base)

	for i := 0; i < length; i++ {
		counts[ (array[i]/digit)%base]++
	}

	for i := 1; i < base; i++ {
		counts[i] += counts[i - 1]
	}

	for i := length - 1; i >= 0; i-- {
		result[counts[(array[i]/digit)%base] - 1] = array[i]
		counts[(array[i]/digit)%base]--
	}

	return result
}

func arrayMax(array []int) int {
	length := len(array)
	if (length == 0) {
		return 0;
	}

	max := array[0]
	for i := 0; i < length; i++ {
		if max < array[i] {
			max = array[i]
		}
	}

	return max
}
