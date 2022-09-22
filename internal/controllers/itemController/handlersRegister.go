package itemController

import (
	"net/http"
)

func HandlersRegister(c *controller) {
	http.HandleFunc("/item/create", c.Create)
	http.HandleFunc("/item/getList", c.GetList)
	http.HandleFunc("/item/remove", c.Delete)
	http.HandleFunc("/item/update", c.Update)
}
