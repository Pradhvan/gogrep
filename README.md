# mygrep
A command line program written in Go that implements Unix grep like functionality.

### Usage

1. Clone the repo 

2. Run the make file and it should give the following output.
```
$ make
go fmt 
go vet
go build -o mygrep .
```

3. Use the binary to run search the file.
```
$ ./mygrep searchterm filetosearchin.txt
```