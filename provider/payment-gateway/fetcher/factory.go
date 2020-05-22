package fetcher

type Factory interface {
	submitHtml()
	doLogin()
	downloadAndParse() []Mutation
}
