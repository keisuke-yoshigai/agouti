package agouti

import (
	"github.com/onsi/ginkgo"
	"github.com/sclevine/agouti/phantom"
	"github.com/sclevine/agouti/webdriver"
	"time"
)

const PHANTOM_HOST = "127.0.0.1"
const PHANTOM_PORT = 8910

var phantomService *phantom.Service

func SetupAgouti() bool {
	phantomService = &phantom.Service{Host: PHANTOM_HOST, Port: PHANTOM_PORT, Timeout: time.Second}
	phantomService.Start()
	return true
}

func CleanupAgouti(ignored bool) bool {
	phantomService.Stop()
	return true
}

func Feature(text string, body func()) bool {
	return ginkgo.Describe(text, body)
}

func Background(body interface{}, timeout ...float64) bool {
	return ginkgo.BeforeEach(body, timeout...)
}

func Scenario(description string, body func()) bool {
	return ginkgo.It(description, body)
}

func Step(description string, body func()) {
	body()
}

func Navigate(url string) *Page {
	session, err := phantom.CreateSession()
	// TODO: test error
	if err != nil {
		ginkgo.Fail(err)
	}

	driver := &webdriver.Driver{session}

	driver.Navigate(url)
	return &Page{driver}
}

