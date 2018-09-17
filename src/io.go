package main

import (
  "log"
  "os"
  "io/ioutil"
)

func DoesExist(filename string) bool {
  _, err := os.Stat(filename)
  return !os.IsNotExist(err)
}

func ReadFromFile(filename string) string {
  b, err := ioutil.ReadFile(filename) 
  if err != nil {
    log.Fatal(err)
  }
  return string(b)
}

func WriteToFile(filename string, text string) {
  file, err := os.Create(filename)
  if err != nil {
    log.Fatalf("failed creating file: %s", err)
  }
  defer file.Close()
 
  _, err = file.WriteString(text)

  if err != nil {
    log.Fatalf("failed writing to file: %s", err)
  }
}

func CreateDirIfNotExist(dir string) {
  if _, err := os.Stat(dir); os.IsNotExist(err) {
    err = os.MkdirAll(dir, 0755)
    if err != nil {
      panic(err)
    }
  }
}