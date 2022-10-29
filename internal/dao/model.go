package dao

type TableName string

const (
	TableNameUser     TableName = "user"
	TableNameBankCard TableName = "bank_card"
	TableNameBank     TableName = "bank"

	TableNameGoods             TableName = "goods"
	TableNameUserGoods         TableName = "user_goods"
	TableNameShop              TableName = "shop"
	TableNameHomeSlider        TableName = "home_slider"
	TableNameActivity          TableName = "activity"
	TableNameHomeItem          TableName = "home_item"
	TableNameDetailItem        TableName = "detail_item"
	TableNameUserCollection    TableName = "user_collection"
	TableNameUserGoodsCheckin  TableName = "user_goods_checkin"
	TableNameOrder             TableName = "order"
	TableNameAmountChange      TableName = "user_amount_change"
	TableNameAmountChangeMonth TableName = "user_amount_change_month"
)
