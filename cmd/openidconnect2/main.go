package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/siuyin/dflt"
	"golang.org/x/oauth2"
)

var (
	clientID     = dflt.EnvString("CLIENT_ID", "goclient")
	clientSecret = dflt.EnvString("CLIENT_SECRET", "your secret here")
	nav          = `<a href="/login" style="display: block">Login (will request credentials if not already logged in)</a>
<a id="logout" href="http://localhost:8081/realms/junk/protocol/openid-connect/logout" style="display: block">Logout</a>
<script>
	const logout=document.getElementById("logout");
	const idtoken=sessionStorage.getItem("idtoken");
	const logoutURI="http://localhost:8081/realms/junk/protocol/openid-connect/logout?post_logout_redirect_uri="
	+ encodeURIComponent("http://localhost:8080/")
	+ "&id_token_hint="+idtoken;
	logout.setAttribute("href",logoutURI);
</script>`
)

func main() {
	fmt.Println("oidc with keycloak")
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, dflt.EnvString("PROVIDER_URL", "http://localhost:8081/realms/myrealm"))
	if err != nil {
		log.Fatal(err)
	}

	oidcConfig := &oidc.Config{ClientID: clientID}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  dflt.EnvString("REDIRECT_URL", "http://localhost:8080/auth/callback"),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<h1>OpenID Connect</h1>`+nav)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		state, err := randString(16)
		if err != nil {
			http.Error(w, "state: Internal error with random string generator", http.StatusInternalServerError)
			return
		}
		nonce, err := randString(16)
		if err != nil {
			http.Error(w, "nonce: Internal error with random string generator", http.StatusInternalServerError)
			return
		}

		setCallbackCookie(w, r, "state", state)
		setCallbackCookie(w, r, "nonce", nonce)

		http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
	})

	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		state, err := r.Cookie("state")
		if err != nil {
			http.Error(w, "state not found", http.StatusBadRequest)
			return
		}

		if r.URL.Query().Get("state") != state.Value {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}

		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		nonce, err := r.Cookie("nonce")
		if err != nil {
			http.Error(w, "nonce not found", http.StatusBadRequest)
			return
		}
		if idToken.Nonce != nonce.Value {
			http.Error(w, "nonce did not match", http.StatusBadRequest)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, `<div><a href="/">Home</a></div>`)
		fmt.Fprintf(w, nav)
		fmt.Fprintf(w, "<pre>")
		w.Write(data)
		fmt.Fprintf(w, "</pre>")
		fmt.Fprintf(w, `<script>
sessionStorage.setItem("idtoken","%s")
</script>`, rawIDToken)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", dflt.EnvIntMust("PORT", 8080)), nil))
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
