
# Monkey Language

Based on the Thorsten Ball book - [Writing An Interpreter In Go](https://interpreterbook.com/)

Right now, only REPL (Read, Evaluate, Print, Loop) is supported. To run the program, you can simply use the command
```
go run main.go
```

# Language Features

let keyword is used for variable initialization
```
let foo = "bar";
let baz = 123;
let abc = [1, 2, 3];
```
functions are first-class citizens in this language, and can be used / passed around as variables
```
let greet = fn(name) {
    return "Hello " + name + " !";
}
```

the monkey language also supports closures
```
let newGreeter = fn(greeting) {
  // `puts` is a built-in function we add to the interpreter
  return fn(name) { puts(greeting + " " + name); }
};

// `hello` is a greeter function that says "Hello"
let hello = newGreeter("Hello");

// Calling it outputs the greeting:
hello("John!"); // => Hello John!
```
