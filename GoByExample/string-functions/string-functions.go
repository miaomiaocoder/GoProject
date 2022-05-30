package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count:	 ", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:    ", s.Index("teste", "e"))
	p("Join:     ", s.Join([]string{"a", "b", "c"}, "任"))
	p("Repeat:   ", s.Repeat("a", 5))
	p("Replace   ", s.Replace("fooooo", "o", "0", -1))
	p("Replace:  ", s.Replace("fooooo", "o", "0", 3))
	p("Split:    ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
