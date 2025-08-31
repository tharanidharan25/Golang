# Interfaces in Go

### How a type "implements" an interface?

Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.

```
type shape interface {
  area() float64
}
```

If a type in your code implements an area method, with the same signature (e.g. accepts nothing and returns a float64), then that object is said to implement the shape interface.

```
type circle struct{
	radius float64
}

func (c circle) area() float64 {
  return 3.14 * c.radius * c.radius
}
```

This is different from most other languages, where you have to explicitly assign an interface type to an object, like with Java:
```
class Circle implements Shape
```

### Interfaces Are Implemented Implicitly
A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation. You may add methods to a type and in the process be unknowingly implementing various interfaces, and that's okay.

### Multiple Interfaces
A type can implement any number of interfaces in Go. For example, the empty interface, `interface{}`, is always implemented by every type because it has no requirements.