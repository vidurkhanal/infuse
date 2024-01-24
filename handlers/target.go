package handlers

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/vidurkhanal/infuse/config"
	"github.com/vidurkhanal/infuse/constants"
	"github.com/vidurkhanal/infuse/utils"
)

func TryTarget(c *fiber.Ctx, t config.Target, fn, method, path string) (*fasthttp.Response, error) {
	switch t.Strategy.Mode {
	case constants.StrategyFallback:
		// fallback
		return Fallback(c, t, fn, method, path)
	case constants.StrategyLoadBalance:
		// load balance
		return LoadBalance(c, t, fn, method, path)
	case constants.StrategySingle:
		// try single target
		return SingleTarget(c, t, fn, method, path)
	default:
		return TryPost(c, t, fn, method, path)
	}
}

func Fallback(c *fiber.Ctx, t config.Target, fn, method, path string) (*fasthttp.Response, error) {
	// Todo
	var wg sync.WaitGroup
	var res *fasthttp.Response
	var err error
	for target_indx, target := range t.Targets {
		wg.Add(1)
		currentPath := fmt.Sprintf("%s.targets[%d]", path, target_indx)
		go func(t config.Target) {
			defer wg.Done()
			res, err = TryTarget(c, t, fn, method, currentPath)
		}(target)
	}
	wg.Wait()
	return res, err
}

func LoadBalance(c *fiber.Ctx, t config.Target, fn, method, path string) (*fasthttp.Response, error) {
	var totalWeight int
	for _, target := range t.Targets {
		if target.Weight == 0 {
			target.Weight = 1
		}
		totalWeight += target.Weight
	}

	randomWeight := utils.GetRngInstance().Gen.Intn(totalWeight + 1)
	for target_indx, target := range t.Targets {
		if target.Weight < randomWeight {
			currentPath := fmt.Sprintf("%s.targets[%d]", path, target_indx)
			return TryTarget(c, target, fn, method, currentPath)
		}
		randomWeight -= target.Weight
	}

	return &fasthttp.Response{}, errors.New("no target found")
}

func SingleTarget(c *fiber.Ctx, t config.Target, fn, method, path string) (*fasthttp.Response, error) {
	return TryTarget(c, t.Targets[0], fn, method, path+".targets[0]")
}

func TryPost(c *fiber.Ctx, t config.Target, fn, method, path string) (*fasthttp.Response, error) {
	return TryTarget(c, t, fn, fiber.MethodPost, path)
}
