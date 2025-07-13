package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/Artur2912/pet1/models"
	"github.com/Artur2912/pet1/config"
	
)


type Handler struct{
	Config *config.Config
	Validate *validator.Validate
}

func NewHandler(cfg *config.Config)*Handler{
	return &Handler{
		Config: cfg,
		Validate: validator.New(),
	}
}
func (h *Handler) CreateUser (c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status bad request"})	
	}
	if err := h.Validate.Struct(user); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status bad request"})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failded hashed password"})
	}
	user.Password = string(hashedPassword)
	if err := h.Config.DB.Create(&user).Error; err != nil{
		return  c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not created"})
	}

	
	var responseusr = models.ResponseUser{
		Name: user.Name,
		Email: user.Email,
		Time: time.Now(),
	}
	return c.JSON(http.StatusCreated, responseusr)
}


func (h *Handler) GetAllUsers(c echo.Context)error{
	var users []models.User
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1{
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit
	if err := h.Config.DB.Limit(limit).Offset(offset).Find(&users).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Not find page"})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser (c echo.Context)error{
	id , err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Invalid JSON"} )
	}
	var user models.User
	if err := h.Config.DB.First(&user, id).Error; err != nil{
		 return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	response := models.ResponseUser{
		Name: user.Name,
		Email: user.Email,
		Time: time.Now(),
	}
	return c.JSON(http.StatusOK, response)
}



func (h *Handler) UpdateUser(c echo.Context)error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid json"})
	}
	var user models.User
	if err := h.Config.DB.First(&user, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	var updateUser models.UpdateUser
	if err := c.Bind(&updateUser); err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid json"})
	}
	if err := h.Validate.Struct(&updateUser); err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input data"})	
	}
	if updateUser.Name != nil{
		user.Name = *updateUser.Name
	}
	if updateUser.Email != nil{
		user.Email = *updateUser.Email
	}
	if updateUser.Password != nil{
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*updateUser.Password), bcrypt.DefaultCost)
		if err != nil{
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed hashed password"})
		}
		user.Password = string(hashedPassword)
	}
	if err := h.Config.DB.Save(&user).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed save user"})
	}

	return c.JSON(http.StatusOK, user)


}


func (h *Handler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := h.Config.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	return c.NoContent(http.StatusNoContent)
}


