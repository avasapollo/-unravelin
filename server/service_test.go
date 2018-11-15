package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/avasapollo/unravelin/data"
	"github.com/avasapollo/unravelin/encoder"
	"github.com/avasapollo/unravelin/printer"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

var (
	le *logrus.Entry
)

func TestMain(m *testing.M) {
	le = logrus.New().WithField("service", "testing")

	os.Exit(m.Run())
}

func TestApiRest_PostForm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	encoderMock := encoder.NewMockHashEncoder(ctrl)
	encoderMock.EXPECT().Encode("https://ravelin.com").Return("hash-test")

	parseMock := data.NewMockParser(ctrl)
	parseMock.EXPECT().ParseMapToData(gomock.Any()).Return(&data.Data{
		WebsiteUrl:         "https://ravelin.com",
		SessionId:          "123123-123123-123123123",
		ResizeFrom:         data.Dimension{},
		ResizeTo:           data.Dimension{},
		CopyAndPaste:       nil,
		FormCompletionTime: 0,
	}, nil)

	testApi := NewApiRest(printer.NewPrinter(le), parseMock, encoderMock)

	req := httptest.NewRequest(http.MethodPost, "/v1/form", bytes.NewBufferString(
		`
			{
    			"author" : "Andrea",
				"eventType": "copyAndPaste",
				"websiteUrl": "https://ravelin.com",
  				"sessionId": "123123-123123-123123123",
  				"pasted": true,
  				"formId": "inputCardNumber"
			}`,
	))

	req.Header.Add("Content-type", "application/json")

	resp := executeRequest(testApi.GetMuxRouter(), req)
	require.Equal(t, http.StatusAccepted, resp.Result().StatusCode)
	output := make(map[string]interface{})
	err := json.Unmarshal(resp.Body.Bytes(), &output)
	require.NoError(t, err)

	webUrl, _ := output["websiteUrl"].(string)
	require.Equal(t, "https://ravelin.com", webUrl)

	require.Equal(t, "123123-123123-123123123", output["sessionId"].(string))
	require.Equal(t, "hash-test", output["hash"].(string))
}

func executeRequest(muxRouter *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	muxRouter.ServeHTTP(rr, req)
	return rr
}
