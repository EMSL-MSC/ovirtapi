package ovirtapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type OvirtObject struct {
	Link
	Con         *Connection `json:"-"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Actions     struct {
		Links []Link `json:"link,omitempty"`
	} `json:"actions,omitempty"`
	Links []Link `json:"link,omitempty"`
}

func (ovirtObject *OvirtObject) DoAction(action string, parameters interface{}) (err error) {
	for _, link := range ovirtObject.Actions.Links {
		if link.Rel == action {
			var body []byte
			body, err = json.Marshal(parameters)
			if err != nil {
				return fmt.Errorf("Unable to Marshal parameters")
			}
			_, err = ovirtObject.Con.Request("POST", ovirtObject.Con.ResolveLink(link.Href), body)
			return
		}
	}
	return errors.New("Action not found")
}

func (ovirtObject *OvirtObject) GetLink(rel string) (*url.URL, error) {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			return ovirtObject.Con.ResolveLink(link.Href), nil
		}
	}
	return nil, errors.New("Link not found")
}

func (ovirtObject *OvirtObject) Delete() error {
	_, err := ovirtObject.Con.Request("DELETE", ovirtObject.Con.ResolveLink(ovirtObject.Href), nil)
	return err
}
