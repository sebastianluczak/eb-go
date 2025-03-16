# Introduction to Go Programming

**Author: Sebastian Łuczak**

**Date: March 2025**
---

## FYI

* All things considered, this is solely my point of view
* I've interacted with Go for less than a day, take that into account
* I've quickly gathered all necessary resources to learn Go in just a few hours

---

## What is it about Go?

* The TypeScript compiler was rewritten in Go, which piqued my curiosity
* Go is often recommended as a go-to language for former PHP developers
* Go can build native binaries with cross-compilation (*rules apply as I discovered afterwards)

---

## Package/Module system in Go

* Go uses a module system to manage dependencies, similar to npm in JavaScript/TypeScript
* Modules are collections of related Go packages
* The `go.mod` file defines the module's properties and dependencies, akin to `package.json`

---

## Building with cross-compilation

* Go makes it easy to build binaries for different operating systems and architectures
* Use the `GOOS` and `GOARCH` environment variables to specify the target OS and architecture

---

## Building with cross-compilation

> Yeah, with a catch 

If your app uses cgo (C language bindings), cross-compilation can be more challenging

* You may need a cross-compiler for the target platform, with additional libraries on your machine
* See: [Reddit Pitfalls](https://www.reddit.com/r/golang/comments/2ut9hw/real_pitfalls_of_crosscompiling_with_cgodisabled/)
---

## Making an API in Go

* Go's standard library includes powerful packages for building web servers and APIs, similar to Express in Node.js
* The `net/http` package provides HTTP client and server implementations
* Use the `http.HandleFunc` function to define endpoints and their handlers

---

## Complete application for REST API:

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```

---

## 15 LINES OF CODE (!):

```go
1.   package main
2.   
3.   import (
4.       "fmt"
5.       "net/http"
6.   )
7.  
8.   func helloHandler(w http.ResponseWriter, r *http.Request) {
9        fmt.Fprintf(w, "Hello, World!")
10.  }
11. 
12.  func main() {
13.      http.HandleFunc("/hello", helloHandler)
14.      http.ListenAndServe(":8080", nil)
15.  }
```

---

## Explanation

* This code sets up a simple web server that responds with "Hello, World!" at the `/hello` endpoint, similar to setting up a route in Express, but with much less verbosity!
* Go emphasizes simplicity and efficiency, which might require a shift in mindset for developers accustomed to PHP's more verbose and flexible nature.

---

## Migrating to Go

* **No Need for Separate Web Server**: Unlike TypeScript (Express) or PHP (Apache/Nginx), Go can handle HTTP requests directly.
* **Powerful Concurrency**: Go's goroutines are simpler and more powerful than JavaScript's async/await.
* **Explicit Error Handling**: Go requires explicit error handling, unlike JavaScript's try/catch.


---

## PHP compared

* **Compile-time Type Safety**
  * Go: ✅ Yes (static typing)
  * PHP: ❌ No (dynamic typing)
* **Prevents Type Errors Before Running**
  * Go: ✅ Yes
  * PHP: ❌ No
* **Explicit Error Handling**
  * Go: ✅ Yes
  * PHP: ❌ No


---

## PHP compared vol.2

* **Memory Safety**
  * Go: ✅ Yes
  * PHP: ❌ No
* **Crashes at Compile-Time Instead of Runtime**
  * Go: ✅ Yes
  * PHP: ❌ No
* **Concurrency Resilience**
  * Go: ✅ Go’s worker model is safer
  * PHP: ❌ PHP requests can crash independently

---

## **Winner? 🚀 GO.**

- **PHP can be made safer** using strict types (`declare(strict_types=1);`) and exception handling, but by default, **Go is inherently safer**.
- In Go, **bugs are caught at compile-time** rather than at runtime.
- Go enforces explicit error handling, while PHP **allows code that can break unpredictably**.
- **Go’s safety comes with verbosity**, but **the trade-off is beneficial** in production environments.

--- 

## Thank You!

* Questions?
* Feedback?
* Connect with me on [LinkedIn](https://www.linkedin.com/in/sebastianmluczak) or [GitHub](https://github.com/sebastianluczak)

---

**Happy Coding!**