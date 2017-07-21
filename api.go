package ovirtapi

import (
	"strings"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"bytes"
)

type Link struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr,omitempty"`
	Id   string `xml:"id,attr,omitempty"`
}

type API struct {
	EndPoint       *url.URL
	UserName       string
	Password       string
	Debug          bool
	Links          []Link `xml:"link"`
	SpecialObjects struct {
		Links Link `xml:"link"`
	} `xml:"special_objects"`
	ProductInfo struct {
		Name    string `xml:"name"`
		Vendor  string `xml:"vendor"`
		Version struct {
			Major       int    `xml:"major"`
			Minor       int    `xml:"minor"`
			Revision    int    `xml:"revision"`
			Build       int    `xml:"build"`
			FullVersion string `xml:"full_version"`
		} `xml:"version"`
	} `xml:"product_info"`
	Summary struct {
		Vms struct {
			Active int `xml:"active"`
			Total  int `xml:"total"`
		} `xml:"vms"`
		Hosts struct {
			Active int `xml:"active"`
			Total  int `xml:"total"`
		} `xml:"hosts"`
		Users struct {
			Active int `xml:"active"`
			Total  int `xml:"total"`
		} `xml:"users"`
		StorageDomains struct {
			Active int `xml:"active"`
			Total  int `xml:"total"`
		} `xml:"storage_domains"`
	} `xml:"summary"`
}

type Fault struct {
	XMLName xml.Name `xml:"fault"`
	Detail string `xml:"detail"`
	Reason string `xml:"reason"`
}

func NewAPI(endpoint string, username string, password string) (*API, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.New("Error parsing endpoint URL")
	}
	api := &API{
		EndPoint: endpointURL,
		UserName: username,
		Password: password,
	}
	body, err := api.Request("GET", endpointURL, nil)
	if err != nil {
		return nil, errors.New("Error connecting to endpoint")
	}
	xml.Unmarshal(body, &api)
	return api, nil
}

func (api *API) ResolveLink(link string) (*url.URL) {
	return api.EndPoint.ResolveReference(&url.URL{Path: link})
}

func (api *API) Request(verb string, requestURL *url.URL, reqBody []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(verb, requestURL.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	if reqBody != nil {
		req.Header.Add("Content-Type", "application/xml")
	}
	req.SetBasicAuth(api.UserName, api.Password)
	if api.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(">", strings.Replace(strings.Replace(string(dump), "\r\n", "\n", -1), "\n", "\n> ", -1))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if api.Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Println("<", strings.Replace(strings.Replace(string(dump), "\r\n", "\n", -1), "\n", "\n< ", -1))
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fault := Fault{}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			xml.Unmarshal(respBody, &fault)
			if fault.Reason != "" {
				return nil, fmt.Errorf("Server Error, Reason: %s, Detail: %s", fault.Reason, fault.Detail)
			}
		}
		return nil, fmt.Errorf("Error getting response from server (Response code %d )", resp.StatusCode)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (api *API) GetLink(rel string) (*url.URL, error) {
	for _, link := range api.Links {
		if rel == link.Rel {
			return api.ResolveLink(link.Href), nil
		}
	}
	return nil, errors.New("Link not found")
}

func (api *API) GetLinkBody(link string, id string) ([]byte, error) {
	url, err := api.GetLink(link)
	if err != nil {
		return nil, errors.New("API missing vms url")
	}
	if id != "" {
		url.Path += "/" + id
	}
	body, err := api.Request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return body, nil
}
