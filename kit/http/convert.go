package http

import (
	"context"
	"fmt"
	"github.com/mikhailbolshakov/decision/kit"
	"regexp"
	"strings"
)

const (
	asc   = "asc"
	desc  = "desc"
	first = "first"
	last  = "last"
)

var sortDirections = map[string]bool{
	"":   true,
	asc:  true,
	desc: false,
}

var sortMissings = map[string]string{
	"":    "",
	first: kit.SortRequestMissingFirst,
	last:  kit.SortRequestMissingLast,
}

// ParseSortBy Converts string like "field1 asc first,field2 desc last,field3 desc,field4" to array of SortRequest
func ParseSortBy(ctx context.Context, sortString string) ([]*kit.SortRequest, error) {
	if sortString == "" {
		return nil, nil
	}

	elements := strings.Split(sortString, ",")

	var res []*kit.SortRequest
	ruleRegex := regexp.MustCompile(fmt.Sprintf("^([a-zA-Z][a-zA-Z0-9]+)(?: (%s|%s)(?: (%s|%s))?)?$", asc, desc, first, last))
	for _, elem := range elements {

		elemParts := ruleRegex.FindStringSubmatch(elem)
		if elemParts == nil {
			return nil, ErrHttpUrlWrongSortFormat(ctx, sortString)
		}

		sortRq := &kit.SortRequest{Field: elemParts[1]}

		sortDirection, ok := sortDirections[elemParts[2]]
		if !ok {
			return nil, ErrHttpUrlWrongSortFormat(ctx, sortString)
		}
		sortRq.Asc = sortDirection

		missing, ok := sortMissings[elemParts[3]]
		if !ok {
			return nil, ErrHttpUrlWrongSortFormat(ctx, sortString)
		}
		sortRq.Missing = missing

		res = append(res, sortRq)
	}
	return res, nil
}
