package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/shdlabs/week21/client/web/server"
	h "github.com/shdlabs/week21/helpers"
	"github.com/shdlabs/week21/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Info("starting web client")
	// Set up a connection to the server.
	conn, err := grpc.NewClient(
		"asus:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %#v", err)
	}

	defer conn.Close()

	c := service.NewQueryUserClient(conn)
	d := service.NewQueryAllClient(conn)

	// Contact the server and print out its response.

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		defer h.DurationLog(time.Now(), "/search")
		q := r.URL.Query().Get("user")
		id, err := strconv.Atoi(q)
		if err != nil {
			log.Error("could not parse ID", "err", err)
			id = 0
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rRes, err := c.GetUser(ctx, &service.UserRequest{Id: int32(id)})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = server.Index(server.User(rRes)).Render(ctx, w); err != nil {
			log.Error("could not render template", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		rRes, err := d.GetAll(ctx, &service.All{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := range rRes.Users {
			if err = server.Index(server.User(rRes.Users[i])).Render(ctx, w); err != nil {
				log.Error("could not render template", "err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	log.Info("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
