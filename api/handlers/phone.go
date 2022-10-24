package handlers

import (
	restservice "task/rest_service"
	"task/rest_service/api/http"
	"task/rest_service/storage"

	"github.com/saidamir98/udevs_pkg/util"

	fiber "github.com/gofiber/fiber/v2"
)

// CreatePhone godoc
// @ID create_phone
// @Router /phone [POST]
// @Summary Create Phone
// @Description Create Phone
// @Tags phone
// @Accept json
// @Produce json
// @Param phone body restservice.PhoneNumber true "PhoneNumber"
// @Success 201 {object} http.Response{data=string} "Phone data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePhone(c *fiber.Ctx) error {
	var phone restservice.Phone
	err := c.BodyParser(&phone)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	id, err := h.store.Phone().Create(c.Context(), phone)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, err.Error())
	}

	return h.handleResponse(c, http.Created, id)
}

// GetPhoneList godoc
// @ID get_phone_list
// @Router /phone [GET]
// @Summary Get Phone List
// @Description  Get Phone List
// @Tags phone
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=restservice.PhoneList} "PhoneList"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPhoneList(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	list, count, err := h.store.Phone().GetList(c.Context(), limit, offset)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, err.Error())
	}

	return h.handleResponse(c, http.OK, restservice.PhoneList{Phones: list, Count: count})
}

// GetPhoneByID godoc
// @ID get_phone_by_id
// @Router /phone/{phone_id} [GET]
// @Summary Get Phone By ID
// @Description Get Phone By ID
// @Tags phone
// @Accept json
// @Produce json
// @Param phone_id path string true "phone_id"
// @Success 200 {object} http.Response{data=restservice.Phone} "PhoneBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPhoneByID(c *fiber.Ctx) error {
	phoneID := c.Params("phone_id")

	if !util.IsValidUUID(phoneID) {
		return h.handleResponse(c, http.InvalidArgument, "phone_id is an invalid uuid")
	}

	phone, err := h.store.Phone().GetByPK(c.Context(), phoneID)
	if err == storage.ErrorNotFound {
		return h.handleResponse(c, http.NotFound, err.Error())
	}

	if err != nil {
		return h.handleResponse(c, http.InternalServerError, err.Error())
	}

	return h.handleResponse(c, http.OK, phone)
}

// UpdatePhone godoc
// @ID update_phone
// @Router /phone [PUT]
// @Summary Update Phone
// @Description Update Phone
// @Tags phone
// @Accept json
// @Produce json
// @Param phone body restservice.Phone true "Phone"
// @Success 200 {object} http.Response{data=string} "Phone data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePhone(c *fiber.Ctx) error {
	var phone restservice.Phone
	err := c.BodyParser(&phone)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	rowsAffected, err := h.store.Phone().Update(c.Context(), phone)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return h.handleResponse(c, http.NotFound, "phone not found")
	}

	return h.handleResponse(c, http.OK, "Successfully updated")
}

// DeletePhone godoc
// @ID delete_phone
// @Router /phone/{phone_id} [DELETE]
// @Summary Delete Phone
// @Description Get Phone
// @Tags phone
// @Accept json
// @Produce json
// @Param phone_id path string true "phone_id"
// @Success 200 {object} http.Response{data=string} "Phone data"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePhone(c *fiber.Ctx) error {
	phoneID := c.Params("phone_id")

	if !util.IsValidUUID(phoneID) {
		return h.handleResponse(c, http.InvalidArgument, "phone_id is an invalid uuid")
	}

	rowsAffected, err := h.store.Phone().Delete(c.Context(), phoneID)
	if err != nil {
		return h.handleResponse(c, http.InternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return h.handleResponse(c, http.NotFound, "phone not found")
	}

	return h.handleResponse(c, http.OK, "Successfully deleted")
}
