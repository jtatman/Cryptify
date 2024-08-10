package trading

import (
	"fmt"
	"sync"

	"github.com/jtatman/Cryptify/cryptify_app/internal/order"
)

// TradingEngine handles the execution of trading orders
type TradingEngine struct {
	orderService *order.OrderService
	trades       map[int]order.Order
	tradesMutex  sync.Mutex
}

// NewTradingEngine creates a new instance of TradingEngine
func NewTradingEngine(orderService *order.OrderService) *TradingEngine {
	return &TradingEngine{
		orderService: orderService,
		trades:       make(map[int]order.Order),
	}
}

// ProcessOrder executes a trading order
func (t *TradingEngine) ProcessOrder(orderID int) {
	t.tradesMutex.Lock()
	defer t.tradesMutex.Unlock()

	// Fetch the order from the order service
	o, exists := t.orderService.GetOrder(orderID)
	if !exists {
		fmt.Printf("Order %d not found\n", orderID)
		return
	}

	// Execute the order (simulate trading)
	fmt.Printf("Order executed: %v\n", o)

	// Record the trade
	t.trades[orderID] = o
}
