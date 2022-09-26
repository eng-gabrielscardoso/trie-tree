# Trie tree in Golang

This repository is a collection of codes for implementing a Trie tree using the Go language for the purposes of the discipline Data Structures II - UFPA.

## Running codes

### Running with SDK

1. You must have the specific version of the Golang specified in the `go.mod` file or above.
2. You must run the codes manually (terminal), do not use VS Code extensio or shortcuts, just dirty your hand.
```shell
$ go run . 
or
$ go run *.go
```
3. You also could build the code and generate the binary that will run in your specific architecture and OS
```shell
$ go build
```

### Running with Docker

> Solution still under development

1. You must have Docker installed in your machine first than all.
2. You have to run the following command in the root directory of the application:
```shell
$ sudo docker build --tag trie-tree .
```
3. Run the container with the following command:
```shell
$ sudo docker run trie-tree
```

## Licence

This project is licensed under the [MIT licence](LICENCE.md).
