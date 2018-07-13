package main

import ( 
  "fmt"
  "log"
  "bufio"
  "os"
  "io"
  "time"
)

func main() {
  f, err := os.OpenFile("songs", os.O_RDONLY, os.ModePerm)
  if err != nil {
    log.Fatalf("open file error: %v", err)
    return
  }
  defer f.Close()

  tags := make(map[string]int)
  reader := bufio.NewReader(f)
  count := 1
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      if err == io.EOF { break }
      log.Fatalf("read line error: %v", err)
    }
    tags[line] = count
    count += 1
  }
  start:= time.Now()
  n := Rebuild(tags)
  t := time.Now()
  fmt.Println("build tags %v", t.Sub(start))

  start = time.Now()
  r:= Find(n, "hello")
  t = time.Now()
  fmt.Println("hello: ", r, t.Sub(start))

  start = time.Now()
  r = Find(n, "st")
  t = time.Now()
  fmt.Println("st", r, t.Sub(start))
}
