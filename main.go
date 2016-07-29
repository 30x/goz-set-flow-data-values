package main

import (
	"github.com/30x/gozerian/go_gateway"
	"github.com/30x/gozerian/pipeline"
	"github.com/30x/goz-set-flow-data-values/setFlowDataValues"
	"os"
	"fmt"
)

// Just an example for testing this plugin
// Install the node Edge Micro proxy on Edge first and configure via main.yaml

func main() {
	pipeline.RegisterDie("setFlowDataValues", setFlowDataValues.CreateFitting)

	yamlReader, err := os.Open("main.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = go_gateway.ListenAndServe(yamlReader)
	if err != nil {
		fmt.Println(err)
	}
}
