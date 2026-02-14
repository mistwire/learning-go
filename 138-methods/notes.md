## Go Methods

### What is a method?

A method is a function with a **receiver** — a special parameter that binds the function to a specific type. Instead of calling `speak(p1)`, you call `p1.speak()`. That's the only real difference between a function and a method.

```go
// function
func speak(p person) {
    fmt.Println("Hello, my name is", p.first)
}

// method — same thing, but the receiver moves before the name
func (p person) speak() {
    fmt.Println("Hello, my name is", p.first)
}
```

### Method syntax

The full method signature follows the same pattern as a regular function (Lesson 134), with a receiver added before the name:

```
func (r ReceiverType) identifier(params) (returns) { code }
```

| Part | Example | Purpose |
|------|---------|---------|
| `func` | `func` | keyword |
| receiver | `(p person)` | binds this method to the `person` type |
| identifier | `speak` | method name |
| params | `()` | parameters (same as any func) |
| returns | | return types (same as any func) |

### Value receivers vs pointer receivers

There are two kinds of receivers, and choosing between them matters:

**Value receiver** — gets a **copy** of the value (pass by value, Lesson 134):

```go
func (p person) speak() {
    p.first = "Modified"    // changes the copy only
    fmt.Println(p.first)    // "Modified"
}

p1 := person{first: "James"}
p1.speak()
fmt.Println(p1.first)       // still "James" — original unchanged
```

**Pointer receiver** — gets a **pointer** to the original value:

```go
func (p *person) changeName(name string) {
    p.first = name          // modifies the original
}

p1 := person{first: "James"}
p1.changeName("Bond")
fmt.Println(p1.first)       // "Bond" — original was changed
```

### When to use each receiver type

| Use a value receiver when... | Use a pointer receiver when... |
|------------------------------|-------------------------------|
| The method only reads data | The method needs to **modify** the receiver |
| The type is small (a few fields) | The type is **large** and copying is expensive |
| You want the type to be immutable in this context | You need **consistent** behavior (see below) |

**Consistency rule:** if *any* method on a type uses a pointer receiver, *all* methods on that type should use pointer receivers. Mixing creates confusion about whether a value or pointer is needed to satisfy an interface.

### Methods on any named type

Methods aren't limited to structs — you can attach them to **any named type**:

```go
type myFloat float64

func (m myFloat) abs() myFloat {
    if m < 0 {
        return -m
    }
    return m
}

f := myFloat(-3.14)
fmt.Println(f.abs())     // 3.14
```

The only restriction: you can only define methods on types declared in the **same package**. You can't add methods to `int` or `string` directly — you need to create your own named type first.

### Method promotion through embedding

When a struct embeds another struct (Lesson 126), the inner struct's methods are **promoted** — they become accessible on the outer struct:

```go
type person struct {
    first string
}

func (p person) speak() {
    fmt.Println("I am", p.first)
}

type secretAgent struct {
    person          // embedded — promotes person's methods
    ltk    bool
}

sa := secretAgent{person: person{first: "James"}, ltk: true}
sa.speak()          // works! promoted from person
```

If the outer struct defines its own method with the same name, it **shadows** the promoted one — this is how Go achieves something similar to method overriding (see Lesson 139 for this in action with interfaces).

### Comparing to Python

| Feature | Python | Go |
|---------|--------|----|
| Syntax | `def speak(self):` | `func (p person) speak()` |
| Receiver | implicit `self` (always a reference) | explicit receiver (value or pointer — you choose) |
| Classes | `class Person:` | no classes — `type person struct{}` + methods |
| Inheritance | `class Agent(Person):` | no inheritance — use embedding for composition |
| Calling | `p1.speak()` | `p1.speak()` (same!) |

The biggest difference: Python's `self` is always a reference to the original object. In Go, a value receiver gives you a copy. You must explicitly choose a pointer receiver if you want to modify the original.

### Method sets and interfaces (preview)

Every type has a **method set** — the collection of methods attached to it. This becomes critical with interfaces (Lesson 139):

- A **value** of type `T` can call methods with value receivers
- A **pointer** of type `*T` can call methods with *both* value and pointer receivers

This means pointer receivers affect whether a type satisfies an interface — more on this in Lesson 139.

### Receiver naming conventions

Go uses **short receiver names** — typically one or two letters based on the type name:

```go
func (p person) speak()         // p for person
func (sa secretAgent) speak()   // sa for secretAgent
func (s server) handleRequest() // s for server
func (db database) query()      // db for database
```

Do **not** use `self` or `this` — that's idiomatic in Python and Java but not in Go.

### Key takeaways

- A **method** is a function with a receiver — it attaches behavior to a type
- **Value receivers** get a copy; **pointer receivers** get a reference to the original
- Use pointer receivers when you need to modify the receiver or the struct is large
- You can define methods on **any named type**, not just structs
- Embedded structs **promote** their methods to the outer type (composition over inheritance)
- Go has **no classes** — types + methods + composition replace class-based OOP
- Use **short receiver names** (`p`, `sa`, `s`) — never `self` or `this`
- Method sets determine which **interfaces** a type satisfies (Lesson 139)
