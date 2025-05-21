package game

import (
	"encoding/json"
	"fmt"
	"github.com/RA341/multipacman/pkg"
	"github.com/rs/zerolog/log"
	"time"
)

type MessageData struct {
	msgInfo       map[string]interface{}
	world         *World
	playerSession *PlayerEntity
}

type MessageHandlerFunc func(data MessageData) map[string]interface{}

type MessageHandler struct {
	handler     MessageHandlerFunc
	messageName string
}

func (handler MessageHandler) WithMiddleware(middleware func(MessageHandlerFunc) MessageHandlerFunc) MessageHandler {
	handler.handler = middleware(handler.handler)
	return handler
}

func CheckGameOverMiddleware(existingFunc MessageHandlerFunc) MessageHandlerFunc {
	return func(data MessageData) map[string]interface{} {
		encodedMsg := existingFunc(data)
		reason := data.world.checkGameOver() // verify the state after message has been handled
		if reason != "" {
			data.world.GameOver(reason)
		}

		return encodedMsg
	}
}

func registerMessageHandlers(opts ...MessageHandler) map[string]MessageHandlerFunc {
	handlers := map[string]MessageHandlerFunc{}

	for _, opt := range opts {
		if opt.messageName == "" {
			log.Warn().Msg("Attempted to register a handler with an empty key. Skipping.")
			continue
		}
		if opt.handler == nil {
			log.Warn().Str("key", opt.messageName).Msg("Attempted to register a nil handler for key. Skipping.")
			continue
		}

		if _, exists := handlers[opt.messageName]; exists {
			log.Warn().Str("key", opt.messageName).Msg("Duplicate handler registration attempted. Overwriting.")
		}

		handlers[opt.messageName] = opt.handler
		//log.Debug().Str("key", opt.messageName).Msg("Registered handler")
	}

	return handlers
}

func MovMessage() MessageHandler {
	name := "pos"
	return MessageHandler{
		messageName: name,
		handler: func(data MessageData) map[string]interface{} {
			x, y, err := getCoordFromMessage(data.msgInfo)
			dir, ok := data.msgInfo["dir"]
			if err != nil || !ok {
				log.Error().Err(err).
					Any("data", data.msgInfo).
					Msg("Unable to find coordinates or direction from message")
				return nil
			}

			data.world.MovePlayer(data.playerSession, x, y)
			data.playerSession.Type = name
			data.playerSession.Dir = dir.(string)

			return data.playerSession.ToMap()
		},
	}
}

func EndGameMessage(reason string) MessageHandler {
	name := "gameover"
	return MessageHandler{
		messageName: name,
		handler: func(_ MessageData) map[string]interface{} {
			return map[string]interface{}{
				"type":   name,
				"reason": reason,
			}
		},
	}
}

func PelletMessage() MessageHandler {
	name := "pel"
	return MessageHandler{
		messageName: name,
		handler: func(data MessageData) map[string]interface{} {
			x, y, err := getCoordFromMessage(data.msgInfo)
			if err != nil {
				log.Error().Err(err).Any("msg", data.msgInfo).Msg("Unable to find coordinates or direction from message")
				return nil
			}
			data.world.EatPellet(x, y)

			return map[string]interface{}{
				"type": name,
				"x":    x,
				"y":    y,
			}
		},
	}
}

func PowerUpMessage(manager *Manager) MessageHandler {
	name := "pow"
	return MessageHandler{
		messageName: name,
		handler: func(data MessageData) map[string]interface{} {
			x, y, err := getCoordFromMessage(data.msgInfo)
			if err != nil {
				return nil
			}

			data.world.EatPowerUp(x, y)

			time.AfterFunc(powerUpTime, func() {
				mess := EndPowerUpMessage(data.world).handler(data)
				marshal, err := json.Marshal(mess)
				if err != nil {
					log.Warn().Err(err).Msg("Unable to marshal json")
					return
				}

				pkg.Elog(manager.broadcastAll(data.world, marshal))
			})

			return map[string]interface{}{
				"type": name,
				"x":    x,
				"y":    y,
			}
		},
	}
}

func KillPlayer() MessageHandler {
	name := "kill"
	return MessageHandler{
		messageName: name,
		handler: func(data MessageData) map[string]interface{} {
			ghostId, exists := data.msgInfo["id"]
			if !exists {
				log.Warn().Any("msg", data.msgInfo).Msg("no pellet ghostId found")
				return nil
			}

			if data.world.IsPoweredUp {
				data.world.GhostEatenAction(SpriteType(ghostId.(string)))
				return map[string]interface{}{
					"type":     name, // ghost eaten
					"spriteId": ghostId,
				}
			}

			data.world.GameOver("pacman was killed")
			return map[string]interface{}{
				"type":     name,
				"spriteId": Pacman, // pacman dead
			}
		},
	}
}

func EndPowerUpMessage(lobbyEntity *World) MessageHandler {
	mesName := "powend"
	return MessageHandler{
		messageName: mesName,
		handler: func(data MessageData) map[string]interface{} {
			lobbyEntity.IsPoweredUp = false
			return map[string]interface{}{
				"type": mesName,
			}
		},
	}
}

func getCoordFromMessage(msgInfo map[string]interface{}) (X float64, Y float64, err error) {
	x, existsX := msgInfo["x"]
	y, existsY := msgInfo["y"]
	if !existsY || !existsX {
		return 0, 0, fmt.Errorf("unable to find coordinates in message")
	}

	//if _, err := strconv.ParseFloat(x.(), 64); err != nil {
	//	return "", "", fmt.Errorf("invalid X coordinate in message")
	//}
	//if _, err := strconv.ParseFloat(y.(string), 64); err != nil {
	//	return "", "", fmt.Errorf("invalid Y coordinate in message")
	//}

	return x.(float64), y.(float64), nil
}
