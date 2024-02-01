package middlewares

import (
	"cmp"
	"errors"
	"slices"
	"strings"

	"github.com/vidurkhanal/infuse/constants"
)

var (
	STRATEGY_MODES []string = []string{
		constants.StrategyFallback,
		constants.StrategySingle,
		constants.StrategyLoadBalance,
		constants.StrategyConcurrent,
	}

	CACHE_MODE []string = []string{constants.CacheSimple, constants.CacheSemantic}
)

type HeaderConfigStrategy struct {
	Mode          string `json:"mode"`
	OnStatusCodes []int  `json:"on_status_codes"`
}

type HeaderConfigCache struct {
	Mode   string `json:"mode"`
	MaxAge int    `json:"max_age"`
}

type HeaderConfigRetry struct {
	Attempts      int   `json:"attempts"`
	OnStatusCodes []int `json:"on_status_codes"`
}

type HeaderConfig struct {
	Strategy      HeaderConfigStrategy `json:"strategy"`
	Provider      constants.Provider   `json:"provider"`
	APIKey        string               `json:"apiKey"`
	Cache         HeaderConfigCache    `json:"cache"`
	Retry         HeaderConfigRetry    `json:"retry"`
	Weight        int                  `json:"weight"`
	OnStatusCodes []int                `json:"onStatusCodes"`
	Targets       []HeaderConfig       `json:"targets"`
}

func (h HeaderConfig) Validate() error {
	if h.Provider != "" && h.APIKey != "" {
		// has provider name and api key
		if !contains(constants.Providers, h.Provider) {
			return errors.New(
				string(h.Provider) + " is an invalid provider or is not supported yet by " + constants.POWERED_BY)
		}
	} else if h.Strategy.Mode != "" && len(h.Targets) > 0 {
		// has strategy mode and targets
		if err := h.Strategy.Validate(); err != nil {
			return err
		}
		for _, target := range h.Targets {
			if err := target.Validate(); err != nil {
				return err
			}
		}
	} else if h.Cache.Mode != "" {
		// has cache mode and max age
		if err := h.Cache.Validate(); err != nil {
			return err
		}
	} else if h.Retry.Attempts > 0 {
		// has retry attempts
	} else {
		return errors.New("invalid config. must have provider & API key, strategy & targets, cache, or retry")
	}
	return nil
}

func (s HeaderConfigStrategy) Validate() error {
	if slices.Contains(STRATEGY_MODES, s.Mode) {
		return nil
	}
	return errors.New("invalid strategy mode")
}

func (c HeaderConfigCache) Validate() error {
	if slices.Contains(CACHE_MODE, c.Mode) {
		return nil
	}
	return errors.New("invalid cache mode")
}

func contains(s []constants.Provider, e constants.Provider) bool {
	for _, a := range s {
		if cmp.Compare(string(a), strings.ToLower(string(e))) == 0 {
			return true
		}
	}
	return false
}
