package v1

import (
	"net/http"
	"refactoring/internal/domain"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) initUserRoutes(api chi.Router) chi.Router {
	return api.Route("/users", func(r chi.Router) {
		r.Get("/", h.searchUsers)
		r.Post("/", h.createUser)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.getUser)
			r.Patch("/", h.updateUser)
			r.Delete("/", h.deleteUser)
		})
	})
}

// @Summary Search Users
// @Tags user
// @ID search-users
// @Product json
// @Success	200		    {object}	dataResponse     "userList"
// @Failure	400,404		{object}	errResponse
// @Failure	500			{object}	errResponse
// @Failure	default		{object}	errResponse
// @Router		/api/v1/users [get]
func (h *Handler) searchUsers(w http.ResponseWriter, r *http.Request) {
	list, err := h.services.GetAll()
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	render.JSON(w, r, dataResponse{Data: list})
}

type createUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *createUserRequest) Bind(r *http.Request) error { return nil }

// @Summary Create user
// @Tags user
// @ID create-user
// @Accept json
// @Product json
// @Param input body createUserRequest true "request form"
// @Success	201		    {object}	idResponse     "id"
// @Failure	400,404		{object}	errResponse
// @Failure	500			{object}	errResponse
// @Failure	default		{object}	errResponse
// @Router		/api/v1/users [post]
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	request := createUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	u := domain.User{
		CreatedAt:   time.Now(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id, err := h.services.Create(u)
	if err != nil {
		render.Render(w, r, ErrServerError(err)) //nolint:errcheck
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, idResponse{ID: id})
}

// @Summary Get user
// @Tags user
// @ID get-user
// @Product json
// @Param	id	path integer true "ID of user"
// @Success	200		    {object}	domain.User    "user"
// @Failure	400,404		{object}	errResponse
// @Failure	500			{object}	errResponse
// @Failure	default		{object}	errResponse
// @Router		/api/v1/users/{id} [get]
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	user, err := h.services.GetByID(id)
	if err != nil {
		render.Render(w, r, ErrServerError(err)) //nolint:errcheck
		return
	}

	render.JSON(w, r, user)
}

type updateUserRequest struct {
	DisplayName *string `json:"display_name"`
}

func (c *updateUserRequest) Bind(r *http.Request) error { return nil }

// @Summary Update user
// @Tags user
// @ID update-user
// @Accept json
// @Param	id	path integer true "ID of user"
// @Param input body updateUserRequest true "update form"
// @Success	204 {string} string "NoContent"
// @Failure	400,404		{object}	errResponse
// @Failure	500			{object}	errResponse
// @Failure	default		{object}	errResponse
// @Router		/api/v1/users/{id} [patch]
func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	request := updateUserRequest{}

	if err = render.Bind(r, &request); err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	input := domain.UpdateUserInput{
		DisplayName: request.DisplayName,
	}

	if err = input.Validate(); err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	if err = h.services.Update(id, input); err != nil {
		render.Render(w, r, ErrServerError(err)) //nolint:errcheck
		return
	}

	render.Status(r, http.StatusNoContent)
}

// @Summary Delete user
// @Tags user
// @ID delete-user
// @Param	id	path integer true "ID of user"
// @Success	204 {string} string "NoContent"
// @Failure	400,404		{object}	errResponse
// @Failure	500			{object}	errResponse
// @Failure	default		{object}	errResponse
// @Router		/api/v1/users/{id} [delete]
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err)) //nolint:errcheck
		return
	}

	if err = h.services.Delete(id); err != nil {
		render.Render(w, r, ErrServerError(err)) //nolint:errcheck
		return
	}

	render.Status(r, http.StatusNoContent)
}
