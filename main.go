package main

import (
	"fmt"
	"time"
	// "io"
	// "io/ioutil"
	// "log"
	"math/rand"
	"os"
	// "runtime"
	// "sync"
	// "syscall"
	// "time"
	//
	"github.com/ehllo-com/worker/cmd"
	// "github.com/hashicorp/go-uuid"
	// "github.com/hashicorp/packer/command"
	// "github.com/hashicorp/packer/packer"
	// "github.com/hashicorp/packer/packer/plugin"
	// "github.com/hashicorp/packer/packer/tmp"
	// "github.com/hashicorp/packer/version"
	// "github.com/mitchellh/cli"
	// "github.com/mitchellh/panicwrap"
	// "github.com/mitchellh/prefixedio"
)

var (
	VERSION = "0.0.1"
)

func main() {
	fmt.Println("==> main starting...")
	os.Setenv("WORKER_UUID", "1234-5678-9101-1121")

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())
}
