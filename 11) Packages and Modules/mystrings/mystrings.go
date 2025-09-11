// by convention, we name our package the same as the directory
package mystrings

// Reverse reverses a string left to right
// Notice that we need to capitalize the first letter of the function
// If we don't then we won't be able to access this function outside of the
// mystrings package
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

// Run go build . Because this isn't a main package, it won't build an executable. However, go build will still compile the package and save it to our local build cache. It's useful for checking for compile errors.