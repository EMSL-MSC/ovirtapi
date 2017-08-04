package ovirtapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Link struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
	Id   string `json:"id,omitempty"`
}

type Connection struct {
	EndPoint       *url.URL
	UserName       string
	Password       string
	Debug          bool
	Links          []Link `json:"link"`
	SpecialObjects struct {
		Links []Link `json:"link"`
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

type RequestError struct {
	StatusCode int
	Detail     string `json:"detail"`
	Reason     string `json:"reason"`
}

func (f RequestError) Error() string {
	if f.Reason != "" && f.Detail != "" {
		return fmt.Sprintf("Server Error, Reason: %s, Detail: %s", f.Reason, f.Detail)
	} else {
		return fmt.Sprintf("Error getting response from server (Response code %d )", f.StatusCode)
	}
}

func NewConnection(endpoint string, username string, password string) (*Connection, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.New("Error parsing endpoint URL")
	}
	con := &Connection{
		EndPoint: endpointURL,
		UserName: username,
		Password: password,
	}
	body, err := con.Request("GET", endpointURL, nil)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &con)
	return con, nil
}

func (con *Connection) ResolveLink(link string) *url.URL {
	return con.EndPoint.ResolveReference(&url.URL{Path: link})
}

func (con *Connection) Request(verb string, requestURL *url.URL, reqBody []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(verb, requestURL.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	if reqBody != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(con.UserName, con.Password)
	if con.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println(">", strings.Replace(strings.Replace(string(dump), "\r\n", "\n", -1), "\n", "\n> ", -1))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if con.Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Println("<", strings.Replace(strings.Replace(string(dump), "\r\n", "\n", -1), "\n", "\n< ", -1))
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fault := RequestError{resp.StatusCode, "", ""}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(respBody, &fault)
		}
		return nil, fault
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (con *Connection) GetLink(rel string) (*url.URL, error) {
	for _, link := range con.Links {
		if strings.ToLower(rel) == link.Rel {
			return con.ResolveLink(link.Href), nil
		}
	}
	return nil, fmt.Errorf("con does not have %s link", rel)

}

func (con *Connection) GetLinkBody(link string, id string) ([]byte, error) {
	url, err := con.GetLink(link)
	if err != nil {
		return nil, err
	}
	if id != "" {
		url.Path += "/" + id
	}
	body, err := con.Request("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return body, nil
}
