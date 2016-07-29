/*
A Gozerian plugin that simply injects its configuration values into the request's control.FlowData.

Sample Configuration:

    - addFlowData:
        contextKeyA: contextValueA
        contextKeyB: contextValueB
*/
package setFlowDataValues

import (
	"errors"
	"github.com/30x/gozerian/pipeline"
	"net/http"
)

type setFlowDataValuesFitting struct {
	config map[string]interface{}
}

// CreateFitting exported function to create the fitting
func CreateFitting(config interface{}) (pipeline.Fitting, error) {

	c, ok := config.(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("Invalid config. Expected map[string]interface{}")
	}

	conf := make(map[string]interface{})
	for k, v := range c {
		key, ok := k.(string)
		if !ok {
			return nil, errors.New("Invalid config. Expected map[string]interface{}")
		}
		conf[key] = v
	}

	return &setFlowDataValuesFitting{config: conf}, nil
}

func (f *setFlowDataValuesFitting) RequestHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		control := w.(pipeline.ControlHolder).Control()
		log := control.Log()

		flowData := control.FlowData()
		for k, v := range f.config {
			flowData[k] = v
			log.Debugf("Setting flow var: %s to: %v\n", k, v)
		}
	}
}

func (f *setFlowDataValuesFitting) ResponseHandlerFunc() pipeline.ResponseHandlerFunc {
	return nil
}
