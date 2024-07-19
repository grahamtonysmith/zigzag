package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	x := []string{
		"2017-10-24",
		"2017-10-25",
		"2017-10-26",
		"2017-10-27",
	}

	y := []opts.KlineData{
		{Value: []int{20, 34, 10, 38}}, // open=20, close=34, low=10, high=38
		{Value: []int{40, 35, 30, 50}},
		{Value: []int{31, 38, 33, 44}},
		{Value: []int{38, 15, 05, 42}},
	}

	charts.NewKLine().SetXAxis(x).AddSeries("GC", y).Render(w)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
