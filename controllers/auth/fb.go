package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"wasgood/app"
	"wasgood/models"
)

var (
	oauthFBConfig = &oauth2.Config{
		RedirectURL:  app.RootURL + "/auth/fb/callback",
		ClientID:     "251347055265766",
		ClientSecret: "c3eb4496d6f2dfd9745e9756bf02c3a5",
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
)

const FBapiVersion = "2.7"

func FBLogin(c *gin.Context) {
	url := oauthFBConfig.AuthCodeURL(url.QueryEscape(c.DefaultQuery("redirect", "/")))
	c.Redirect(http.StatusMovedPermanently, url)
}

func FBCallback(c *gin.Context) {
	redirect_url, err := url.QueryUnescape(c.Query("state"))
	if err != nil {
		redirect_url = "/"
	}

	code := c.Query("code")
	token, err := oauthFBConfig.Exchange(oauth2.NoContext, code)
	check(err)

	user := fbUpdateUserInfo(token)
	Login(user, c)
	c.Redirect(http.StatusTemporaryRedirect, app.RootURL+redirect_url)
}

func fbUpdateUserInfo(token *oauth2.Token) *models.User {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(
		fmt.Sprintf("https://graph.facebook.com/v%s/me?fields=id,first_name,last_name,picture,email&access_token=%s",
			FBapiVersion,
			token.AccessToken))
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	providerID, _ := strconv.Atoi(data["id"].(string))

	var email string
	if v, ok := data["email"]; ok {
		email = v.(string)
	}

	user := &models.User{
		Provider:   models.FB,
		ProviderID: providerID,
		Name:       data["first_name"].(string),
		LastName:   data["last_name"].(string),
		Photo:      data["picture"].(map[string]interface{})["data"].(map[string]interface{})["url"].(string),
		Email:      email,
	}

	user.Save()

	return user
}
