package survivor

import (
	"encoding/json"
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
	router.Get("/{sid}", c.getSurvivorHandler)
	router.Put("/", c.addSurvivorHandler)
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

func (c *SurvivorChiController) getSurvivorHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sid := chi.URLParam(r, "sid")

	survivor, err := c.survivorService.GetSurvivor(ctx, sid)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	render.JSON(w, r, survivor)
}

func (c *SurvivorChiController) addSurvivorHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var survivor domain.Survivor

	err := json.NewDecoder(r.Body).Decode(&survivor)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	sid, err := c.survivorService.AddSurvivor(ctx, survivor)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	render.JSON(w, r, sid)
}
