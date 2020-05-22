package models

type Booking struct {
	Id        int           `json:"id" binding:"required"`
	UserId    string        `json:"user_id" binding:"required"`
	Detail    BookingDetail `json:"booking_detail" binding:"required"`
	TotalFare float32       `json:"total_fare" binding:"required"`
	Status    string        `json:"status" binding:"required"`
	Notes     string        `json:"notes"`
}

type BookingDetail struct {
	Item      string             `json:"item"`
	Fare      BookingFare        `json:"fare"`
	Passenger []BookingPassenger `json:"passenger"`
}

type BookingFare struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type BookingPassenger struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
	Contact  string `json:"contact"`
}
