servBinPath = ./bin/server
gameBinPath = ./bin/game
clientBinPath = ./bin/cl

all: serv

game: build_game run_game
serv: build_serv run_serv
cl: build_client run_client

build_serv:
	go build -o $(servBinPath) ./cmd/server/main.go

run_serv:
	./$(servBinPath)

build_game:
	go build -o $(gameBinPath) ./cmd/game/main.go

run_game:
	./$(gameBinPath)

build_client:
	go build -o $(clientBinPath) ./cmd/game/main.go

run_client:
	./$(clientBinPath)

create_default_save:
	go build -o default_save ./cmd/helpers/create_default_save.go && \
		mkdir saves && \
			./default_save --path-to-file="./saves/default.111" && \
				rm -f ./default_save