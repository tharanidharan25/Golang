# Memory layout of structs

> In Go, structs sit in memory in a contiguous block, with fields placed one after another as defined in the struct.

> Field ordering matters a lot.

```
type stats struct {
	NumPosts uint8
	Reach    uint16
	NumLikes uint8
}
```
`In the above struct definition, Go will pad the memory usage of NumPosts to 2 bytes instead of one and will pad the memory usage of numLikes to 2 bytes instead of one, since Reach will need 2 bytes. GO does this for execution speed, but this can lead to increased memory usage.`

> if you have a specific reason to be concerned about memory usage, aligning the fields by size (largest to smallest) can help.