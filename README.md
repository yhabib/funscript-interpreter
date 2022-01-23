# Funscript Interpreter

A toy interpreter for a toy language.

**Note**: This repository is heavily inspired on my learnings from this amazing [book](https://interpreterbook.com/)

## Funscript

Funscript is a new&awesome language resulting of weirdly mergin JavaScript and Kotlin.

The syntax:

```js
var two = 2;
var add = fun(x, y) {
  return x + y;
}
var result = add(two, 3);
```

## Interpreter

An interpreter is a program that takes other programs as their input and produce something meaningful.
There are different types of interpreters from very simple ones to very complex ones:

- Parse the source code(input) and perform its beahaviour directly. Eg: Lisp early versions
- Translate source code into some kind of intermediate representation(bytecode) and then immediately excute this. Eg: Pyhton, Ruby, ...
- JIT interpreters that: compile the input just-in-time into native machine code that gets then executed. Eg: JavaScript

### Lexing

From Source Code to an Abstract Syntax Tree

Lexing is the first step for a program to be interpreted.

The lexer, takes an input, a string, and transforms it into tokens. This process is known as "lexical analysis". Something like `var sum = 2 + 2;` is parsed into a collection of tokens like [VAR, IDENTIFIER(sum), EQUAL_SIGN, NUMBER(2), PLUS_SIGN, NUMBER(2), SEMICOLON], where each entry represents a token type to be interpreted later.

#### Basics

Branch: [lexer/basics](https://github.com/yhabib/funscript-interpreter/tree/lexer/basics)

In this first iteration we defined two modules: Token and Lexer.

A `Token` is defined by a `Type` and a `Literal`. The former will take one of the values that our interpreter will support(defined as constants). While the latter will be the literal value that we are tokenizing.

The `Lexer` will implement all the required to navigate the input, source code, and convert each character of this string into a Token.

In this first iteration we test it against a very simple input with some basic values `=+(){},;`.

#### Advance

Branch: [lexer/advance](https://github.com/yhabib/funscript-interpreter/tree/lexer/advance)

In this second iteration we are enhancing the previously created modules to tokenize this basic program:

```js
var two = 2;
var add = fun(x, y) {
  return x + y;
}
var result = add(two, 3);
```

`Lexer` needs to be able to process the new tokens. Multiletter ones require to be treated a bit specially as we need to read the whole identifier, for that we create two new functions `readNewIdentifier` and `readNumber` that keep reading characters until its done.

We also need to start thinking about white spaces. We don't care about them but we need to consume them `eatWhitespace` does basically this.

Our interpreter is now able to process this simple program. Next, we will extern our token set to fully support our language.

#### Wrapping Up

Branch: [lexer/final](https://github.com/yhabib/funscript-interpreter/tree/lexer/final)

Now we just need to complete our lexer so it is able to understand the rest of funscript.

First we need to add support to the missing basic operators like `-` or `*`. This should be easy to accomplish as we only have to extend our symbols.

We keep adding non meaningful code to our test case as we still don't care about "correctness". The Lexer only cares about tokens.

We add now the missing keywords, `TRUE`, `FALSE`, `IF` and `ELSE`.
