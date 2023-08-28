package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/joho/godotenv"
)

var (
	months     = []string{"1996-07", "1996-08", "1996-09", "1996-10", "1996-11", "1996-12", "1997-01"}
	categories = []string{
		"Beverages",
		"Condiments",
		"Confections",
		"Dairy Products",
		"Grains/Cereals",
		"Meat/Poultry",
	}
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < len(months); i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func barStack() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Order Trend",
		}),
	)
	bar.SetXAxis(months)
	for _, v := range categories {
		bar.AddSeries(v, generateBarItems())
	}
	bar.SetSeriesOptions(charts.WithBarChartOpts(opts.BarChart{
		Stack: "stackA",
	}))

	return bar
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	page := components.NewPage()
	page.AddCharts(barStack())
	f, err := os.Create("html/bar.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.Writer(f))

	fs := http.FileServer(http.Dir("html"))
	log.Printf("running server at http://%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")), logRequest(fs)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
