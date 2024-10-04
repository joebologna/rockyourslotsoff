all: build
	echo $@

frodo:
	go generate ./...
	awk '/use strict/{next} /module.exports/{exit}{print}' api/gen/vslot_service.gen.client.js > api/gen/vslot_service.gen.client.js.tmp; mv api/gen/vslot_service.gen.client.js.tmp api/gen/vslot_service.gen.client.js

build: frodo
	go build .

test: build
	./slots &
	bats testapi.bats
	-pkill slots

clean:
	rm -f slots
	rm -fr api/gen
