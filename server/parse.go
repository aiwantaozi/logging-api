package server

import (
	"fmt"
	"net/url"

	"github.com/rancher/norman/parse"
	"github.com/rancher/norman/types"
)

func URLParser(schemas *types.Schemas, url *url.URL) (parse.ParsedURL, error) {
	parsedURL, err := parse.DefaultURLParser(schemas, url)
	if err != nil {
		return parse.ParsedURL{}, err
	}
	fmt.Println(parsedURL.ID)
	return parsedURL, nil
}
