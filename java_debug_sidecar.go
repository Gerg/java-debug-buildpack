package main

import (
  "fmt"
  "net/http"
  "os/exec"
  "os"
)

func threadDump(w http.ResponseWriter, req *http.Request) {
  portCmd := exec.Command("bash", "-c", "netstat -lntp tcp | grep 8080 | awk '{print $7}' | egrep -o1 '[0-9]+'")
  portCmd.Stderr = os.Stderr
  port, err := portCmd.Output()
  if err != nil {
    fmt.Println("Netstat Error: ", port, err)
  }

  portStr := string(port[:])


  killCommand := exec.Command("bash", "-c", fmt.Sprintf("kill -3 %s", portStr))
  killCommand.Stderr = os.Stderr
  fmt.Println(killCommand.String())
  dump, err := killCommand.Output()
  if err != nil {
    fmt.Println("Kill Error: ", dump, err)
  }

  fmt.Fprintf(w, string(dump[:]) )
}

func main() {
  http.HandleFunc("/threaddump", threadDump)
  http.ListenAndServe(":8081", nil)
}
