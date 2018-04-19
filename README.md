# Brainfuck inpterpeter
Brainfuck is an esoteric programming language created in 1993.  This go program compiles and executes  Brainfuck source code.

## Building 
To build the project, [go](https://golang.org/) has to be installed on your system, then run the following commands:

```bash
$ cd src
``` 
```bash
$ go build -o ../bin/bf
```

## Running

### Hello world

```bash
$ ./bf ../resources/hw.bf
Hello World!
```

### First 11 Fibonacci numbers
```bash
$ ./bf ../resources/fibonacci.bf
1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89
```

## Resources
* [Brainfuck programming language](https://en.wikipedia.org/wiki/Brainfuck)
* [Brainfuck example programs](http://www.hevanet.com/cristofd/brainfuck/)
