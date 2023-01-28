package handlers

import (
	"gorm-raw/db"
	"gorm-raw/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddNewUser(ctx *gin.Context) {
	d := db.GetDB()
	var user models.User

	ctx.BindJSON(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	err = d.QueryRow(
		`
			INSERT INTO "users" ("email","password") VALUES ($1, $2)
			RETURNING "id", "email", "password";
		`,
		user.Email,
		hashedPassword,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"data":    user,
			"message": "error",
			"error":   err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusCreated,
			"data":    user,
			"message": "Created",
		})
	}
}

func FindByEmail(userEmail string) (*models.User, error) {
	d := db.GetDB()
	var user models.User

	err := d.QueryRow(
		`
			SELECT users.id, users.email, users.password FROM users WHERE users.email = $1
		`,
		userEmail,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func FindByEmailGorm(userEmail string) (*models.User, error) {
	d := db.GetGormDB()
	var user models.User

	err := d.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(ctx *gin.Context) {
	var user models.User

	ctx.BindJSON(&user)

	userByEmail, err := FindByEmail(string(user.Email))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "invalid email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "wrong password",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": userByEmail,
	})
}

func LoginGorm(ctx *gin.Context) {
	var user models.User

	ctx.BindJSON(&user)

	userByEmail, err := FindByEmailGorm(string(user.Email))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "invalid email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(user.Password))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "wrong password",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": userByEmail,
	})
}
