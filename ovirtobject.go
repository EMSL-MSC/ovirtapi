package ovirtapi

import (
	"errors"
	"net/url"
)

type OvirtObject struct {
	Link
	Api         *API   `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Actions     struct {
		Links     []Link `json:"link,omitempty"`
	} `json:"actions,omitempty"`
	Links       []Link `json:"link,omitempty"`
}

func (ovirtObject *OvirtObject) DoAction(action string) (err error) {
	for _, link := range ovirtObject.Actions.Links {
		if link.Rel == action {
			_, err = ovirtObject.Api.Request("POST", ovirtObject.Api.ResolveLink(link.Href), nil)
			return
		}
	}
	return errors.New("Action not found")
}

func (ovirtObject *OvirtObject) GetLink(rel string) (*url.URL, error) {
	for _, link := range ovirtObject.Links {
		if rel == link.Rel {
			return ovirtObject.Api.ResolveLink(link.Href), nil
		}
	}
	return nil, errors.New("Link not found")
}

func (ovirtObject *OvirtObject) Delete() error {
	_, err := ovirtObject.Api.Request("DELETE", ovirtObject.Api.ResolveLink(ovirtObject.Href), nil)
	return err
}
