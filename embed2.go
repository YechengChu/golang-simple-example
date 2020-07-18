package main

import (
  "fmt"
)

type Log struct {
  msg string
}

type Customer struct {
  Name string
  Log
}

func main() {
  c := &Customer{"Sean", Log{"1 - Hello everyone!"}}
  fmt.Println(c)
  fmt.Println("----------------------------------")
  c.Add("2 - Welcome to Computation class!")
  fmt.Println(c)
  fmt.Println("----------------------------------")
  d := Customer{"Barry", Log{"1 - Hi!"}}
  d.Add("2 - Welcome to Mobile Systems class!")
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
  return c.Name + "\nLog:\n" + (&(c.Log)).String()
} // String
