# The Go programming language
The idea is to have a note for the programming patterns
with small snippets to quickly revisit and understand

This denotes the package or module structuring. Do not understand much at this point. 
```
package main
```

## Import
Provides the necessary libraries for operations

For e.g., this provides the I/O functions from the looks of it.

```
import "fmt"
```

Here is a [list](https://pkg.go.dev/std) of standard packages in Go 

end of statement need not be explicitly declared.

Q. how does it identify then?
## Operators
```
&& AND
|| OR
```

## Variables

Refer to <strong>[variables.go](./variables.go)</strong> and <strong>[constant.go](./constant.go)</strong> for notes


## Looping

Refer to <strong>[for.go](./for.go)</strong> for examples and patterns

## Conditional statements

Refer to <strong>[if-else.go](./if-else.go)</strong> and <strong>[switch.go](./switch.go)</strong> for examples and patterns

## Arrays

Refer to <strong>[arrays.go](./arrays.go)</strong>

## Slices

Slices are basically dynamic arrays with extended functionalities. 

- make
- copy
- Equal

Refer to <strong>[slices.go](./slices.go)</strong>

Note: Requires library "slices"

## Map

Functions
- make
- clear

Refer to <strong>[map.go](./map.go)</strong>

## Range

Sort of like an iterator on iterable objects

1. Range on arrays and slices provides both the index and value for each entry.

2. Range on map iterates over key/value pairs.

3. Range can also iterate over just the keys of a map.

4. Range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.

Refer to <strong>[range.go](./range.go)</strong> for examples of the discussion above. 

## Functions


