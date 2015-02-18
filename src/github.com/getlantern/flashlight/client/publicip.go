package client

import (
	"io/ioutil"
	"net/http"

	"github.com/getlantern/fronted"
)

func lookupPublicIp(fd fronted.Dialer) {
	client := fd.DirectHttpClient()
	resp, err := client.Get("http://geo.getiantem.org/lookup")
	if err != nil {
		log.Errorf("Unable to lookup public ip: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var bodyString = ""
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			bodyString = string(body)
		}
		log.Errorf("Unexpected response status %d: %v", resp.StatusCode, bodyString)
		return
	}
	log.Debugf("Public ip is: %v", resp.Header.Get("X-Reflected-Ip"))
}
