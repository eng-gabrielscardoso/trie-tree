# Architecture

The application is divided into modules, as like any project in Golang. The directory structure is like below:

```shell
.
├── ARCHITECTURE.md
├── AUTHORS.md
├── core
│   ├── Node.go
│   └── Trie.go
├── go.mod
├── LICENCE.md
├── main.go
└── README.md
```

## Root directory

In root directory you will find documentation files and two important files for application: ```go.mod``` and ```main.go```.

The ```go.mod``` is an important file that contais the Go version and informations about the module, i.e. the dependencies and versions of any.

The ```main.go``` is the core script that contains all the logic to initialise the application. In there, you will find the necessary comands that wrapper the main functionalities of application.

## /core

In the ```/core``` directory, you will find the following modules: ```Node.go``` and ```Trie.go```. Which modules should be used when creating the application, and both represents the abstraction of the concepts of node and trie tree.

Go inside then to get more knowledge about functionalities.
