# oidc-javascript is a HTML5 oidc app

Go is only used as a static file server

## Keycloak setup
keycloak is an openid connect compliant identity provider.

1. Log in to keycloak admin console.
1. Create a realm. Eg. "Beyond Broadcast LLP".
   Switch to the realm using the drop-down menu at top-left of screen.
1. Create a client. Eg "My HTML5 app".
   Ensure client Access Type is "public".
1. Create users.
   Create user credentials.

## Running the app

```
cd cmd/oidc-javascript
go run main.go
```

Point your browser to http://localhost:8080/   
