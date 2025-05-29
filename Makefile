# TARGET ?= wasm
TARGET = target.json

cart.wasm: main.go
	tinygo build -target $(TARGET) -o cart.wasm .

run: cart.wasm
	tic80 --fs . --cmd 'load cart.wasmp & import binary cart.wasm & save & run'

clean:
	rm -f cart.wasm
