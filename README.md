# Learning Go

Notes from [Todd McLeod's Udemy course](https://www.udemy.com/course/learn-how-to-code/) | [Course repo](https://github.com/GoesToEleven/learn-to-code-go-version-03)

Each lesson file includes inline annotations and a `REMARKS` section at the bottom with key takeaways, cross-references to related lessons, and common patterns.

## Topics Covered

| # | Topic | Key Concepts |
|---|-------|--------------|
| 004 | User Input | `fmt.Scanln`, pointers `&`, error handling `if err != nil` |
| 013-014 | Variables | `:=`, `var`, blank identifier `_`, zero values, multiple assignment |
| 017-018 | Types | type conversion `Type(value)`, strong typing, `%v` `%T` format verbs |
| 024-026 | Constants | bit shifting `<<`, `iota`, custom types (`ByteSize`) |
| 027-030 | Exercises | variable declaration patterns, print verbs `%d %b %#x`, numeric limits |
| 033 | Scope | package/block scope, variable shadowing, exported vs unexported, methods preview |
| 049-054 | Modules | external packages, `go.mod`, transitive deps, git tags, version pinning |
| 066-070 | Control Flow | `init()`, `if`/`else`, logical operators `&& \|\| !`, comma-ok idiom, `switch`, `select` |
| 071 | Concurrency | goroutines, channels, `select`, concurrency vs parallelism |
| 072-075 | Loops | `for` (3 forms), `break`, `continue`, nested loops, `range`, modulus `%` |
| 076-092 | Exercises | control flow, loops, slices, maps, comma-ok idiom |
| 096-097 | Arrays | fixed-size, `[...]` literals, `len()`, value types |
| 098-105 | Slices | dynamic sizing, `append`, slicing `[i:j]`, `make([]T, len, cap)`, delete, 2D slices |
| 106-108 | Slice Internals | shared underlying arrays, `copy()`, pass-by-value with pointers, side effects |
| 109-115 | Exercises | arrays, slices, composite literals, multidimensional slices |
| 116-119 | Maps | `map[K]V`, `make`, `for-range`, `delete()`, comma-ok idiom `v, ok := m[key]` |
| 121-124 | Exercises | map operations, word counting |
| 125-127 | Structs | named types, embedded structs (composition), anonymous structs, field promotion, shadowing |
| 129-132 | Struct Exercises | structs with slice fields, `map[string]struct`, embedded structs, anonymous structs with composite fields |
| 133-134 | Function Syntax | `func`, parameters vs arguments, returns, multiple return values, named returns, pass by value |
| 135-136 | Variadic & Unfurling | `...Type`, variadic as slice, unfurling slices `slice...`, `fmt.Println` is variadic |
| 137 | Defer | `defer`, LIFO stack, deferred cleanup, resource management |
| 138 | Methods | receiver `(r Type)`, value receiver, dot notation, methods on named types |
| 139 | Interfaces & Polymorphism | `interface`, implicit satisfaction, method sets, polymorphic functions |
| 140 | Stringer Interface | `fmt.Stringer`, `String() string`, `strconv.Itoa`, named types with methods |

## Running Examples

```bash
go run <folder>/main.go
```
