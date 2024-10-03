package main

import (
	"log"
	"net/http"
	"os"
	"slots/api"
	slots_rpc "slots/api/gen"
	"slots/vslot"

	"github.com/monadicstack/frodo/rpc"
	"github.com/rs/cors"
)

func main() {
	var (
		buf []byte
		err error
	)
	if buf, err = os.ReadFile("index.html"); err != nil {
		panic(err)
	}

	handler := http.NewServeMux()
	service := api.VSlotServiceHandler{MyVSlot: vslot.NewMyVSlot(42)}

	gateway := slots_rpc.NewVSlotServiceGateway(&service, rpc.WithMiddleware(cors.AllowAll().ServeHTTP))
	handler.HandleFunc("/", gateway.ServeHTTP)
	handler.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		w.Write(buf)
	})

	if err = http.ListenAndServe("localhost:8998", handler); err != nil {
		log.Fatal(err)
	}
}
