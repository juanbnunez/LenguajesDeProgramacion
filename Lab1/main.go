package main

import (
	"fmt"
	"reflect"
	"time"
)

//-----------------------------------------------------
// Structs

type Transaction struct {
	Type   string
	Amount int
	Date   time.Month
}

type Client struct {
	Id               int
	Name             string
	AccountType      string
	Balance          int
	TransactionsList []Transaction
}

//-----------------------------------------------------

// A collection of Client
type clients []Client

var clientsBD clients

//-----------------------------------------------------
//Validations

// Valid account types
var validAccountTypes = map[string]bool{
	"Savings": true,
	"Salary":  true,
	"Checks":  true,
}

// Valid transaction types
// Not used yet
var validTransactionType = map[string]bool{
	"Deposit":    true,
	"Withdrawal": true,
}

//-----------------------------------------------------

func (c *clients) AddClient(clientPtr *Client) {
	found := false
	if validAccountTypes[clientPtr.AccountType] {
		for n := 0; n < len(*c); n++ {
			if clientPtr.Id == (*c)[n].Id {
				found = true
				client := &(*c)[n]
				if clientPtr.Name != client.Name {
					client.Name = clientPtr.Name
				}
				if clientPtr.AccountType != client.AccountType {
					client.AccountType = clientPtr.AccountType
				}
				if clientPtr.Balance != client.Balance {
					client.Balance = clientPtr.Balance
				}
				if !reflect.DeepEqual(clientPtr.TransactionsList, client.TransactionsList) {
					client.TransactionsList = clientPtr.TransactionsList
				}
			}
		}
		if !found {
			*c = append(*c, *clientPtr)
		}
	} else {
		fmt.Println("Invalid account type:", clientPtr.AccountType)
	}
}

//-----------------------------------------------------

// Function 2: Get clients with more withdrawals than deposits
func GetClientsWithMoreWithdrawals(clientsList clients) clients {
	var result clients
	for _, client := range clientsList {
		withdrawals := 0
		deposits := 0
		for _, transaction := range client.TransactionsList {
			if transaction.Type == "Withdrawal" {
				withdrawals++
			} else if transaction.Type == "Deposit" {
				deposits++
			}
		}
		if withdrawals > deposits {
			result = append(result, client)
		}
	}
	return result
}

// Function 3: Filter transactions for a specific client and month
func FilterTransactionsForClientAndMonth(clientsList clients, clientId int, targetMonth time.Month) []Transaction {
	var result []Transaction
	for _, client := range clientsList {
		if client.Id == clientId {
			for _, transaction := range client.TransactionsList {
				if transaction.Date == targetMonth {
					result = append(result, transaction)
				}
			}
			break
		}
	}
	return result
}

// Function 4: Calculate final balance using map and reduce
func CalculateFinalBalance(clientsList *clients) map[int]int {
	result := make(map[int]int)
	for idx, client := range *clientsList {
		balance := client.Balance
		for _, transaction := range client.TransactionsList {
			if transaction.Type == "Deposit" {
				balance += transaction.Amount
			} else if transaction.Type == "Withdrawal" {
				balance -= transaction.Amount
			}
		}
		(*clientsList)[idx].Balance = balance // Update balance directly through pointer
		result[client.Id] = balance
	}
	return result
}

//-----------------------------------------------------

func PrintClients(clientsBD clients) {
	for _, client := range clientsBD {
		fmt.Printf("Client ID: %d\n", client.Id)
		fmt.Printf("Name: %s\n", client.Name)
		fmt.Printf("Account Type: %s\n", client.AccountType)
		fmt.Printf("Balance: %d\n", client.Balance)
		fmt.Println("Transactions:")
		for _, transaction := range client.TransactionsList {
			fmt.Printf("  Type: %s, Amount: %d, Date: %s\n", transaction.Type, transaction.Amount, transaction.Date)
		}
		fmt.Println("-------------")
	}
}

//-----------------------------------------------------

