CC=make
subfolders=hello nrequest
modules=$(addsuffix .wasm, $(subfolders))

all: $(modules)

%.wasm: %
	$(MAKE) -C $<