package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"time"

	"gitlab.com/jeshuamorrissey/mmo_server/auth_server/packet"
	"gitlab.com/jeshuamorrissey/mmo_server/database"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbConnect() *mongo.Database {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error while connecting to database: %v\n", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Could not ping database: %v\n", err)
	}

	return client.Database(database.DatabaseName)
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	// Connect to the database.
	db := dbConnect()
	GenerateTestData(db)

	// Main control loop.
	log.Printf("Listening for connections on :5000...\n")

	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving connection from %v\n", conn.RemoteAddr())

		// Make a new buffer to read from.
		go packet.RunSession(db, bufio.NewReader(conn), bufio.NewWriter(conn))
	}
}