// FillData adds the clients and transactions
func FillData() {
	//Transactions and clients
	juanTransactions := []Transaction{
		{Type: "Deposit", Amount: 20000, Date: time.January},
		{Type: "Withdrawal", Amount: 5000, Date: time.February},
		{Type: "Deposit", Amount: 10000, Date: time.March},
		{Type: "Withdrawal", Amount: 344324, Date: time.August},
		{Type: "Withdrawal", Amount: 2344354, Date: time.March},
		{Type: "Deposit", Amount: 10000, Date: time.March},
		{Type: "Withdrawal", Amount: 10000, Date: time.March},
		{Type: "Deposit", Amount: 10000, Date: time.March},
	}
	juan := Client{Id: 12345678, Name: "Juan", AccountType: "Savings", Balance: 20000, TransactionsList: juanTransactions}
	clientsBD.AddClient(&juan)

	mariaTransactions := []Transaction{
		{Type: "Deposit", Amount: 20000, Date: time.January},
		{Type: "Withdrawal", Amount: 122, Date: time.January},
		{Type: "Deposit", Amount: 100654433, Date: time.March},
		{Type: "Withdrawal", Amount: 344324, Date: time.January},
		{Type: "Withdrawal", Amount: 22322233, Date: time.July},
		{Type: "Deposit", Amount: 10000, Date: time.March},
		{Type: "Withdrawal", Amount: 10000, Date: time.July},
		{Type: "Deposit", Amount: 4433543, Date: time.March},
	}
	maria := Client{Id: 98765432, Name: "Maria", AccountType: "Savings", Balance: 0, TransactionsList: mariaTransactions}
	clientsBD.AddClient(&maria)

	jennyTransactions := []Transaction{
		{Type: "Withdrawal", Amount: 14547468734, Date: time.January},
		{Type: "Withdrawal", Amount: 46347623, Date: time.January},
		{Type: "Withdrawal", Amount: 263642726, Date: time.March},
		{Type: "Withdrawal", Amount: 378468346, Date: time.January},
		{Type: "Withdrawal", Amount: 46738346, Date: time.July},
		{Type: "Deposit", Amount: 25676246, Date: time.March},
		{Type: "Withdrawal", Amount: 267383, Date: time.July},
		{Type: "Deposit", Amount: 2467835442, Date: time.March},
	}
	jenny := Client{Id: 14546577, Name: "Jenny", AccountType: "Salary", Balance: 0, TransactionsList: jennyTransactions}
	clientsBD.AddClient(&jenny)

	pedroTransactions := []Transaction{
		{Type: "Deposit", Amount: 345643, Date: time.January},
		{Type: "Deposit", Amount: 865438, Date: time.January},
		{Type: "Deposit", Amount: 34567898, Date: time.December},
		{Type: "Deposit", Amount: 23454, Date: time.December},
		{Type: "Withdrawal", Amount: 245764, Date: time.December},
		{Type: "Deposit", Amount: 45769, Date: time.March},
		{Type: "Withdrawal", Amount: 24567, Date: time.July},
		{Type: "Deposit", Amount: 2456794, Date: time.March},
	}
	pedro := Client{Id: 44556788, Name: "Pedro", AccountType: "Salary", Balance: 0, TransactionsList: pedroTransactions}
	clientsBD.AddClient(&pedro)

	joseTransactions := []Transaction{
		{Type: "Deposit", Amount: 345643, Date: time.January},
		{Type: "Withdrawal", Amount: 865438, Date: time.May},
		{Type: "Deposit", Amount: 235875, Date: time.May},
		{Type: "Deposit", Amount: 2343454, Date: time.May},
		{Type: "Withdrawal", Amount: 247885764, Date: time.December},
		{Type: "Deposit", Amount: 4575, Date: time.May},
		{Type: "Withdrawal", Amount: 2222, Date: time.June},
		{Type: "Withdrawal", Amount: 45678546, Date: time.March},
	}
	jose := Client{Id: 77468939, Name: "Jose", AccountType: "Checks", Balance: 0, TransactionsList: joseTransactions}
	clientsBD.AddClient(&jose)

	lizTransactions := []Transaction{
		{Type: "Deposit", Amount: 345643, Date: time.January},
		{Type: "Deposit", Amount: 356485676, Date: time.June},
		{Type: "Withdrawal", Amount: 467978645, Date: time.August},
		{Type: "Withdrawal", Amount: 436586347, Date: time.May},
		{Type: "Withdrawal", Amount: 23465, Date: time.December},
		{Type: "Deposit", Amount: 4357896, Date: time.May},
		{Type: "Withdrawal", Amount: 2567653, Date: time.June},
		{Type: "Withdrawal", Amount: 23567562, Date: time.January},
	}
	liz := Client{Id: 22340045, Name: "Liz", AccountType: "Checks", Balance: 0, TransactionsList: lizTransactions}
	clientsBD.AddClient(&liz)

}

func main() {
	FillData()

	// Original list of clients
	fmt.Println("Original List of Clients:")
	PrintClients(clientsBD)

	// Function 2: Clients with more withdrawals than deposits
	fmt.Println("\nClients with More Withdrawals than Deposits:")
	clientsWithMoreWithdrawals := GetClientsWithMoreWithdrawals(clientsBD)
	fmt.Println(clientsWithMoreWithdrawals)

	// Function 3: Filter transactions for a specific client and month
	clientId := 12345678
	targetMonth := time.March
	fmt.Printf("\nTransactions for Client %d in %s:\n", clientId, targetMonth)
	filteredTransactions := FilterTransactionsForClientAndMonth(clientsBD, clientId, targetMonth)
	fmt.Println(filteredTransactions)

	// Function 4: Calculate final balance
	fmt.Println("\nFinal Balances:")
	finalBalances := CalculateFinalBalance(&clientsBD)
	fmt.Println(finalBalances)

	// After transactions list of clients
	fmt.Println("\nAfter transactions list of Clients:")
	PrintClients(clientsBD)
}
