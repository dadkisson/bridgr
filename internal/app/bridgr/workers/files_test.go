package workers_test

import (
	"bridgr/internal/app/bridgr/config"
	"bridgr/internal/app/bridgr/workers"
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type HTTPMock struct {
	http.RoundTripper
}

func (m HTTPMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString("OK")),
		Header:     make(http.Header),
	}, nil
}

var defaultConf = config.BridgrConf{
	Files: config.Files{
		Items: []config.FileItem{
			{Source: "/source1", Target: "packages/files/file1", Protocol: "file"},
			{Source: "http://nothing.net/file2", Target: "packages/files/file2", Protocol: "http"},
			{Source: "ftp://nothing.net/file3", Target: "packages/files/file3", Protocol: "ftp"},
		},
	},
}

var stubWorker = workers.Files{
	Config: &defaultConf,
	HTTP:   &http.Client{Transport: &HTTPMock{}},
}

func TestFilesSetup(t *testing.T) {
	err := stubWorker.Setup()
	if err != nil {
		t.Error("Error running Setup")
	}
}

// Can't run this currently because it creates the base directory on the real OS.
// TODO: add a filesytem abstraction to bridgr for better testing os calls.
// func TestFilesRun(t *testing.T) {
// 	err := stubWorker.Run()
// 	if err != nil {
// 		t.Error("Error running Run")
// 	}
// }

func TestFilesName(t *testing.T) {
	f := workers.Files{}
	if f.Name() != "Files" {
		t.Errorf("Files worker does not provide the correct Name() response (%s)", f.Name())
	}
}
