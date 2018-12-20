package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {
	// user.Current needs cgo, this does not
	u, err := user.LookupId(strconv.Itoa(os.Getuid()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not get current user: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("I am %s! UID=%s GID=%s\n", u.Username, u.Uid, u.Gid)
}
