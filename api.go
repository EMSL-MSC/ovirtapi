package ovirtapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel,omitempty"`
	Id   string `json:"id,omitempty"`
}

type API struct {
	EndPoint       *url.URL
	UserName       string
	Password       string
	Debug          bool
	Links          []Link `json:"link"`
	SpecialObjects struct {
		Links Link `json:"link"`
	} `json:"special_objects"`
	ProductInfo struct {
		Name    string `json:"name"`
		Vendor  string `json:"vendor"`
		Version struct {
			Major       int    `json:"major"`
			Minor       int    `json:"minor"`
			Revision    int    `json:"revision"`
			Build       int    `json:"build"`
			FullVersion string `json:"full_version"`
		} `json:"version"`
	} `json:"product_info"`
	Summary struct {
		Vms struct {
			Active int `json:"active"`
			Total  int `json:"total"`
		} `json:"vms"`
		Hosts struct {
			Active int `json:"active"`
			Total  int `json:"total"`
		} `json:"hosts"`
		Users struct {
			Active int `json:"active"`
			Total  int `json:"total"`
		} `json:"users"`
		StorageDomains struct {
			Active int `json:"active"`
			Total  int `json:"total"`
		} `json:"storage_domains"`
	} `json:"summary"`
}

type Fault struct {
	XMLName xml.Name `json:"fault"`
	Detail  string   `json:"detail"`
	Reason  string   `json:"reason"`
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
		return nil, err
	}
	xml.Unmarshal(body, &api)
	return api, nil
}

func (api *API) ResolveLink(link string) *url.URL {
	return api.EndPoint.ResolveReference(&url.URL{Path: link})
}

func (api *API) Request(verb string, requestURL *url.URL, reqBody []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(verb, requestURL.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	api.Debug = true
	if reqBody != nil {
		req.Header.Add("Content-Type", "application/xml")
	} else {
		req.Header.Add("Accept", "application/json")
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
		if strings.ToLower(rel) == link.Rel {
			return api.ResolveLink(link.Href), nil
		}
	}
	return nil, fmt.Errorf("api does not have %s link", rel)

}

func (api *API) GetLinkBody(link string, id string) ([]byte, error) {
	url, err := api.GetLink(link)
	if err != nil {
		return nil, err
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
