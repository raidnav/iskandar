package fetcher

type BCA struct{}

func (bca *BCA) submitHtml() {

}

func (bca *BCA) doLogin() {

}

func (bca *BCA) downloadAndParse() []Mutation {
	return []Mutation{}
}

func NewBCAFetcher() Factory {
	return &BCA{}
}
