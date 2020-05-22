package payment_gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PaymentGatewayResponse struct {
	AccountNo string
	Amount    float32
	Success   bool
	Reason    string
}

func Request(accountNo string, amount float32) PaymentGatewayResponse {
	payload := []byte(fmt.Sprintf("`{\"accountNo\": \"%s\", \"amount\": %f}`", accountNo, amount))
	response, err := http.Post("https://fintech.com/external/pay", "text/json", bytes.NewBuffer(payload))
	if err != nil {
		panic("")
	}
	defer response.Body.Close()
	data := PaymentGatewayResponse{}
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic("There was an error during parsing json from provider.\n" + err.Error())
	}
	return data
}
