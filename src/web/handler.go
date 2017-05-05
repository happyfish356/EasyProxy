package web

import (
	"net/http"
	"proxy"
)

const (
	StatisticsUrl = "/statistic"
)

func Statistic(writer http.ResponseWriter, request *http.Request) {
	proxy.Record()
	Render(writer, "statistic", StatisticHtml, proxy.StatisticData())
}
