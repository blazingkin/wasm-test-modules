CC=make
subfolders=hello

all: $(subfolders)
	$(CC) -C $(subfolders)