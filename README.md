# Gotris
tetris implementation in go.
## ⚠️WIP Disclaimer:
This is work in progress in very early stage.

## Try it :

```
go get alfonsodev/gotris
cd $GOPATH/src/github.com/alfonsodev/gotris
go build
./gotris
```

## Current state

The current state just shows the Tetrominos.
- Use enter key to start.
- Up arraw switch between different types of tetrominos.
- Down arrow to rotate them.

## Improvements:
Input should include the package read the key event as a primitive and emit input.Key type menu, then comunicate via channels to the other components


![Gotris demo](../master/docs/gotris.gif)
