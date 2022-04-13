package config

type (
	CheckoutParam struct {
		EventName      string `json:"event_name" bson:"event_name" binding:"required"`
		UserName       string `json:"username" bson:"username" binding:"required"`
		TicketCheckout int32  `json:"ticket_checkout" bson:"ticket_checkout" binding:"required,numeric,min=1"`
	}

	PayParam struct {
		PriceUniqueValue float64 `json:"price_unik" bson:"price_unik" binding:"required"`
	}
)
