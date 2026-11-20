package modifiers

import(
	"financialExchange/enum"
)

// принимает sellerQuantity типа float64
func ModifierForOrder(sellerQuantity float64) func(*enum.Order) *enum.Order{
	// замыкание которое «запоминает» sellerQuantity
	return func(order *enum.Order) *enum.Order{
		// Увеличиваем поле Filled на переданное количество
		order.Filled += sellerQuantity

		// Filled больше или равно требуемому Quantity — ставим статус Filled
		if order.Filled >= order.Quantity{
			order.Status = enum.OrderStatusFilled
		} else{
			order.Status = enum.OrderStatusPartiallyFilled
		}
		return order
	}
}