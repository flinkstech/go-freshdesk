package querybuilder

import (
	"fmt"
	"net/url"
	"strings"
)

type Query struct {
	query string
}

func (query Query) URLSafe() string {
	v := url.Values{}
	var clean string
	if query.query[0] == byte('(') {
		clean = query.query[1:]
	}
	if clean[len(clean)-1] == byte(')') {
		clean = clean[:len(clean)-1]
	}
	v["query"] = []string{fmt.Sprintf(`"%s"`, clean)}
	return strings.Replace(v.Encode(), "+", "%20", -1)
}

type leftQuery struct {
	leftPart string
}

func BuildQuery(q Query) string {
	return q.query
}

func Parameter(parameter string) leftQuery {
	return leftQuery{
		leftPart: parameter,
	}
}

func AllOf(queries ...Query) Query {
	n := len(queries)
	out := "("
	for i, q := range queries {
		out += q.query
		if i < (n - 1) {
			out += " AND "
		}
	}
	out += ")"
	return Query{
		query: out,
	}
}

func AnyOf(queries ...Query) Query {
	n := len(queries)
	out := "("
	for i, q := range queries {
		out += q.query
		if i < (n - 1) {
			out += " OR "
		}
	}
	out += ")"
	return Query{
		query: out,
	}
}

func (left leftQuery) Equals(value int) Query {
	return Query{
		query: fmt.Sprintf("%s:%d", left.leftPart, value),
	}
}

func (left leftQuery) Is(value string) Query {
	return Query{
		query: fmt.Sprintf("%s:'%s'", left.leftPart, value),
	}
}

func (left leftQuery) GreaterThan(value int) Query {
	return Query{
		query: fmt.Sprintf("%s:>%d", left.leftPart, value),
	}
}

func (left leftQuery) LessThan(value int) Query {
	return Query{
		query: fmt.Sprintf("%s:<%d", left.leftPart, value),
	}
}

func (left leftQuery) GreaterThanString(value string) Query {
	return Query{
		query: fmt.Sprintf("%s:>'%s'", left.leftPart, value),
	}
}

func (left leftQuery) LessThanString(value string) Query {
	return Query{
		query: fmt.Sprintf("%s:<'%s'", left.leftPart, value),
	}
}

func (left leftQuery) IsTrue() Query {
	return Query{
		query: fmt.Sprintf("%s:true", left.leftPart),
	}
}

func (left leftQuery) IsFalse() Query {
	return Query{
		query: fmt.Sprintf("%s:false", left.leftPart),
	}
}

func (left leftQuery) IsNil() Query {
	return Query{
		query: fmt.Sprintf("%s:null", left.leftPart),
	}
}
