package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"github.com/ethereum/go-ethereum/log"

	"github.com/eniac-x-labs/manta-relayer/manager/types"
)

type Registry struct {
	signService types.SignService
}

func NewRegistry(signService types.SignService) *Registry {
	return &Registry{
		signService: signService,
	}
}

func (registry *Registry) SignStateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.SignStateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}
		if len(request.StateRoot.String()) == 0 {
			c.JSON(http.StatusBadRequest, errors.New("state root must not be nil"))
			return
		}
		var signature []byte
		var err error

		signature, err = registry.signService.SignStateBatch(request)

		if err != nil {
			c.String(http.StatusInternalServerError, "failed to sign state")
			log.Error("failed to sign state", "error", err)
			return
		}
		if _, err = c.Writer.Write(signature); err != nil {
			log.Error("failed to write signature to response writer", "error", err)
		}
	}
}

func (registry *Registry) PrometheusHandler() gin.HandlerFunc {
	h := promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer, promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{MaxRequestsInFlight: 3},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
