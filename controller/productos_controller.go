
package controller
import (
  "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)
type IDProd struct {
	Id []string              `json:"results"`
}

type Foto struct {
	Url string                 `json:"url"`
}

type Product struct {
	Id    string               `json:"id"`
	Title string               `json:"title"`
	Price float64              `json:"price"`
	Aq int     `json:"available_quantity"`
	Pictures []PictureMeli	   `json:"pictures"`
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

// PREGUNTAS SIN RESPONDER
type QuestionMeli struct {
	Date_created string   `json:"date_created"`
	Text string           `json:"text"`
	Status string         `json:"status"`
}

type QuestionsMeli struct {
	Questions []QuestionMeli  `json:"questions"`
}

// ESTRUCTURA PARA ENVIAR AL FRONT

type Item struct {
	Title string
	Quantity int
	Price float64
	FirstPicture string
}

type Sold_Item struct {
	Title string
	Sold_Quantity int
	Unit_Price float64
	Subtotal float64
}

type Sale_Order struct {
	Sold_Items [] Sold_Item
	Sale_date string
	Total  float64
	Total_Delivery float64
}

type Unanswered_Question struct {
	Question_date string
	Title string
	Question_text string
}

type Dashboard struct {
	Items [] Item
	Sales_Orders [] Sale_Order
	Unanswered_Questions [] Unanswered_Question
}

func GetItem(c*gin.Context){
  token := c.Query("token")
	userid := c.Query("userid")
	var url string = "https://api.mercadolibre.com/users/" + userid + "/items/search?access_token=" + token
	req, err := http.Get(url)
  if err!= nil {
    fmt.Error("Error",err.Error())
    return
  }
  defer resp.Body.Close()
  data,err := ioutil.ReadAll(resp.Body)

  var IdP IDProd
  json.Unmarshal(dataItemsId, &IdP)
  if err != nil {
		c.JSON(500, err.Error())
		return
	}
	ch := make(chan Item)
	var response[] Item

	for i := 0; i < len(res.Results); i++ {
		go itemCollector(token, res.Results[i], ch, c)
	}
	for i := 0; i < len(res.Results); i++ {
		response = append(response, <-ch)
	}
	c.JSON(200, response)
  }

  func itemCollector(token, userid string, ch chan Item, c *gin.Context ){
	var url string = "https://api.mercadolibre.com/items?ids="+ userid +"&attributes=id,price,category_id,title,pictures&access_token=" + token
	req, erro := http.Get(url)
	if req.StatusCode != 200 || erro != nil{
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	defer req.Body.Close()
	data, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	var res[] Items
	fmt.Println(string(data))
	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	var resp Item
	resp.Category = res[0].Category
	resp.Id = res[0].Body.Id
	resp.Picture = res[0].Body.Pictures[0]["url"]
	resp.Price = res[0].Body.Price
	resp.Title = res[0].Body.Title
	ch <- resp
}
