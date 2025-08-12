package main

import (
	"os"
)

const default_user = "default_user"

func currentUsername() string {
	if u := os.Getenv("USER"); u != "" { // Unix-like
		return u
	}
	if u := os.Getenv("USERNAME"); u != "" { // Windows
		return u
	}
	return default_user
}