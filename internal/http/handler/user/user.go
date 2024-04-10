package user

import (
	"log/slog"
	"net/http"

	"tp_back/internal/exception"
	"tp_back/internal/http/handler/user/request"
	"tp_back/internal/http/handler/user/response"
	"tp_back/internal/lib/rq"
	"tp_back/internal/lib/rs"
	"tp_back/internal/lib/vl"
	"tp_back/internal/service/user"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Handler struct {
	log     *slog.Logger
	service user.Service
}

func New(log *slog.Logger, service user.Service) *Handler {
	return &Handler{log: log, service: service}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.handler.user.Register"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req request.Register

		log.Info("decoding request body")
		if err := rq.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", slog.Any("error", err))
			w.WriteHeader(exception.Code(err))
			render.JSON(w, r, rs.Err(err.Error()))
			return
		}
		log.Info("decoded request body", slog.Any("request", req))

		log.Info("validating request body")
		if vlErrs := vl.Validate(req); vlErrs != nil {
			log.Error("validation error", slog.Any("validation errors", vlErrs))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, rs.ValidationErr(vlErrs))
			return
		}
		log.Info("validate request body", slog.Any("request", req))

		log.Info("running service.Register")
		u, err := h.service.Register(user.RegisterDTO{Email: req.Email, Password: req.Password})
		if err != nil {
			log.Error("failed to run service.Register")
			w.WriteHeader(exception.Code(err))
			render.JSON(w, r, rs.Err(err.Error()))
			return
		}
		log.Info("ran service.Register", slog.Any("user", u))

		render.JSON(w, r, rs.OK(response.Register{
			Uuid:      u.Uuid,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}))
	}
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.handler.user.Login"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req request.Login

		log.Info("decoding request body")
		if err := rq.DecodeJSON(r.Body, &req); err != nil {
			log.Error("failed to decode request body", slog.Any("error", err))
			w.WriteHeader(exception.Code(err))
			render.JSON(w, r, rs.Err(err.Error()))
			return
		}
		log.Info("decoded request body", slog.Any("request", req))

		log.Info("validating request body")
		if vlErrs := vl.Validate(req); vlErrs != nil {
			log.Error("validation error", slog.Any("validation error", vlErrs))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, rs.ValidationErr(vlErrs))
			return
		}
		log.Info("validate request body", slog.Any("request", req))

		log.Info("running service.Login")
		token, err := h.service.Login(user.LoginDTO{Email: req.Email, Password: req.Password})
		if err != nil {
			log.Error("failed to run service.Login")
			w.WriteHeader(exception.Code(err))
			render.JSON(w, r, rs.Err(err.Error()))
			return
		}
		log.Info("ran service.Login")

		render.JSON(w, r, rs.OK(response.Login{
			Token: token,
		}))
	}
}
