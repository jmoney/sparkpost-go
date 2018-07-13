package sparkpost

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// DefaultTrackingDomain get the default tracking domain for the account
func (sparkpostAPIClient *Client) DefaultTrackingDomain() (string, error) {
	trackingDomains, err := sparkpostAPIClient.TrackingDomains()
	if err == nil {
		for _, trackingDomain := range *trackingDomains {
			if trackingDomain.Default {
				return trackingDomain.Domain, nil
			}
		}
	}
	return "", err
}

// TrackingDomains retrieve the tracking domains for the account
func (sparkpostAPIClient *Client) TrackingDomains() (*[]TrackingDomain, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", sparkpostAPIClient.apiURL, "tracking-domains"), nil)
	req.Header.Add("Authorization", sparkpostAPIClient.apiKey)
	resp, err := sparkpostAPIClient.http.Do(req)
	if err != nil {
		return &[]TrackingDomain{}, err
	}
	defer resp.Body.Close()

	results := &TrackingDomainResults{}
	rawJSON, err := ioutil.ReadAll(resp.Body)
	json.NewDecoder(strings.NewReader(fmt.Sprintf("%s", rawJSON))).Decode(results)

	return &(*results).Results, nil
}

// SendingDomains returns the sending domain results from sparkpost
func (sparkpostAPIClient *Client) SendingDomains() (*[]SendingDomainResult, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", sparkpostAPIClient.apiURL, "sending-domains"), nil)
	req.Header.Add("Authorization", sparkpostAPIClient.apiKey)
	resp, err := sparkpostAPIClient.http.Do(req)
	if err != nil {
		return &[]SendingDomainResult{}, err
	}
	defer resp.Body.Close()

	results := &SendingDomains{}
	rawJSON, err := ioutil.ReadAll(resp.Body)
	json.NewDecoder(strings.NewReader(fmt.Sprintf("%s", rawJSON))).Decode(results)

	return &(*results).SendingDomainResult, nil
}

// SendingDomain returns the sending domain results from sparkpost
func (sparkpostAPIClient *Client) SendingDomain(sendingDomain string) (*SendingDomainResult, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", sparkpostAPIClient.apiURL, "sending-domains", sendingDomain), nil)
	req.Header.Add("Authorization", sparkpostAPIClient.apiKey)
	resp, err := sparkpostAPIClient.http.Do(req)
	if err != nil {
		return &SendingDomainResult{}, err
	}
	defer resp.Body.Close()

	results := &SendingDomain{}
	rawJSON, err := ioutil.ReadAll(resp.Body)
	json.NewDecoder(strings.NewReader(fmt.Sprintf("%s", rawJSON))).Decode(results)

	return &(*results).SendingDomainResult, nil
}
