// go mod init github.com/guycole/daring-cyclops/worker

package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Banner splash message
const Banner = "Daring Cyclops Worker V0.0"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	log.Println(Banner)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	/*
		worker := game.NewWorker("gameId")
		log.Println(worker)

		message1 := `{"command":["newPlayer", "player1uuid", "CaptainRank", "BlueTeam"]}`

		var result map[string]interface{}
		json.Unmarshal([]byte(message1), &result)

		game.DispatchCommand(message1, *worker)
	*/

	/*
		player1 := game.NewPlayer("player1", game.Player1, game.CaptainRank, game.BlueTeam)
		log.Println(player1)
		game.PlayerAdd(player1, game1)
		playerTest := game.PlayerFind(game.Player1, game1)
		log.Println(playerTest)
	*/

	/*
		ship1 := game.NewShip("shipName", game.Player1, game.FighterShip, game.BlueTeam)
		game.ShipAdd(ship1, game1)

		test := `{"command":["m", "one", "two", "three"]}`
		log.Println(test)

		var result map[string]interface{}
		json.Unmarshal([]byte(test), &result)
		log.Println(result)
		log.Println(result["command"])

		//zzz := game.NewJsonCommand(test, player1)

		rawCommand := game.NewRawCommand(game.Player1, test)
		log.Println(rawCommand)
	*/

	//	demoCommand := game.NewTextCommand("gate", "bogus")
	//	log.Println(demoCommand)

	//	game.DispatchCommand(demoCommand, *demoGame)

	//game.CommandGa("shipId", *demoGame)
}
