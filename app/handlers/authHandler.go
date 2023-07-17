package handlers

import (
	"auth/dto/request"
	"auth/service"
	"fmt"
	"github.com/azerpost/dashboard-lib/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	loginService service.AuthService
}

func (lh *AuthHandler) SignUp(c *gin.Context) {
	var request request.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errs.NewInternalServerError("Invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	} else {
		createdUser, err := lh.loginService.SignUp(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			fmt.Println(createdUser)
		}
	}

	// Hash the password
	//hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Failed to hash password",
	//	})
	//	return
	//}
	//// Create the user
	//user := models.User{Email: request.Email, Password: string(hash)}
	//if user.IsMailValid() {
	//	fmt.Println("is valid")
	//}
	//result := database.DB.Create(&user) // pass pointer of data to Create
	//if result.Error != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Failed to created user",
	//	})
	//	return
	//}
	//// Respond
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "Succesfuly created user",
	//})
}

func (lh *AuthHandler) Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errs.NewInternalServerError("Invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	} else {
		tokenString, err := lh.loginService.Login(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie(`Authorization`, tokenString, 3600*24*30, "", "", false, true)
			c.JSON(http.StatusOK, gin.H{
				"status": "okey",
			})
		}
	}
	// Look up requested user
	//var user models.User
	//initializers.DB.First(&user, "email = ?", body.Email)
	//if user.ID == 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Invalid email or password",
	//	})
	//	return
	//}
	//// compare sent in pass with saved user pass hash
	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "Invalid password",
	//	})
	//	return
	//}
	//
	//// Create a new token object, specifying signing method and the claims
	//// you would like it to contain.
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"sub": user.ID,
	//	"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	//})
	//
	//// Sign and get the complete encoded token as a string using the secret
	//tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "failed to create token",
	//	})
	//	return
	//}
	//
	//// send it back
	//c.JSON(http.StatusOK, gin.H{
	//	"token": tokenString,
	//})
	//c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie(`Authorization`, tokenString, 3600*24*30, "", "", false, true)
}

func NewAuthHandler(os service.AuthService) *AuthHandler {
	return &AuthHandler{os}
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
