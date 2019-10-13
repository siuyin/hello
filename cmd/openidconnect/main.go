package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/siuyin/dflt"
	"golang.org/x/oauth2"
)

var (
	providerURI  = dflt.EnvString("IDP_URL", "https://siuyin.ice.ibmcloud.com/oidc/endpoint/default")
	clientID     = dflt.EnvString("CLIENT_ID", "de1603d7-7473-4217-8efd-2f5314573ca7")
	clientSecret = dflt.EnvString("CLIENT_SECRET", "OmlelKVSBu")
)

func main() {
	fmt.Println("OpenID Connect example")

	// root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, path is: %q\n", html.EscapeString(r.URL.Path))
	})

	// login handler
	state := "state should be returned unmodified" // Don't do this in production.
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, providerURI)
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})
	if err != nil {
		log.Fatalf("provider: %v", err)
	}
	var claims struct {
		ScopesSupported []string `json:"scopes_supported"`
		ClaimsSupported []string `json:"claims_supported"`
	}
	if err := provider.Claims(&claims); err != nil {
		log.Println(err)
	}
	fmt.Println(claims.ScopesSupported)
	fmt.Println(claims.ClaimsSupported)
	fmt.Println(provider.Endpoint())
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "https://rasp.beyondbroadcast.com/auth/oidc/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Redirecting to ID provider: %s\n", providerURI)
		http.Redirect(w, r, config.AuthCodeURL(state,
			oauth2.SetAuthURLParam("code_challenge", "abc123ifjsalfjldfsIOfjldsjflasjlfsdjlfdslkjlsdfjlfslkfs"),
		),
			http.StatusFound)
	})

	//http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	//	log.Printf("Redirecting to ID provider: %s\n", providerURI)
	//	http.Redirect(w, r, providerURI, http.StatusFound)
	//})

	// oidc callback
	http.HandleFunc("/auth/oidc/callback", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Processing Authorization response")
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"),
			oauth2.SetAuthURLParam("code_verifier", "abc123ifjsalfjldfsIOfjldsjflasjlfsdjlfdslkjlsdfjlfslkfs"),
		)
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "could not get raw ID token", http.StatusInternalServerError)
			return
		}
		// Parse and verify ID Token payload.
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "could not verify ID token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var claims struct {
			GroupIDs []string `json:"groupIds"`
		}
		if err := idToken.Claims(&claims); err != nil {
			http.Error(w, "Failed extract claims: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(claims)
		fmt.Printf("idToken: %s\n", idToken)

		userInfo, err := provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
		if err != nil {
			http.Error(w, "Failed to get userinfo: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("userInfo: %s\n", userInfo)

		resp := struct {
			OAuth2Token *oauth2.Token
			UserInfo    *oidc.UserInfo
			IDToken     *oidc.IDToken
		}{oauth2Token, userInfo, idToken}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	// static content handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// start serving
	log.Fatal(http.ListenAndServeTLS(":8080", "/h/certbot/rasp.beyondbroadcast.com/fullchain.pem",
		"/h/certbot/rasp.beyondbroadcast.com/privkey.pem", nil))
}
