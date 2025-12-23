package problem

import (
	"net/url"
)

var problems = set[ProblemInfo]{}
var problemsDetails = map[string]*set[ProblemInfo]{}

func registerProblem(p ProblemInfo) {
	problems.Add(p.withoutDetails())
	if p.hasDetails() {
		basePath := p.basePath()
		d, _ := problemsDetails[basePath]
		d.Add(p)
		problemsDetails[basePath] = d
	}
}

func (p ProblemInfo) basePath() string {
	u, _ := url.Parse(p.Type)
	return u.Path
}

func RegistryListProblems() []ProblemInfo {
	return problems.List()
}

func RegistryProblemDetails(path string) []ProblemInfo {
	details, ok := problemsDetails[path]
	if !ok || details == nil {
		return nil
	}
	return details.List()
}

func List() []ProblemInfo {
	return problems.List()
}

func ListDetails(path string) []ProblemInfo {
	details, ok := problemsDetails[path]
	if !ok || details == nil {
		return nil
	}
	return details.List()
}
