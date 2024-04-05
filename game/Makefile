.DEFAULT_GOAL = debug 

debug_name = debug_player_0
debug_pin = 8080

release_name = release_player_0
release_pin = 133

debug: debug_build debug_run

debug_build: # build debug version (no stripped)
	g++ -c src/*.cpp -std=c++17 -m64 -g -w -Wall -I include && \
	 g++ *.o -o bin/debug/main -lSDL2main -lSDL2 -lSDL2_image -lSDL2_ttf -lboost_system

debug_run: # run debug version
	  ./bin/debug/main $(debug_name) $(debug_pin)

release: release_build release_run # build and than run


release_build: # build release version (stripped)
	g++ -c src/*.cpp -std=c++17 -m64 -w -Wall -I include && \
	 g++ *.o -o bin/release/main -lSDL2main -lSDL2 -lSDL2_image -lSDL2_ttf -lboost_system && \
	 strip bin/release/main

release_run: # run release version
	./bin/release/main $(release_name) $(release_pin)


clear: # clear all .o files
	rm -f *.o
