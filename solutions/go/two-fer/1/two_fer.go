// Package twofer provides a function for sharing cookies.
package twofer

import "fmt"

// ShareWith returns a phrase indicating who gets the extra cookie.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}