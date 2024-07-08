// GetPetsByType godoc
// @Summary Get pets by type
// @Description Get a list of pets by type
// @Tags pets
// @Produce json
// @Param type query string true "Pet Type"
// @Success 200 {array} model.Pet
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /pets/type [get]

package api
