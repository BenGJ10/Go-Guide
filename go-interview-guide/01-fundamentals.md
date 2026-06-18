# Go Fundamentals

## 1. Why `package main`?

### Short Interview Answer

> `package main` tells the Go compiler that this package is an executable program rather than a reusable library package. The Go runtime looks for a `main()` function inside `package main` and starts execution from there.

---

### Detailed Explanation

Go programs are organized into packages.

Example:

```go
package mathutils
```

This is a library package.

```go
package main
```

This is a special package recognized by the Go compiler.

When Go sees:

```go
package main

func main() {
    fmt.Println("Hello")
}
```

it generates an executable binary.

Without `package main`, Go assumes you're creating a reusable package.

---

## Follow-up Interview Question

### What happens if there is no main function?

```go
package main
```

but

```go
func test() {}
```

Output:

```bash
function main is undeclared in package main
```

Because Go needs an entry point.

---

### Internal Working

Execution flow:

```text
OS
 ↓
Generated Binary
 ↓
Go Runtime Initialization
 ↓
main.main()
```

Before your `main()` executes:

* Go Runtime initializes
* Scheduler starts
* Garbage Collector initializes
* Memory allocator initializes

Then:

```go
main()
```

runs.

---

## 2. Difference Between `go run` and `go build`

One of the most common Go interview questions.

---

### Short Interview Answer

> `go run` compiles the source code and immediately executes the generated binary. The binary is temporary and removed after execution.
>
> `go build` only compiles the code and generates a reusable executable binary without running it.

---

## Example

### go run

```bash
go run main.go
```

Internally:

```text
Source Code
    ↓
Compile
    ↓
Temporary Binary
    ↓
Execute
    ↓
Delete Binary
```

---

### go build

```bash
go build
```

Internally:

```text
Source Code
    ↓
Compile
    ↓
Permanent Binary
```

Output:

```bash
./myapp
```

Run later:

```bash
./myapp
```

---

## Real Usage

Development:

```bash
go run .
```

Production:

```bash
go build
```

then

```bash
./app
```

---

## Follow-up

### Which is faster?

For repeated executions:

```bash
go build
```

because compilation happens only once.

---

## 3. How Does Go Work Behind the Scenes?

Interviewers ask this to understand whether you know what happens after writing code.

---

### High-Level Flow

```text
Go Source Code
       ↓
Compiler
       ↓
Machine Code
       ↓
Executable Binary
       ↓
Operating System
       ↓
CPU Executes Instructions
```

---

### Step 1: Write Go Code

```go
fmt.Println("Hello")
```

Human-readable code.

---

### Step 2: Compilation

Go is a compiled language.

```bash
go build
```

The Go compiler:

* Checks syntax
* Checks types
* Resolves imports
* Optimizes code
* Generates machine code

---

### Step 3: Linking

All required packages are linked.

Example:

```go
fmt
net/http
database/sql
```

are combined into one executable.

---

### Step 4: Binary Creation

Produces:

```bash
app.exe
```

or

```bash
app
```

This binary contains:

* Your code
* Go runtime
* Scheduler
* Garbage collector

---

### Step 5: Execution

OS loads binary into memory.

Runtime initializes:

```text
Heap
Stack
Scheduler
GC
```

Then:

```go
main()
```

starts.

---

## 4. How Memory is Managed in Go?

Extremely common interview question.

---

### Short Interview Answer

> Go manages memory automatically using a Garbage Collector. Memory can be allocated on the stack or heap. The compiler decides where variables live through escape analysis.

This answer alone impresses many interviewers.

---

## Stack vs Heap

### Stack

Fast memory.

Example:

```go
func main() {
    x := 10
}
```

`x` usually goes to stack.

Characteristics:

* Very fast
* Automatically cleaned
* Function scoped

---

### Heap

Used when data survives beyond function scope.

Example:

```go
func create() *int {
    x := 10
    return &x
}
```

`x` cannot remain on stack because function ends.

Compiler moves it to heap.

---

## Escape Analysis

One of Go's biggest interview topics.

---

### Question

How does Go decide stack vs heap?

Answer:

> Go uses escape analysis during compilation. If a variable escapes the function scope, it is allocated on the heap; otherwise it remains on the stack.

Example:

```go
func main() {
    x := 10
}
```

Stack.

---

Example:

```go
func create() *int {
    x := 10
    return &x
}
```

Heap.

---

You can check:

```bash
go build -gcflags="-m"
```

Output:

```text
moved to heap: x
```

---

## Garbage Collection

Go automatically frees unused heap memory.

Example:

```go
func main() {
    data := make([]int, 1000000)

    data = nil
}
```

Once nothing references that memory:

```text
GC identifies it
GC frees it
```

---

### Why GC?

Without GC:

```c
malloc()
free()
```

Developer manually manages memory.

Risk:

* Memory leaks
* Dangling pointers
* Double free

Go removes these issues.

---

## Interview Question

### What is Escape Analysis?

Best Answer:

> Escape analysis is a compiler optimization used to determine whether a variable can safely stay on the stack or must be allocated on the heap. Variables that escape the current function scope are moved to the heap, where they are managed by the garbage collector.

---

## One-Liner Revision Sheet

### Why package main?

> Special package that creates an executable program. Execution starts from `main()`.

### Why main()?

> Entry point of a Go application.

### go run

> Compile + Execute + Delete temporary binary.

### go build

> Compile and create reusable executable.

### How Go works?

> Source Code → Compiler → Machine Code → Binary → Runtime → main()

### Memory Management

> Go uses stack, heap, escape analysis, and garbage collection.

### Escape Analysis

> Determines whether a variable lives on stack or heap.

### Stack

> Fast, function-local, auto-cleaned.

### Heap

> Long-lived objects, managed by GC.

### Garbage Collector

> Concurrent mark-and-sweep collector that automatically reclaims unused heap memory.

---