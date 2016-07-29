package setFlowDataValues

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/30x/gozerian/pipeline"
	"github.com/30x/gozerian/test_util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetFlowDataValues", func() {

	It("should add values from config", func() {
		//use the pipe fitting to create the pipeline
		pipeline := createPipeline(&testFitting{})

		//get our request handler
		requestHandler := pipeline.RequestHandlerFunc()

		//create a mock request
		req, err := http.NewRequest("GET", "http://example.com/", nil)
		if err != nil {
			log.Fatal(err)
		}

		w := httptest.NewRecorder()
		requestHandler(w, req)
	})

})

//mock endpoint url. "http://localhost:8181/verifiers/apikey"
func createPipeline(testFitting pipeline.FittingWithID) pipeline.Pipe {

	conf := make(map[interface{}]interface{})
	conf["testKeyA"] = "testValueA"
	conf["testKeyB"] = []string{"testValueB1", "testValueB2"}

	fitting, err := CreateFitting(conf)

	Expect(err).Should(BeNil(), "Error creating fitting")
	Expect(fitting).ShouldNot(BeNil(), "Error creating fitting")

	//create our handler from our fitting
	handler := fitting.RequestHandlerFunc()

	reqFittings := []pipeline.FittingWithID{
		test_util.NewFittingFromHandlers("test", handler, nil),
		testFitting,
	}

	pipeDef := pipeline.NewDefinition(reqFittings, []pipeline.FittingWithID{})

	pipe := pipeDef.CreatePipe(fmt.Sprintf("%d", time.Now().UnixNano()))

	return pipe
}

type testFitting struct {
}

func (f *testFitting) RequestHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		control := w.(pipeline.ControlHolder).Control()
		flowData := control.FlowData()
		Expect(flowData["testKeyA"]).To(Equal("testValueA"))
		Expect(flowData["testKeyB"]).To(Equal([]string{"testValueB1", "testValueB2"}))
	}
}

func (f *testFitting) ResponseHandlerFunc() pipeline.ResponseHandlerFunc {
	return nil
}

func (f *testFitting) ID() string {
	return "test"
}