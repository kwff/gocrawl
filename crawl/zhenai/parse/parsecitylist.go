package parse

import (
	"crawl/engine"

	"regexp"
)

const urlRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(urlRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{
		Requests: nil,
		Items:    nil,
	}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})

	}
	return result
}
