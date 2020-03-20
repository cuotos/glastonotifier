package songkick

type SongkickEventResponse struct {
	ResultsPage struct {
		Status string `json:"status"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
		Results struct {
			Event Event `json:"event"`
		} `json:"results"`
	}
}

type Event struct {
	Location struct {
		City string  `json:"city"`
		Lng  float64 `json:"lng"`
		Lat  float64 `json:"lat"`
	} `json:"location"`
	Popularity  float64 `json:"popularity"`
	URI         string  `json:"uri"`
	DisplayName string  `json:"displayName"`
	ID          int     `json:"id"`
	Type        string  `json:"type"`
	Start       struct {
		Time     string `json:"time"`
		Date     string `json:"date"`
		Datetime string `json:"datetime"`
	} `json:"start"`
	AgeRestriction string        `json:"ageRestriction"`
	Performance    []Performance `json:"performance"`
	Venue          Venue         `json:"venue"`
	Status         string        `json:"status"`
}

type Performance struct {
	Artist       Artist `json:"artist"`
	DisplayName  string `json:"displayName"`
	BillingIndex int    `json:"billingIndex"`
	ID           int    `json:"id"`
	Billing      string `json:"billing"`
}

type Artist struct {
	URI         string `json:"uri"`
	DisplayName string `json:"displayName"`
	ID          int    `json:"id"`
	Identifier  []struct {
		Href string `json:"href"`
		Mbid string `json:"mbid"`
	} `json:"identifier"`
}

type Venue struct {
	MetroArea struct {
		URI         string `json:"uri"`
		DisplayName string `json:"displayName"`
		Country     struct {
			DisplayName string `json:"displayName"`
		} `json:"country"`
		ID int `json:"id"`
	} `json:"metroArea"`
	City struct {
		URI         string `json:"uri"`
		DisplayName string `json:"displayName"`
		Country     struct {
			DisplayName string `json:"displayName"`
		} `json:"country"`
		ID int `json:"id"`
	} `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	URI         string  `json:"uri"`
	DisplayName string  `json:"displayName"`
	Street      string  `json:"street"`
	ID          int     `json:"id"`
	Website     string  `json:"website"`
	Phone       string  `json:"phone"`
	Capacity    int     `json:"capacity"`
	Description string  `json:"description"`
}
