package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/wostzone/directory/internal"
	"github.com/wostzone/gateway/pkg/lib"
)

var directoryConfig = &internal.DirectoryConfig{}

func main() {
	gatewayConfig, err := lib.SetupConfig("", internal.DirectoryPluginID, directoryConfig)

	svc := internal.NewDirectoryService()
	err = svc.Start(gatewayConfig, directoryConfig)
	if err != nil {
		logrus.Errorf("Directory Service: Failed to start")
		os.Exit(1)
	}
	lib.WaitForSignal()
	svc.Stop()
	os.Exit(0)
}
