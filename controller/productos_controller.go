
package controller
import (
  "net/http"
  "github.com/gin-gonic/gin"
  productosservice
)

func AddProdcuto(){
  resp, err := http.Get("https://api.mercadolibre.com/users/651268893/items/search?access_token=APP_USR-2760149476611182-110500-2af3caa0dd8bce845f0493a720b4d10d-651268893")
	if err != nil {
		fmt.Errorf("Error",err.Error())
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	bodyString := string(data)
	fmt.Println(bodyString)

	var itemsIds ItemsIds
	json.Unmarshal(data, &itemsIds)
	fmt.Printf("%+v\n", itemsIds)
}
