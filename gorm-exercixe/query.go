package main

import (
	"fmt"
	"gorm-exercixe/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getOneItemByName(db *gorm.DB, name string) {
	var item entity.Item
	result := db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", name)).First(&item)
	printResult(item, result)
}

func getItemsByName(db *gorm.DB, name string) {
	var items []entity.Item
	result := db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&items)
	printResult(items, result)
}

func getCustomers(db *gorm.DB, amount int) {
	var customers []entity.Customers
	result := db.Order("name").Limit(amount).Find(&customers)
	printResult(customers, result)
}

func getCustomersOffset(db *gorm.DB, amount int, offset int) {
	var customers []entity.Customers
	result := db.Limit(amount).Offset(offset).Find(&customers)
	printResult(customers, result)
}

func getItemsWithTwoParams(db *gorm.DB, param1 string, param2 string) {
	var items []entity.Item
	result := db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", param1)).Where("name NOT ILIKE ?", fmt.Sprintf("%%%s%%", param2)).Find(&items)
	printResult(items, result)
}

func getPurchaseDetailsMoreThan20(db *gorm.DB) {
	var purchase_details []entity.PurchaseDetail
	subquery := db.Select("purchase_id").Group("purchase_id").Having("SUM(quantity) > ?", 20).Table("purchase_details")
	result := db.Debug().Where("purchase_id IN (?)", subquery).Find(&purchase_details)
	printResult(purchase_details, result)
}

func getTotalQty(db *gorm.DB) {
	var total int
	result := db.Debug().Model(&entity.Item{}).Select("SUM(quantity)").Group("item_category_id").Having("item_category_id = 1").Find(&total)
	printResult(total, result)
}

func createItem(db *gorm.DB) {
	item := entity.Item{
		Name:           "Vitamin C 1000 IU",
		Price:          10000,
		Quantity:       5,
		ItemCategoryId: 1,
	}
	result := db.Create(&item)
	printResult(item, result)
}

func incrementQtyAllItems(db *gorm.DB) {
	var item entity.Item
	result := db.Model(&item).Where("item_category_id = ?", 1).Update("quantity", gorm.Expr("quantity + ?", 5))
	printResult(item, result)
}

func softDelete(db *gorm.DB) {
	var item entity.Item
	result := db.Model(&item).Clauses(clause.Returning{}).Where("item_id = ?", 1).Delete(&item)
	printResult(item, result)
}

func getDetailsHeader(db *gorm.DB) {
	var header []entity.PurchaseHeader
	result := db.Preload("PurchaseDetails").First(&header)
	printResult(header, result)
}

func getItemCategory(db *gorm.DB) {
	var item []entity.ItemCategory
	result := db.Debug().Model(&entity.ItemCategory{}).Preload("Item").Find(&item)
	printResult(item, result)
}

func test(db *gorm.DB) {
	var pd []entity.PurchaseHeader
	result := db.Debug().Preload("PurchaseDetail").Find(&pd)
	printResult(pd, result)
}
