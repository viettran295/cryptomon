CONKY=/etc/conky
APP=cryptomon

all: build
	sudo chmod a+wrx ${APP}
	sudo mv ./${APP} ${CONKY}

build: 
	go build -o ${APP}
