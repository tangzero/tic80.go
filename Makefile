NAME = example
WASM = out/$(NAME).wasm
TARGET = target.json
IMPORT_WASM = load cart.wasmp & import binary $(WASM) & save

$(WASM): main.go
	tinygo build -target $(TARGET) -o $@ .

run: $(WASM)
	tic80 --fs . --cmd '$(IMPORT_WASM) & run'

export-mac: $(WASM)
	mkdir -p out/mac
	tic80 --fs . --cli --cmd '$(IMPORT_WASM) & export mac out/mac/$(NAME) & exit'

export-linux: $(WASM)
	mkdir -p out/linux
	tic80 --fs . --cli --cmd '$(IMPORT_WASM) & export linux out/linux/$(NAME) & exit'

export-win: $(WASM)
	mkdir -p out/win
	tic80 --fs . --cli --cmd '$(IMPORT_WASM) & export win out/win/$(NAME) & exit'

clean:
	rm -rf .local $(WASM) $(NAME)

