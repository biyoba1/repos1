package statTracker

type DepthOrder struct {
	Price   float64
	BaseQty float64
}

type OrderBook struct {
	Exchange string
	Pair     string
	Asks     []*DepthOrder
	Bids     []*DepthOrder
}

