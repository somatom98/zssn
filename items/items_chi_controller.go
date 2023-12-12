package items

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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
	ctx := r.Context()

	items, err := c.itemsRepository.GetAllItems(ctx)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	render.JSON(w, r, items)
}

func (c *ItemsChiController) getItemByNameHandler(w http.ResponseWriter, r *http.Request) {
}
