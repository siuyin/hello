package main

import (
	"context"
	"fmt"
	"log"

	oidc "github.com/coreos/go-oidc"
	"github.com/siuyin/dflt"
	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("keycloak with openid-connect")
	ctx := context.TODO()
	pURL := dflt.EnvString("PROVIDER_URL", "http://192.168.93.15:32674/auth/realms/home")
	provider, err := oidc.NewProvider(ctx, pURL)
	if err != nil {
		log.Fatalf("provider %s: %v", pURL, err)
	}
	fmt.Printf("provider endpoint: %#v\n", provider.Endpoint())

	oauth2Config := oauth2.Config{
		ClientID:    "golang",
		RedirectURL: "http://192.168.93.15:32798/",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	fmt.Printf("oauth2.config: %#v: the remainder is the same as for the openid connect example.", oauth2Config)
}
