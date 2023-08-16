# Basic Golang Programming

This repository contains basic syntax like variable, datatype, and control flow (conditional and looping).

This is just kind of a note for who want to learn golang from the very basic (it only covers the surface). To explore more deep about golang as programming language, you can read the [Full Documentation](https://go.dev/doc/)

## Introduction

Golang is a Programming Language created by google in November 10, 2009. It gains its popularity. Golang is short of Go and Language.

### **Simplicity and Readability**

Go's syntax is designed to be simple and readable, making it easy for developers to write and understand code.

### **Open Source**

Go is open source, meaning that its source code is available to the public. This fosters a collaborative community of developers who contribute to its development and create third-party libraries and tools.

[Read More...](https://codilime.com/blog/what-is-go-language/#:~:text=A%20short%20history%20of%20the%20Go%20language&text=That%27s%20why%20on%20September%2021,a%20public%20open%20source%20project.)

## Online Compiler

Golang has an online compiler, if you want to try golang without need to setup in your local computer, you can go to [Go Playground](https://go.dev/play/)

## Hello World Program

```
package main
import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

## How to create a Variable?

```
// without datatype in defining
var say = "hello"

// define explicitly
var toName string = "John"

// multiple declare and define
var width, heigth int = 100, 70

// declared without defined
var declared int

// short declaration
shortVar := "this is a short defined variable"
```
