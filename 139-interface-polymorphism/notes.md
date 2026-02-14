## Go Interfaces & Polymorphism

### What is an interface?

An interface is a type that defines a **method set** — a list of method signatures that a type must have in order to "satisfy" the interface. Interfaces describe **behavior**, not data.

```go
type human interface {
    speak()
}
```

This says: "any type that has a `speak()` method with no parameters and no return values is a `human`."

### Implicit satisfaction (no `implements` keyword)

In Go, a type satisfies an interface **automatically** — there is no `implements` keyword like in Java or C#. If the methods match, the type satisfies the interface. Period.

```go
type person struct {
    first string
}

func (p person) speak() {
    fmt.Println("I am", p.first)
}
// person satisfies `human` because it has a speak() method — nothing else needed
```

This is sometimes called **structural typing** (vs. Java's **nominal typing** where you must explicitly declare `implements`). It means you can define an interface *after* the types that satisfy it already exist, which keeps packages decoupled.

### What is polymorphism?

Polymorphism means "many forms" — a single function can work with **different types** as long as they satisfy the same interface. The behavior changes based on the **concrete type** passed in at runtime.

```go
func saySomething(h human) {
    h.speak()   // calls whatever speak() the concrete type has
}

saySomething(person{first: "Jenny"})
// output: I am Jenny

saySomething(secretAgent{person: person{first: "James"}, ltk: true})
// output: I am a secret agent James
```

One function, two different types, two different behaviors — that's polymorphism.

### How interfaces relate to structs and methods

Go's approach to object-oriented design uses three building blocks:

| Concept | Purpose | Lesson |
|---------|---------|--------|
| Structs | Define **data** (fields) | 125-126 |
| Methods | Attach **behavior** to types | 138 |
| Interfaces | Enable **polymorphism** via shared behavior | 139 |

Together, these three replace the class-based OOP found in languages like Python, Java, and C#. Go has **no classes, no inheritance, no `extends`** — only composition and interfaces.

### Interface values are a type + value pair

Under the hood, an interface variable holds two things:

1. The **concrete type** (e.g., `person`)
2. The **concrete value** (e.g., `person{first: "Jenny"}`)

```go
var h human
h = person{first: "Jenny"}
// h internally stores: (type=person, value=person{first:"Jenny"})
```

This is sometimes called a "fat pointer." You can inspect it with `fmt.Printf("%T %v\n", h, h)`.

### The empty interface

The empty interface `interface{}` (or its alias `any` since Go 1.18) has **zero methods**, so every type satisfies it:

```go
func printAnything(v any) {
    fmt.Println(v)
}

printAnything(42)        // int
printAnything("hello")   // string
printAnything(true)      // bool
```

This is why `fmt.Println` can accept arguments of any type — its signature uses the empty interface.

### Embedding and method shadowing with interfaces

When a struct embeds another struct, the inner struct's methods are **promoted** (see Lesson 126). But if the outer struct defines its own method with the same name, it **shadows** the promoted one:

```go
type person struct {
    first string
}

type secretAgent struct {
    person        // embeds person — promotes person.speak()
    ltk    bool
}

func (p person) speak() {
    fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
    fmt.Println("I am a secret agent", sa.first)
}
```

Both `person` and `secretAgent` satisfy the `human` interface, but each with their **own** `speak()` method. The promoted method from `person` is shadowed by `secretAgent`'s version.

### Comparing to Python

| Feature | Python | Go |
|---------|--------|----|
| Polymorphism | Duck typing at runtime — no formal interface needed | Interfaces — compile-time checked, but still implicit |
| Declaration | No `implements` needed (duck typing) | No `implements` needed (structural typing) |
| Safety | Errors at runtime if method missing | Errors at **compile time** if method missing |
| Philosophy | "If it walks like a duck..." | Same idea, but the compiler verifies it for you |

Go gives you the flexibility of duck typing with the safety of compile-time checks — the best of both worlds.

### Interface design guidelines

- **Keep interfaces small** — Go proverb: *"The bigger the interface, the weaker the abstraction"* (Rob Pike). The best interfaces have one or two methods (e.g., `io.Reader`, `io.Writer`, `fmt.Stringer`).
- **Accept interfaces, return structs** — functions should take interface parameters (flexible inputs) but return concrete types (specific outputs).
- **Define interfaces where they're used**, not where the types are defined — this keeps packages decoupled.
- **Don't create interfaces prematurely** — wait until you have two or more types that need to share behavior.

### Common standard library interfaces

| Interface | Method | Package |
|-----------|--------|---------|
| `fmt.Stringer` | `String() string` | `fmt` |
| `io.Reader` | `Read(p []byte) (n int, err error)` | `io` |
| `io.Writer` | `Write(p []byte) (n int, err error)` | `io` |
| `error` | `Error() string` | built-in |
| `sort.Interface` | `Len()`, `Less()`, `Swap()` | `sort` |

Implementing these standard interfaces makes your types work seamlessly with the standard library.

### Key takeaways

- An **interface** defines a method set — any type with those methods satisfies it automatically
- Go uses **implicit** (structural) satisfaction — no `implements` keyword
- **Polymorphism** = one function works with multiple types via a shared interface
- Interface values store a **(type, value)** pair internally
- The **empty interface** (`any` / `interface{}`) is satisfied by every type
- Keep interfaces **small** — one or two methods is ideal
- **Accept interfaces, return structs** — a key Go design pattern
- Structs + methods + interfaces = Go's alternative to class-based OOP
