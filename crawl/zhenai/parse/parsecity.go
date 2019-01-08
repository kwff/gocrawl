package parse

import (
	"crawl/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{
		Requests: nil,
		Items:     nil,
	}
	for _, m := range matches {
		name:=string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c,name)
			},
		})

	}
	return result
}