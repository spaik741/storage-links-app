package web

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"log"
	"storage-links-app/models"
	"storage-links-app/repository"
)

func GetLinks(ctx fiber.Ctx) error {
	var chatId = ctx.Params("chatId")
	linksEntity, err := repository.FindByChatId(chatId)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	links := linksEntity.Links
	if links == nil || len(links) == 0 {
		return ctx.JSON(models.LinksResponse{Links: make([]string, 0)})
	}
	linksResponse := models.LinksResponse{
		Links: links,
	}
	return ctx.JSON(linksResponse)
}

func Save(ctx fiber.Ctx) error {
	body := ctx.Body()
	var request models.LinksRequest
	if err := json.Unmarshal(body, &request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	linksEntity, err := repository.FindByChatId(request.ChatId)
	if err == nil {
		linksEntity.Links = append(linksEntity.Links, request.Link)
		err = repository.Update(linksEntity)
	} else {
		log.Printf(err.Error())
		err = repository.Save(mapToEntity(request))
	}
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func Clear(ctx fiber.Ctx) error {
	chatId := ctx.Query("chatId")
	if len(chatId) == 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("chatId is empty")
	}
	entity, err := repository.FindByChatId(chatId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	entity.Links = entity.Links[:0]
	err = repository.Update(entity)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func mapToEntity(r models.LinksRequest) models.LinksEntity {
	return models.LinksEntity{
		ChatId: r.ChatId,
		Links:  []string{r.Link},
	}
}
