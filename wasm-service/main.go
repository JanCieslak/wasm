package main

import (
	"os/exec"
)

func main() {
  command := exec.Command("asc", "./code/asc/test.ts", "--outFile", "./code/asc/test.wasm", "--optimize")
  err := command.Run()
  if err != nil {
    panic(err)
  }

  // output, err := command.Output()
  // if err != nil {
  //   panic(err)
  // }

  // fmt.Printf("Sucessfully compiled to wasm, %s \n", output)
}
