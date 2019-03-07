package main

import (
  "log"
  "net/http"
)

func getListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }

  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
