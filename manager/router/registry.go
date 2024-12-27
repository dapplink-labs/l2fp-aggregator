package router

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/l2fp-aggregator/manager/types"
)

type Registry struct {
	signService types.SignService
}

func NewRegistry(signService types.SignService) *Registry {
	return &Registry{
		signService: signService,
	}
}

func (registry *Registry) SignMsgHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request types.SignMsgRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("invalid request body"))
			return
		}
		if len(request.TxHash) == 0 || request.BlockNumber == nil || request.TxType != "" {
			c.JSON(http.StatusBadRequest, errors.New("tx_hash, block_number and tx_type must not be nil"))
			return
		}
		var result *types.SignResult
		var err error

		result, err = registry.signService.SignMsgBatch(request)

		if err != nil {
			c.String(http.StatusInternalServerError, "failed to sign msg")
			log.Error("failed to sign msg", "error", err)
			return
		}
		if _, err = c.Writer.Write(result.Signature.Serialize()); err != nil {
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
