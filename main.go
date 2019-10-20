package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"log"

	"github.com/daragao/sibsapi/client"
)

func printConsentSteps(sibsClient *client.Client) {
	aspspCde := "CGDPT"
	// bytesResp, err := sibsClient.GetAccount(aspspCde, "71525dacac1763b812af4e83af61847")
	bytesResp, err := sibsClient.GetAccount(aspspCde, "71525dacac1763b812af4e83af61848")
	// bytesResp, err := sibsClient.GetBalances(aspspCde, "71525dacac1763b812af4e83af61847")
	// bytesResp, err := sibsClient.GetTransactions(aspspCde, "71525dacac1763b812af4e83af61847")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytesResp))

	var account struct {
		Account client.AccountReference `json:"account"`
	}
	err = json.Unmarshal(bytesResp, &account)
	if err != nil {
		log.Fatal(err)
	}

	consentPayload := client.ConsentPayload{
		RecurringIndicator:       false,
		ValidUntil:               "2019-11-18T00:00:00Z", //"1900-01-01T00:00:00Z", // TODO: Make this date dynamic
		FrequencyPerDay:          1,
		CombinedServiceIndicator: false,
	}
	consentPayload.Access.Accounts = []client.AccountReference{account.Account}
	consentPayload.Access.Balances = []client.AccountReference{account.Account}
	consentPayload.Access.Transactions = []client.AccountReference{account.Account}
	// consentPayload.Access.AvailableAccounts = "all-accounts"
	// consentPayload.Access.AllPSD2 = "all-accounts"

	//fmt.Printf("%+v\n", consentPayload)

	bytesResp, err = sibsClient.NewConsent(aspspCde, consentPayload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytesResp))

	var consentResponse client.ConsentResponseResource
	err = json.Unmarshal(bytesResp, &consentResponse)
	if err != nil {
		log.Fatal(err)
	}

	bytesResp, err = sibsClient.GetConsent(aspspCde, consentResponse.ConsentID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytesResp))
}

func listAvailableASPSP(sibsClient *client.Client) []client.Aspsp {
	bytesResp, err := sibsClient.ListAvailableASPSP()
	if err != nil {
		log.Fatal(err)
	}

	var available struct {
		AspspList []client.Aspsp `json:"aspsp-list"`
	}
	err = json.Unmarshal(bytesResp, &available)
	if err != nil {
		log.Fatal(err)
	}

	return available.AspspList
}

func listAccounts(sibsClient *client.Client, aspspCde string) []client.Account {
	bytesResp, err := sibsClient.ListAccounts(aspspCde)
	if err != nil {
		log.Fatal(err)
	}

	var accountList struct {
		AccountList []client.Account `json:"accountList"`
	}
	err = json.Unmarshal(bytesResp, &accountList)
	if err != nil {
		log.Fatal(err)
	}
	return accountList.AccountList
}

func loadConfiguration(file string) (string, string) {
	var config struct {
		Secret   string `json:"secret,omitempty"`
		ClientID string `json:"client-id,omitempty"`
	}
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config.ClientID, config.Secret
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	clientID, _ := loadConfiguration("config.json")

	sibsClient := &client.Client{http.DefaultClient, clientID}

	// printConsentSteps(sibsClient)

	/*
		allASPSP := listAvailableASPSP(sibsClient)
		for _, value := range allASPSP {
			fmt.Println(value.Name)
		}
	*/

	/*
		accountList := listAccounts(sibsClient, "CGDPT")
		for _, value := range accountList {
			fmt.Println(value.Name)
		}
	*/

	allASPSP := listAvailableASPSP(sibsClient)
	for _, aspsp := range allASPSP {
		fmt.Printf("%s\t%s\n", aspsp.ASPSPCDE, aspsp.Name)
		// fmt.Printf("\t\t%+v\n", aspsp)
		accountList := listAccounts(sibsClient, aspsp.ASPSPCDE)
		for _, acc := range accountList {
			fmt.Printf("\t\t%s\n", acc.Name)
			// fmt.Printf("\t\t%+v\n", acc)
		}
	}
}
