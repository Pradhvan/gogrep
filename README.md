# mygrep
A command line program written in Go that implements Unix grep like functionality.

## Running the program locally

1. Clone the repo 

2. Run the make file with `make` command and it should give the following output.
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

## Flags supported by mygrep

* **-i**: Make the seach case sensitive.
* **-c**: Count number of matches.
* **-o**: Store the search results in a file.
* **-B**: "Print *'n'* lines before the match"



## Usage
```
$ ./mygrep -flag[optional] searchword filname.txt or directory 
```

* Search *'go'* in *'word.txt'*
```
$ ./mygrep foo word.txt
words.txt: Yet another go program.
words.txt: Learning Go.
words.txt: Real-World Go Programming
words.txt: Jhon Bodners Go Book
```

* *Case sensitive* search for 'go'
```
$ ./mygrep -i go word.txt
words.txt: Yet another go program.
```

* *Count* total search result for 'go'
```
$ ./mygrep -c go word.txt
Total matches found for 'go' in words.txt: 4
```

* Search 'go' in 'words' *directory*
```
$ ./mygrep go words
words/inner/inner-inner/inner-inner1.txt: Real-World Go Programming
words/inner/inner1.txt: Real-World Go Programming
words/outer.txt: Learning Go.
```

* *Store* match result for 'go' in 'output.txt'
```
$ ./mygrep -o output.txt go word.txt
$ cat output.txt
words.txt: Yet another go program.
words.txt: Learning Go.
words.txt: Real-World Go Programming
words.txt: Jhon Bodners Go Book
```

* Print 'n' number of lines before a match
```
$ ./mygrep -B 1 go word.txt
words.txt:this is a text here.
words.txt:Yet another go program.
words.txt:Find me!
words.txt:Learning Go.
words.txt:An Idomatic Approach to
words.txt:Real-World Go Programming
```

* Using multiples flags
```
$ ./mygrep -B 1 -o output.txt -i -c go word.txt
Total matches found for 'go' in word.txt: 2
```
