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

const (
	ibmProvider     = "https://siuyin.ice.ibmcloud.com/oidc/endpoint/default"
	ibmClientID     = "de1603d7-7473-4217-8efd-2f5314573ca7"
	ibmClientSecret = "OmlelKVSBu"
)

var (
	providerURI  = dflt.EnvString("IDP_URL", ibmProvider)
	clientID     = dflt.EnvString("CLIENT_ID", ibmClientID)
	clientSecret = dflt.EnvString("CLIENT_SECRET", ibmClientSecret)
	redirectURL  = dflt.EnvString("REDIRECT_URL", "https://rasp.beyondbroadcast.com/auth/oidc/callback")
)

func main() {
	fmt.Println("OpenID Connect example")

	// root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, path is: %q\n", html.EscapeString(r.URL.Path))
	})

	// OIDC setup
	provider := oidcProvider(providerURI)
	pm := providerMetadata(provider)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "email", "profile", "roles"},
	}

	// login handler
	state := "state should be returned unmodified" // Don't do this in production.
	http.Handle("/login", newLoginHandler(&config, state))

	// logout handler
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r,
			//fmt.Sprintf("%s?redirect_uri=%s", pm.EndSessionEndpoint, url.QueryEscape("http://localhost:8082/")), // no-longer supported: see https://www.keycloak.org/docs/latest/upgrading/index.html#openid-connect-logout
			//fmt.Sprintf("%s?post_logout_redirect_uri=%s&client_id=goclient", pm.EndSessionEndpoint, url.QueryEscape("http://localhost:8082/")),
			pm.EndSessionEndpoint,
			http.StatusFound)
	})

	// oidc callback
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})
	http.Handle("/auth/oidc/callback", newCallbackHandler(&config, state, provider, verifier))

	// static content handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// start serving
	log.Println("Starting http server.")
	log.Fatal(http.ListenAndServe(":8082", nil))
	//log.Fatal(http.ListenAndServeTLS(":8080", "/h/certbot/rasp.beyondbroadcast.com/fullchain.pem",
	//	"/h/certbot/rasp.beyondbroadcast.com/privkey.pem", nil))
}
func oidcProvider(url string) *oidc.Provider {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, url)
	if err != nil {
		log.Fatalf("provider: %v", err)
	}
	return provider
}

type providerClaims struct {
	ScopesSupported    []string `json:"scopes_supported"`
	ClaimsSupported    []string `json:"claims_supported"`
	JWKSURI            string   `json:"jwks_uri"`
	EndSessionEndpoint string   `json:"end_session_endpoint"`
}

func providerMetadata(provider *oidc.Provider) *providerClaims {
	// for details, see: https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
	claims := providerClaims{}
	if err := provider.Claims(&claims); err != nil {
		log.Println(err)
	}
	fmt.Printf("scopes supported: %v\n", claims.ScopesSupported)
	fmt.Printf("claims supported: %v\n", claims.ClaimsSupported)
	fmt.Printf("jwks URI: %v\n", claims.JWKSURI)
	fmt.Printf("end session endpoint: %v\n", claims.EndSessionEndpoint)
	fmt.Printf("provider endpoint: %v\n", provider.Endpoint())
	return &claims
}

type loginHandler struct {
	config *oauth2.Config
	state  string
}

func newLoginHandler(config *oauth2.Config, state string) *loginHandler {
	h := new(loginHandler)
	h.config = config
	h.state = state
	return h
}
func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Redirecting to ID provider: %s\n", providerURI)
	log.Printf("URL: %s", h.config.AuthCodeURL(h.state,
		oauth2.SetAuthURLParam("code_challenge", "abc123ifjsalfjldfsIOfjldsjflasjlfsdjlfdslkjlsdfjlfslkfs"),
	))
	http.Redirect(w, r, h.config.AuthCodeURL(h.state,
		oauth2.SetAuthURLParam("code_challenge", "abc123ifjsalfjldfsIOfjldsjflasjlfsdjlfdslkjlsdfjlfslkfs"),
	),
		http.StatusFound)
}

//http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
//	log.Printf("Redirecting to ID provider: %s\n", providerURI)
//	http.Redirect(w, r, providerURI, http.StatusFound)
//})

type callbackHandler struct {
	config   *oauth2.Config
	state    string
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

func newCallbackHandler(config *oauth2.Config, state string, provider *oidc.Provider, verifier *oidc.IDTokenVerifier) *callbackHandler {
	h := callbackHandler{config, state, provider, verifier}
	return &h
}
func (h *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Processing Authorization response")
	if r.URL.Query().Get("state") != h.state {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := h.config.Exchange(context.Background(), r.URL.Query().Get("code"),
		oauth2.SetAuthURLParam("code_verifier", "abc123ifjsalfjldfsIOfjldsjflasjlfsdjlfdslkjlsdfjlfslkfs"),
	)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("code token expiry: %v\n", oauth2Token.Expiry)
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "could not get raw ID token", http.StatusInternalServerError)
		return
	}
	// Parse and verify ID Token payload.
	idToken, err := h.verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		http.Error(w, "could not verify ID token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("idToken: %s\n", idToken)

	userInfo, err := h.provider.UserInfo(context.Background(), oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		http.Error(w, "Failed to get userinfo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("userInfo: %v\n", userInfo)

	resp := struct {
		OAuth2Token   *oauth2.Token
		UserInfo      *oidc.UserInfo
		IDToken       *oidc.IDToken
		IDTokenClaims *json.RawMessage
	}{oauth2Token, userInfo, idToken, new(json.RawMessage)}
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, "Failed extract claims: "+err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
