package parse

import (
	"crawl/engine"
	"crawl/model"
	"regexp"
)

//var nameRe = regexp.MustCompile(`<span class="nickName"[^>]*>([^<]+)</span>`)
var topRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div>`) //ageRe
var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var footRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^<]+)</div>`) //hukouRe
var hukouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{
		Name:     "",
		Marriage: "",
		Age:      "",
		Height:   "",
		Weight:   "",
		Income:   "",
		Hukou:    "",
	}

	//profile.Name = extractString(contents, nameRe)
	profile.Name = name
	//profile.Purple = extractByte(contents, topRe)[6:]
	profile.Marriage = extractString(contents, marriageRe)
	profile.Age = extractString(contents, ageRe)
	//profile.Xinzuo = extractByte(contents, topRe)[2]
	profile.Height = extractString(contents, heightRe)
	profile.Weight = extractString(contents, weightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Hukou = extractString(contents, hukouRe)

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 1 {
		return string(matches[1])
	} else {
		return ""
	}

}

func extractByte(contents []byte, re *regexp.Regexp) []string {
	matches := re.FindAllSubmatch(contents, -1)
	var result []string
	for _, m := range matches {
		result = append(result, string(m[1]))
	}

	return result

}
