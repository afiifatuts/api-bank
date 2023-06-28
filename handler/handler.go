package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/afiifatuts/bankmnc/helper"
	"github.com/afiifatuts/bankmnc/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		isExist := helper.IsRegistered(username)

		if isExist {
			req := helper.GetUser(username)

			match := helper.CheckPasswordHash(password, req.Password)

			if !match {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "Wrong password/username",
				})
				ctx.Abort()
				return
			}

			req.IsLogin = true
			token, err := helper.GenerateToken(username, time.Minute*10) // Set token expiration time
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Failed to generate token",
				})
				ctx.Abort()
				return
			}

			data := gin.H{
				"username": req.Username,
				"is_login": req.IsLogin,
				"token":    token,
			}

			ctx.JSON(http.StatusOK, gin.H{
				"success":   true,
				"message":   "Login Successful!",
				"user_info": data,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid username",
			})
		}
	}
}

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authPayload := ctx.MustGet("payload").(*helper.Payload)

		if authPayload.IsLogin {
			authPayload.ExpiredAt = time.Now()
			fmt.Println(authPayload.ExpiredAt)
			ctx.JSON(http.StatusOK, gin.H{
				"success":   true,
				"user_info": authPayload.Username,
				"message":   "Logout Successful!",
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Not logged in",
			})
		}
	}
}

func Payment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		amount := ctx.PostForm("amount")
		toAccount := ctx.PostForm("to_account")
		merchant := ctx.PostForm("merchant")
		amountInt, _ := strconv.Atoi(amount)
		authPayload := ctx.MustGet("payload").(*helper.Payload)

		// Check token validity
		err := authPayload.Valid()
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Token has expired",
			})
			return
		}

		//Check user validity
		isValidUser := helper.GetUser(toAccount)
		if isValidUser == (model.User{}) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid User",
			})
			return
		}

		// Create a new Transaction object
		payment := model.Transaction{
			ID:          uuid.New().String(),
			FromAccount: authPayload.Username,
			ToAccount:   toAccount,
			Merchant:    merchant,
			Amount:      float64(amountInt),
			Timestamp:   time.Now().String(),
		}

		data := gin.H{
			"id_transaction": payment.ID,
			"amount":         payment.Amount,
			"from_account":   payment.FromAccount,
			"to_account":     payment.ToAccount,
			"merchant":       payment.Merchant,
			"timestamp":      payment.Timestamp,
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Payment Success",
			"transaction": data,
		})
	}
}
