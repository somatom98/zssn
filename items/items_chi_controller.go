package items

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/somatom98/zssn/domain"
)

type ItemsChiController struct {
	itemsRepository domain.ItemsRepository
}

func NewChiController(itemsRepository domain.ItemsRepository) *ItemsChiController {
	return &ItemsChiController{
		itemsRepository: itemsRepository,
	}
}

func (c *ItemsChiController) GetRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/", c.getAllItemsHandler)
	router.Get("/{name}", c.getItemByNameHandler)
	return router
}

func (c *ItemsChiController) getAllItemsHandler(w http.ResponseWriter, r *http.Request) {
}

func (c *ItemsChiController) getItemByNameHandler(w http.ResponseWriter, r *http.Request) {
}
