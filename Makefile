CONKY=/etc/conky
POLYBAR=~/.config/polybar
APP=cryptomon

all: build
	sudo chmod a+wrx ${APP}
	sudo cp ./${APP} ${CONKY} 
	sudo cp ./${APP} ${POLYBAR}

build: 
	go build -ldflags="-w -s" -o ${APP}
