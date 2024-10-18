CONKY=/etc/conky
POLYBAR=~/.config/polybar
DEST=${CONKY} ${POLYBAR}
APP=cryptomon
ENV=.env

all: build
	sudo chmod a+x ${APP}
	for dest in ${DEST}; do \
		sudo cp ./${APP} ${ENV} $$dest; \
	done

build: 
	go build -ldflags="-w -s" -o ${APP}
