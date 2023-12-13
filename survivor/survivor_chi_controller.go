package survivor

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/somatom98/zssn/domain"
)

type SurvivorChiController struct {
	survivorService domain.SurvivorService
}

func NewChiController(survivorService domain.SurvivorService) *SurvivorChiController {
	return &SurvivorChiController{
		survivorService: survivorService,
	}
}

func (c *SurvivorChiController) GetRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/", c.getAllSurvivorsHandler)
	return router
}

func (c *SurvivorChiController) getAllSurvivorsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	survivors, err := c.survivorService.GetAllSurvivors(ctx)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	render.JSON(w, r, survivors)
}
