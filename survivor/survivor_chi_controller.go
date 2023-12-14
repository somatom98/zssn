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
	router.Put("/", c.addSurvivorHandler)
	router.Get("/{sid}", c.getSurvivorHandler)
	router.Patch("/{sid}", c.reportSurvivorStatusHandler)
	router.Get("/{sid}/location", c.updateSurvivorLocationHandler)
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

func (c *SurvivorChiController) updateSurvivorLocationHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sid := chi.URLParam(r, "sid")
	var location domain.Location

	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	err = c.survivorService.UpdateSurvivorLocation(ctx, sid, location)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}
}

func (c *SurvivorChiController) reportSurvivorStatusHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sid := chi.URLParam(r, "sid")
	var survivorStatusReport domain.SurvivorStatusReport

	err := json.NewDecoder(r.Body).Decode(&survivorStatusReport)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}

	err = c.survivorService.ReportSurvivorStatus(ctx, sid, survivorStatusReport)
	if err != nil {
		rErr := domain.NewError(err)
		render.Render(w, r, rErr)
		return
	}
}
