// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	commandCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "daring_cyclops_command_total",
		Help: "total count of commands since boot",
	})

	commandPopulation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "daring_cyclops_active_commands",
		Help: "current population of active commands",
	})

	gameCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "daring_cyclops_game_total",
		Help: "total count of games since boot",
	})

	gamePopulation = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "daring_cyclops_active_games",
		Help: "current population of active games",
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
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
