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

//
//type OrderBook struct {
//	Asks []*DepthOrder `json:"asks"`
//	Bids []*DepthOrder `json:"bids"`
//}
