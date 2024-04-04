.DEFAULT_GOAL = debug 

debug:
	g++ -c src/*.cpp -std=c++17 -m64 -g -Wall -I include && \
	 g++ *.o -o bin/debug/main -lSDL2main -lSDL2 -lSDL2_image -lSDL2_ttf -lboost_system  && \
	  ./bin/debug/main

release:
	g++ -c src/*.cpp -std=c++17 -m64 -O3 -w -Wall -I include && g++ *.o -o bin/release/main -s -lSDL2main -lSDL2 -lSDL2_image -lSDL2_ttf && ./bin/release/main
