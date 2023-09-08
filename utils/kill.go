package utils

import (
	"log"
	"os/exec"
  "github.com/abdfnx/gosh"
)

func findProc(name string) string {

    out, err := exec.Command("pgrep", name).Output()

    if err != nil {
        log.Fatal(err)
    }

    return string(out)
}

func Kill(name string) {
  id := findProc(name)

  gosh.Run("kill " + id)
}
