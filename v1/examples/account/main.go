package main

import (
	"fmt"
	"github.com/payjp/payjp-go/v1"
	"time"
	"strings"
)

func main() {
	payjpService := payjp.New("sk_test_c62fade9d045b54cd76d7036", nil)

	// Get Account
	fmt.Println("Get Account")
	account, _ := payjpService.Account.Retrieve()
	fmt.Println("  Id:", account.ID)
	fmt.Println("  CreatedAt: ", account.CreatedAt.Format(time.RFC1123Z))
	fmt.Println("  Email: ", account.Email)
	fmt.Println("  Customer:")
	fmt.Println("    Id:", account.Customer.ID)
	fmt.Println("    Email: ", account.Customer.Email)
	fmt.Println("    LiveMode:", account.Customer.LiveMode)
	fmt.Println("    DefaultCard:", account.Customer.DefaultCard)
	fmt.Println("    Description:", account.Customer.Description)
	fmt.Println("    Cards: ", len(account.Customer.Cards))
	fmt.Println("    Subscriptions: ", len(account.Customer.Subscriptions))
	fmt.Println("  Merchant:")
	fmt.Println("    Id", account.Merchant.ID)
	fmt.Println("    BankEnabled:", account.Merchant.BankEnabled)
	fmt.Println("    BrandsAccepted:", strings.Join(account.Merchant.BrandsAccepted, ", "))
	fmt.Println("    ContactPhone:", account.Merchant.ContactPhone)
	fmt.Println("    Country:", account.Merchant.Country)
	fmt.Println("    CurrenciesSupported:", account.Merchant.CurrenciesSupported)
	fmt.Println("    DefaultCurrency:", account.Merchant.DefaultCurrency)
	fmt.Println("    DetailsSubmitted:", account.Merchant.DetailsSubmitted)
	fmt.Println("    LiveModeActivatedAt:", account.Merchant.LiveModeActivatedAt.Format(time.RFC1123Z))
	fmt.Println("    LiveModeEnabled:", account.Merchant.LiveModeEnabled)
	fmt.Println("    ProductDetail:", account.Merchant.ProductDetail)
	fmt.Println("    ProductName:", account.Merchant.ProductName)
	fmt.Println("    SitePublished:", account.Merchant.SitePublished)
	fmt.Println("    URL:", account.Merchant.URL)
}
