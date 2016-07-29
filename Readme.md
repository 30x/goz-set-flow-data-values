# setFlowDataValues

A Gozerian plugin that simply injects its configuration values into the request's control.FlowData. Note: Configuration must map to map[string]interface{}.

Sample Configuration:

    - setFlowDataValues:
        flowDataKeyA: flowDataValueA
        flowDataKeyB:
          - flowDataValueB1
          - flowDataValueB2


# Executing tests

(Optional) Install dependencies: `glide install --strip-vcs`

Execute Tests via `go test`
