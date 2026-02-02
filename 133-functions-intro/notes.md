## Go Functions

### What is a function?

A function is a reusable block of code that performs a specific task. We've already been using one — `func main()` is the entry point of every Go program.

### Basic syntax

```go
func functionName(param1 type, param2 type) returnType {
    // function body
    return value
}
```

- `func` — keyword to declare a function
- `functionName` — follows the same naming rules as variables (camelCase, exported if Uppercase)
- parameters go inside `()` — each must have a name and type
- return type comes after the `()` — omit it if the function returns nothing

### Parameters and arguments

- **Parameters** — the variables listed in the function definition
- **Arguments** — the actual values passed when calling the function

```go
func greet(name string) {       // "name" is a parameter
    fmt.Println("Hello", name)
}

greet("Chris")                  // "Chris" is an argument
```

If multiple parameters share the same type, you can shorten:
```go
func add(x, y int) int {       // x and y are both int
    return x + y
}
```

### Return values

Go functions can return zero, one, or multiple values.

**No return value:**
```go
func sayHello() {
    fmt.Println("Hello")
}
```

**Single return value:**
```go
func double(x int) int {
    return x * 2
}
```

**Multiple return values:**
```go
func swap(a, b string) (string, string) {
    return b, a
}

x, y := swap("hello", "world")  // x = "world", y = "hello"
```

This is one of Go's distinctive features — most languages only allow one return value. Multiple returns are idiomatic in Go, especially for returning a result and an error.

### Named return values

You can name your return values. They act as variables declared at the top of the function:

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("cannot divide by zero")
        return  // "naked return" — returns the named values
    }
    result = a / b
    return
}
```

Named returns can help document what a function returns, but naked returns can hurt readability in longer functions. Use them sparingly.

### Everything in Go is passed by value

When you pass an argument to a function, Go **copies** the value. The function works on the copy, not the original:

```go
func increment(x int) {
    x++                    // modifies the local copy only
}

num := 10
increment(num)
fmt.Println(num)           // still 10 — the original was not changed
```

This applies to all types: ints, strings, structs, arrays. However, slices and maps contain internal pointers, so while the slice/map header is copied, changes to the underlying data are visible to the caller (covered in earlier lessons on reference types).

### Variadic functions

A function can accept a variable number of arguments using `...` before the type:

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)       // 6
sum(1, 2, 3, 4, 5) // 15
```

- The variadic parameter must be the **last** parameter
- Inside the function, `nums` is a `[]int` (a slice)
- `fmt.Println` is a variadic function — that's why it accepts any number of arguments

### The func keyword creates a type

In Go, functions are a **type** — they are first-class values. You can:
- Assign a function to a variable
- Pass a function as an argument to another function
- Return a function from a function

```go
// function assigned to a variable
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4))  // 7
```

This will be explored more in upcoming lessons (anonymous functions, closures, callbacks).

### Key takeaways

- Functions are declared with `func`, followed by name, parameters, return type(s), and body
- Parameters must be typed — Go does not infer parameter types
- Multiple return values are a core Go pattern (especially `result, err`)
- Go is **pass by value** — arguments are always copied
- Variadic parameters (`...type`) accept zero or more arguments as a slice
- Functions are first-class values — they can be stored in variables and passed around
- Exported functions start with an uppercase letter (e.g., `fmt.Println`)
