# Pointer Receivers
<b>A receiver type on a method can be a pointer.</b>

- Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
- However, methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

#### Pointer Receiver
```
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
```

#### Non-Pointer Receiver
```
type car struct {
	color string
}

func (c car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "white"
}
```