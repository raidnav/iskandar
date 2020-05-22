package fetcher

type Mandiri struct{}

func (mandiri *Mandiri) submitHtml() {

}

func (mandiri *Mandiri) doLogin() {

}

func (mandiri *Mandiri) downloadAndParse() []Mutation {

	return []Mutation{}
}

func NewMandiriFetcher() Factory {
	return &Mandiri{}
}
