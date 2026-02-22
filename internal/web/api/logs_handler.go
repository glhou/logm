package api

import (
	"encoding/json"
	"net/http"

	"github.com/glhou/logm/internal/core"
	"github.com/glhou/logm/internal/db"
	"github.com/glhou/logm/internal/web/jsonw"
	"go.uber.org/zap"
)

type LogRequest struct {
	Service string `json:"service"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

type Log struct {
	Service string `json:"service"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

type LogResponse struct {
	Log    Log    `json:"log"`
	Status string `json:"status"`
}

type LogsResponse struct {
	Logs []Log `json:"logs"`
}

func newLogHandler(app *core.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := jsonw.WriteError(w, http.StatusBadRequest, "Wrong method")
			if err != nil {
				app.Logger.Error("Error", zap.Error(err))
			}
			return
		}
		var req LogRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			err := jsonw.WriteError(w, http.StatusBadRequest, "invalid JSON")
			if err != nil {
				app.Logger.Error("Error", zap.Error(err))
			}
			return
		}

		q := db.New(app.Db)
		ctx := r.Context()
		log, err := q.CreateLog(ctx, db.CreateLogParams{
			Service: req.Service,
			Level:   req.Level,
			Message: req.Message,
		})
		if err != nil {
			err := jsonw.WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			if err != nil {
				app.Logger.Error("Error", zap.Error(err))
			}
			return
		}

		err = jsonw.WriteJSON(w, http.StatusOK, LogResponse{
			Log: Log{
				Service: log.Service,
				Level:   log.Level,
				Message: log.Message,
			},
			Status: "ok",
		})
		if err != nil {
			app.Logger.Error("Error", zap.Error(err))
		}
	}
}

func listLogHandler(app *core.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			err := jsonw.WriteError(w, 400, "Wrong method")
			if err != nil {
				app.Logger.Error("Error", zap.Error(err))
			}
			return
		}
		q := db.New(app.Db)
		ctx := r.Context()
		logs, err := q.ListLogs(ctx, 100)
		if err != nil {
			err := jsonw.WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
			if err != nil {
				app.Logger.Error("Error", zap.Error(err))
			}
			return
		}
		var ress []Log
		for _, log := range logs {
			ress = append(ress, Log{
				log.Service,
				log.Message,
				log.Level,
			})
		}
		app.Logger.Info("Test", zap.String("logs", ress[0].Message))
		err = jsonw.WriteJSON(w, http.StatusOK, LogsResponse{ress})
		if err != nil {
			app.Logger.Error("Error", zap.Error(err))
		}
	}
}
