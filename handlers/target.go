package handlers

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/vidurkhanal/infuse/config"
	"github.com/vidurkhanal/infuse/constants"
)

func TryTarget(c *fiber.Ctx, t config.Target, fn, method, path string) (fiber.Map, error) {
	switch t.Strategy.Mode {
	case constants.StrategyFallback:
		// fallback
		return Fallback(c, t, fn, method, path)
	case constants.StrategyLoadBalance:
		// load balance
		return LoadBalance(c, t, fn, method, path)
	default:
		// try single target
		return SingleTarget(c, t, fn, method, path)
	}
}

func Fallback(c *fiber.Ctx, t config.Target, fn, method, path string) (fiber.Map, error) {
	// Todo
	return nil, nil
}

func LoadBalance(c *fiber.Ctx, t config.Target, fn, method, path string) (fiber.Map, error) {
	var totalWeight int
	for _, target := range t.Targets {
		if target.Weight == 0 {
			target.Weight = 1
		}
		totalWeight += target.Weight
	}

	randomWeight := rand.Float32() * float32(totalWeight)
	for target_indx, target := range t.Targets {
		if float32(target.Weight) < randomWeight {
			currentPath := fmt.Sprintf("%s.targets[%d]", path, target_indx)
			return SingleTarget(c, target, fn, method, currentPath)
		}
		randomWeight -= float32(target.Weight)
	}
	return nil, nil
}

func SingleTarget(c *fiber.Ctx, t config.Target, fn, method, path string) (fiber.Map, error) {
	return nil, nil
}
