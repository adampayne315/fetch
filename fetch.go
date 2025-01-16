package main

//go:generate oapi-codegen --config=cfg.yaml api.yml


import "fmt"
import "time"
import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// sample:
/*
"retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
*/

type ReceiptRepo struct {
  Receipts   map[string]Receipt
	NextId string
	Lock   sync.Mutex
}

// Make sure we conform to ServerInterface

var _ StrictServerInterface = (*ReceiptRepo)(nil)

func NewReceiptRepo() *ReceiptRepo {
	return &ReceiptRepo{
		Receipts:   make(map[string]Receipt),
		NextId: "a",
	}
}

// store a new receipt
func (r *ReceiptRepo) PostReceiptsProcess(ctx context.Context, request PostReceiptsProcessRequestObject) (PostReceiptsProcessResponseObject, error) {
	var newReceipt receipt
	// unmarshall the post body
	if err := c.BindJSON(&newReceipt); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
		return
    }
	c.IndentedJSON(http.StatusCreated, newReceipt)

}

func main() {
    router := gin.Default()
    router.POST("/receipts/process", addReceipt)

    router.Run("localhost:8080")
}