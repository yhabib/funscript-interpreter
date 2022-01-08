# Funscript Interpreter

I'm writing this interpreter as a learning exercise.

**Note**: This repository is heavily inspired on my learnings from this amazing [book](https://interpreterbook.com/)

## Funscript

Funscript is a new&awesome language resulting of weirdly mergin JavaScript and Kotlin.

The syntax:

```js
var two = 2         // Defines a re-assignable variable

var add = fun(x, y) {
  x + y;
}

var result = add(two, 3)
```

## Interpreter

### Lexing

From Source Code to an Abstract Syntax Tree

* Lexical Analysis: Transforms source code into tokens, eg: `var sum = 2 + 2;` -> [VAR, IDENTIFIER("sum), EQUAL_SIGN, INTEGER(2), PLUS_SIGN, INTEGER(2), SEMICOLON]