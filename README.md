# Reverse DNS lookup

## Introduction
A simple implementation of goroutines to handle the reverse DNS lookup of a network.
It receives a CIDR network range and it searches for all the domains which can get.

## Usage
### Direct execution
```sh
go run functions.go main.go -cidr <CIDR range> [-t <number of threads>]
```


### Compiling
```sh
go build -o binary/reverse-dns-lookup functions.go main.go
```

And can be **executed** as:
```sh
./binary/reverse-dns-lookup -cidr <CIDR range> [-t <number of threads>]
```

## TODO
- Improve handling of input parameters (validate, force, etc)
- Store results in output file
- Define the format of the output

## Note for haters
- Yup, go doesn't talks about "threads". For you, when it says ~~threads~~, it means paralell goroutines. 