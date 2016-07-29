# setFlowDataValues

A Gozerian plugin that simply injects its configuration values into the request's control.FlowData.

Sample Configuration:

    - setFlowDataValues:
        contextKeyA: contextValueA
        contextKeyB: contextValueB


# Executing tests

(Optional) Install dependencies: `glide install --strip-vcs`

Execute Tests via `go test`
