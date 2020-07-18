package main

import (
  "fmt"
)

type Log struct {
  msg string
}

type Customer struct {
  Name string
  log  *Log
}

func main() {
  c := new(Customer)
  // c.Name = "Sean"
  // c.log = new(Log)
  // c.log.msg = "1 - Hello everyone!"
  // The above three lines have same effect as the line below
  c = &Customer{"Sean", &Log{"1 - Hello everyone!"}}
  fmt.Println(c)
  fmt.Println("----------------------------------")
  c.log.Add("2 - Welcome to computation class!")
  fmt.Println(c)
  fmt.Println("----------------------------------")
  d := Customer{"Barry", &Log{"1 - Hi!"}}
  d.log.Add("2 - Welcome to Mobile Systems class!")
  fmt.Println(d)
  fmt.Printf("The type of d is %T\n", d)
} // main

func (l *Log) Add(s string) {
	l.msg += "\n" + s
} // Add

func (l *Log) String() string {
	return l.msg
} // String

func (c *Customer) String() string {
	return c.Name + "\nLog:\n" + (c.log).String()
} // String
