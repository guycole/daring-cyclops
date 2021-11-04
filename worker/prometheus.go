// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	durationCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "daring_cyclops_game_duration",
		Help: "duration of game",
	})

	commandCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "daring_cyclops_command_total",
		Help: "total count of commands since boot",
	})

	commandPopulation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "daring_cyclops_active_commands",
		Help: "current population of active commands",
	})

	playerCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "daring_cyclops_player_total",
		Help: "total count of players since boot",
	})

	playerPopulation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "daring_cyclops_active_players",
		Help: "population of active players",
	})
)

func main() {
	log.Println("start start")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
