package api

import (
	"fmt"
	"net/http"
	"strconv"
)

func getAlbums(w http.ResponseWriter, r *http.Request) {
	yrStart := r.URL.Query().Get("start_year")
	yrEnd := r.URL.Query().Get("end_year")
	var startYear, endYear int

	if yrStart != "" {
		startYear, _ = strconv.Atoi(yrStart)
	}

	if yrEnd != "" {
		endYear, _ = strconv.Atoi(yrEnd)
	}

	fmt.Printf("getAlbums from=%d to=%d\n", startYear, endYear)
}
