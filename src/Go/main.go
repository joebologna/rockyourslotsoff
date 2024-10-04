package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"slots/api"
	slots_rpc "slots/api/gen"
	"slots/vslot"

	"github.com/davidrenne/frodo/rpc"
	"github.com/rs/cors"
)

func main() {
	var (
		buf []byte
		err error
	)

	fmt.Println("Starting server on localhost:8998")

	handler := http.NewServeMux()
	service := api.VSlotServiceHandler{MyVSlot: vslot.NewMyVSlot(vslot.WinningSeed, vslot.LoseAmount)}

	gateway := slots_rpc.NewVSlotServiceGateway(&service, rpc.WithMiddleware(cors.AllowAll().ServeHTTP))
	handler.HandleFunc("/", gateway.ServeHTTP)
	handler.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		if buf, err = os.ReadFile("index.html"); err != nil {
			panic(err)
		}
		w.Write(buf)
	})
	js := "api/gen/vslot_service.gen.client.js"
	handler.HandleFunc("/"+js, func(w http.ResponseWriter, r *http.Request) {
		if buf, err = os.ReadFile(js); err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(buf)
	})

	if err = http.ListenAndServe("localhost:8998", handler); err != nil {
		log.Fatal(err)
	}
}
