package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"wasgood/app"
	"wasgood/models"
)

var (
	oauthVKConfig = &oauth2.Config{
		RedirectURL:  app.RootURL + "/auth/vk/callback",
		ClientID:     "5589322",
		ClientSecret: "6ygBqDKSP1ERYoSMS1I1",
		Scopes:       []string{"email"},
		Endpoint:     vk.Endpoint,
	}
)

const VKapiVersion = "5.53"

func VKLogin(c *gin.Context) {
	url := oauthVKConfig.AuthCodeURL(url.QueryEscape(c.DefaultQuery("redirect", "/")), oauth2.SetAuthURLParam("display", "popup"), oauth2.SetAuthURLParam("v", VKapiVersion))
	c.Redirect(http.StatusMovedPermanently, url)
}

func VKCallback(c *gin.Context) {
	redirect_url, err := url.QueryUnescape(c.Query("state"))
	if err != nil {
		redirect_url = "/"
	}

	code := c.Query("code")
	token, err := oauthVKConfig.Exchange(oauth2.NoContext, code)
	check(err)

	user := vkupdateUserInfo(token)
	Login(user, c)
	c.Redirect(http.StatusTemporaryRedirect, app.RootURL+redirect_url)
}

func vkupdateUserInfo(token *oauth2.Token) *models.User {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	providerID := int(token.Extra("user_id").(float64))
	var email string
	if v := token.Extra("email"); v != nil { // user can deny access to email
		email = v.(string)
	}

	resp, err := client.Get(
		fmt.Sprintf("https://api.vk.com/method/users.get?user_ids=%d&fields=photo_50&name_case=Nom&v=%s&access_token=%s",
			providerID,
			VKapiVersion,
			token.AccessToken))
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	data := raw["response"].([]interface{})[0].(map[string]interface{})

	user := &models.User{
		Provider:   models.VK,
		ProviderID: providerID,
		Name:       data["first_name"].(string),
		LastName:   data["last_name"].(string),
		Photo:      data["photo_50"].(string),
		Email:      email,
	}

	user.Save()

	return user
}
