package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default returns an Engine instance with the Logger and Recovery middleware already attached

	r := gin.Default()

	channelItem := make(chan string)
	channelItem2 := make(chan string)

	bruh := func(itemID string) {
		newItem := new(ItemResponse)
		apiCall := "https://api.mercadolibre.com/items/" + itemID
		response, err := http.Get(apiCall)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			json.Unmarshal(data, newItem)
			fmt.Println("Br0000000000000000000000000000000000")
			fmt.Println(newItem.SiteID)
			fmt.Println(newItem.SellerID)
			fmt.Println(newItem.CategoryID)
			channelItem <- (string(data))
		}
	}

	sis := func() {
		time.Sleep(time.Second * 10)
		channelItem2 <- "response de la api 2"
	}

	//GET of an single item extended data
	r.GET("/show/:itemID", func(c *gin.Context) {
		itemID := c.Param("itemID")
		bruh(itemID)
		go sis()
		select {
		case primerParte := <-channelItem:
			fmt.Println("Primera parte:", primerParte)
		case segundaParte := <-channelItem2:
			fmt.Println("Segunda parte:", segundaParte)
		}
		time.Sleep(time.Second * 10)
		c.JSON(200, gin.H{
			"asd": itemID,
		})
	})
	r.Run()
}

//MakePool Creates pool
func MakePool(poolSize, jobs int) (pool, jobDone, allDone chan bool) {
	// Channel to coordinate the number of concurrent goroutines.
	pool = make(chan bool, poolSize)

	// Fill the channel with tokens.
	for i := 0; i < poolSize; i++ {
		pool <- true
	}

	jobDone = make(chan bool)
	allDone = make(chan bool)

	go func() {
		for i := 0; i < jobs; i++ {
			// Add a token each time a job is done
			<-jobDone
			pool <- true
		}
		allDone <- true
	}()

	return
}

/*
Endpoints de utilidad:
https://api.mercadolibre.com/items/:item_id
https://api.mercadolibre.com/sites/:site_id
https://api.mercadolibre.com/users/:user_id
https://api.mercadolibre.com/categories/:category_id
*/

//ItemResponse because not dynamic?
type ItemResponse struct {
	ID                           string        `json:"id"`
	SiteID                       string        `json:"site_id"`
	Title                        string        `json:"title"`
	Subtitle                     interface{}   `json:"subtitle"`
	SellerID                     int           `json:"seller_id"`
	CategoryID                   string        `json:"category_id"`
	OfficialStoreID              interface{}   `json:"official_store_id"`
	Price                        int           `json:"price"`
	BasePrice                    int           `json:"base_price"`
	OriginalPrice                interface{}   `json:"original_price"`
	CurrencyID                   string        `json:"currency_id"`
	InitialQuantity              int           `json:"initial_quantity"`
	AvailableQuantity            int           `json:"available_quantity"`
	SoldQuantity                 int           `json:"sold_quantity"`
	SaleTerms                    []interface{} `json:"sale_terms"`
	BuyingMode                   string        `json:"buying_mode"`
	ListingTypeID                string        `json:"listing_type_id"`
	StartTime                    time.Time     `json:"start_time"`
	StopTime                     time.Time     `json:"stop_time"`
	Condition                    string        `json:"condition"`
	Permalink                    string        `json:"permalink"`
	Thumbnail                    string        `json:"thumbnail"`
	SecureThumbnail              string        `json:"secure_thumbnail"`
	Pictures                     []interface{} `json:"pictures"`
	VideoID                      string        `json:"video_id"`
	Descriptions                 []interface{} `json:"descriptions"`
	AcceptsMercadopago           bool          `json:"accepts_mercadopago"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping                     struct {
	} `json:"shipping"`
	InternationalDeliveryMode string `json:"international_delivery_mode"`
	SellerAddress             struct {
	} `json:"seller_address"`
	SellerContact interface{} `json:"seller_contact"`
	Location      struct {
	} `json:"location"`
	Geolocation struct {
	} `json:"geolocation"`
	CoverageAreas       []interface{} `json:"coverage_areas"`
	Attributes          []interface{} `json:"attributes"`
	Warnings            []interface{} `json:"warnings"`
	ListingSource       string        `json:"listing_source"`
	Variations          []interface{} `json:"variations"`
	Status              string        `json:"status"`
	SubStatus           []interface{} `json:"sub_status"`
	Tags                []interface{} `json:"tags"`
	Warranty            string        `json:"warranty"`
	CatalogProductID    interface{}   `json:"catalog_product_id"`
	DomainID            string        `json:"domain_id"`
	ParentItemID        interface{}   `json:"parent_item_id"`
	DifferentialPricing interface{}   `json:"differential_pricing"`
	DealIds             []interface{} `json:"deal_ids"`
	AutomaticRelist     bool          `json:"automatic_relist"`
	DateCreated         time.Time     `json:"date_created"`
	LastUpdated         time.Time     `json:"last_updated"`
	Health              float64       `json:"health"`
	CatalogListing      bool          `json:"catalog_listing"`
}

//Item is the struct to be used when u want 2 return the full extended data of the item
type Item struct {
	ID                string
	SiteID            Site
	Title             string
	Subtitle          string
	SellerID          User
	CategoryID        Category
	Price             float32
	BasePrice         float32
	OriginalPrice     float32
	CurrencyID        string
	InitialQuantity   int
	AvailableQuantity int
	SoldQuantity      int
	DateCreated       string
	LastUpdated       string
}

//Site is the struct that contains the site info
type Site struct {
	SiteID             string
	ID                 string
	Name               string
	CountryID          string
	SaleFeesMode       string
	MercadopagoVersion int
}

//User is the struct that contains user info
type User struct {
	SellerID         int
	ID               int
	Nickname         string
	RegistrationDate string
}

//Category is the struct that cointains category info
type Category struct {
	CategoryID               string
	ID                       string
	Name                     string
	TotalItemsInThisCategory int
	Picture                  string
}
