package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	db := InitDb()

	//////    SIMPLE QUERY     //////
	// getOneItemByName(db, "sakit")
	// getItemsByName(db, "obat")
	// getCustomers(db, 5)
	// getCustomersOffset(db, 5, 5)
	// getItemsWithTwoParams(db, "obat", "sakit")

	//////    GROUP HAVING     //////
	// getPurchaseDetailsMoreThan20(db)
	// getTotalQty(db)

	//////    CREATE, UPDATE, DELETE    //////
	// createItem(db)
	// incrementQtyAllItems(db)
	// softDelete(db)
	getDetailsHeader(db)
	// getItemCategory(db)
	// test(db)
}

func printResult(data any, result *gorm.DB) {
	if result.Error != nil {
		fmt.Println("error: ", result.Error)
	}
	fmt.Println("rows affected: ", result.RowsAffected)
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}
