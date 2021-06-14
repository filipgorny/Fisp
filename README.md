# Custom Lisp language parsing and evaluating made in Go lang.

## *Work in progress. Currently rewriting from Typescript to Go lang.*

*This Readme.md is temporary, will rewrite it later to something more adequate.*

## How the wxl language works

Wxl is a [Lisp](lishttps://en.wikipedia.org/wiki/Lispp) language, than means it's syntax is made of lists and it uses [polish notation](https://en.wikipedia.org/wiki/Polish_notation).
The keyword that is a name of the function is the first element of a list, and the rest are the arguments.

For example expression `(+ 2 3)` means something like:

"To evaluate this function use `add` method. Call it with arguments: 2, 3."

## Supported syntax at the moment

### Math

* adding: `+`
* substracting: `-`
* dividie: `/`
* multiple: `*`

### Setting variables

* set a variable: `set`
Example: `(set a 5)`

### Functions

* declare function: `fun`
For example: `(fun add (a b) (+ a b))` and then to use it: `(add 2 3)`

## About the source code, contribution and adding functionality

/WILL ADD THIS SECTION LATER/