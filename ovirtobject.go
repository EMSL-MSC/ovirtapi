package ovirtapi

import (
	"errors"
	"net/url"
)

type OvirtObject struct {
	Link
	Api			*API `xml:"-"`
	Name    string `xml:"name"`
	Description    string `xml:"description,omitempty"`
	Actions []Link `xml:"actions>link,omitempty"`
	Links   []Link `xml:"link,omitempty"`
}

func (ovirtObject *OvirtObject) DoAction(action string) (err error) {
	for _, link := range ovirtObject.Actions {
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
