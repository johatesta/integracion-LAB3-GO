package productos
type ItemsIdMeli struct {
	Id []string              `json:"results"`
}

type PictureMeli struct {
	Url string                 `json:"url"`
}

type Productos struct {

    Nombre    string `json:"nombre"`
    cantidad   int64 `json:"cantidad"`
    Cantidad_Disponible string `json:"Cantidad Disponible"`
    Precio  string `json:"precio"`
    Fotos []PictureMeli `json:"Fotos"`
}

// ITEMS VENDIDOS
type SingleItemMeli struct {
	Title string                `json:"title"`
}

type Order_ItemsMeli struct {
	SingleItem SingleItemMeli    `json:"item"`
	Quantity int                 `json:"quantity"`
	Unit_price float64           `json:"unit_price"`
	Full_Unit_Price float64      `json:"full_unit_price"`
}

type ResultMeli struct {
	Order_Items []Order_ItemsMeli `json:"order_items"`
	Total_amount float64          `json:"total_amount"`
	Paid_amount float64           `json:"paid_amount"`
	Date_closed string            `json:"date_closed"`
}

type SoldItemMeli struct {
	Result []ResultMeli            `json:"results"`
}
