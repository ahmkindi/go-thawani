## github.com/ahmkindi/go-thawani

An unofficial golang SDK for [thawani](https://docs.thawani.om/), just calls there APIs easily switch between dev and prod mode by changing the constants as shown below.

### Usage

#### Initialize your module

`go mod init example.com/test`

#### Get go-thawani

`go get github.com/ahmkindi/go-thawani`

```go
package main

import (
  "fmt"

  "github.com/ahmkindi/go-thawani"
)

const (
  THAWANI_BASE_URL = "https://uatcheckout.thawani.om"
  THAWANI_API_KEY = "rRQ26GcsZzoEhbrP2HZvLYDbn9C9et"
  THAWANI_PUBLISHABLE_KEY = "HGvTMLDssJghr9tlN9gr4DVYt0qyBy"
)

func main() {
	thawaniHost, err := url.Parse(THAWANI_BASE_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse thawani url: %w", err)
	}

	ThawaniClient := thawani.NewClient(
                      &http.Client{Timeout: time.Second * 20},
                      thawaniHost,
                      THAWANI_API_KEY,
                      THAWANI_PUBLISHABLE_KEY
                  )

   customer, err := ThawaniClient.CreateCustomer(thawani.CreateCustomerReq{
     ClientCustomerId: user.ID.String(),
   })
   if err != nil {
     return fmt.Errorf("failed to create customer: %w", err)
   }
   customerId := customer.Data.Id

   fmt.Print("Created a customer with id: ", customerId)
}
```

### Whats left?
There is still loads to implement, feel free to add and make a PR. Currently with the available functions you can:
* create a customer
* create a session
* get a session

Thats enough for a simple workflow.
