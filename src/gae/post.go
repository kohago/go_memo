package main

import "time"

type Post struct {
	Author  string
	UserID string
	Message string
	Posted  time.Time
}
