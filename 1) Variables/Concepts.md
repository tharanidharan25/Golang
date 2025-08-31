# Variables Concepts

## Basic Types
- int
- float64
- bool
- string

## Declaration
#### Normal Declaration
```
var username string
```

#### Walrus Operator Declaration - Easiest way to declare variables
```
username := "username"
```

## Type Sizes
Integers, uints, floats and complex all have type sizes.

#### Whole Numbers (No Decimals)
`int   int8   int16   int32   int64`

#### Positive Numbers (No Decimals)
uint stands for "Unsigned Integers"

`uint  uint8  uint16  uint32  uint64`

#### Signed Decimals
`float32 float64`

#### Complex Numbers (a Complex Number Has a Real and Imaginary Part)
`complex64 complex128`

The size represents how many bits in the memory will be used to store the variable

The default 'int' and 'uint' types refer to their 32 and 64 bits depending on the environment of the user.

"Standard" sizes that should be used unless there's a specific performance need are:
- int
- uint
- float64
- complex128

#### Converting between types:
```
temperatureFloat := 88.26
temperatureInt := int64(temperatureFloat)
```

Casting a float to an integer in this way truncates the floating point portion.

## Formatting Strings in GO
#### Default Representation (%v)
```
s := fmt.Sprintf("I am %v years old", 10)

O/P: I am way 10 years old
```

#### String (%s)
```
s := fmt.Sprintf("I am %s years old", "way too many")

O/P: I am way too many years old
```

#### Integer (%d)
```
s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old
```

#### Float (%f)
```
s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old
```