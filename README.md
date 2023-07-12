# funcreturn
Checks whether there is a newline between functions

[![Build Status](https://travis-ci.org/hsnks100/funcreturn.svg?branch=master)](https://travis-ci.org/ssgreg/nlreturn)
[![Go Report Status](https://goreportcard.com/badge/github.com/hsnks100/funcreturn)](https://goreportcard.com/report/github.com/ssgreg/nlreturn)
[![Coverage Status](https://coveralls.io/repos/github/hsnks100/funcreturn/badge.svg?branch=master&service=github)](https://coveralls.io/github/ssgreg/nlreturn?branch=master)

Linter requires a new line between functions

# Example

Examples of incorrect code:

```go 
func foo() {
} 
func bar() { // want new line
}
```

```go 
func foo() {
}


func bar() { // want only one new line
}
```

Examples of correct code:

```go
func foo() {
}

func bar() { // want new line
}
```
