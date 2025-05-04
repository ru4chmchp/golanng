# Variables

In Go, there are three main ways to declare variables. Each method is suitable for a specific context and is widely used, although their popularity varies depending on the situation.

## ✅ 1. Use `var` with clear data type

```bash
    var name string = "ru4chmchp"
    var age int = 21
    var pi float64 = 60.3
    var isLearning bool = true
```

Use when you want to clear data types, especially in global declarations or when the type is complex.

## ✅ 2. Use `var` with type inference

```bash
    var name = "John"
```

Go will infer type from assigned value (`string` in the above example).

## ✅ 3. Use short declaration `:=` (only used in functions)

```bash
score := 95
```

+ Most common in real code, especially internal function logic.
+ Automatic type inferance, neat.
+ Can't be used outside of functions (`global scope`).