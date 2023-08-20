package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/byte3/galactic.wallet/config"
	"github.com/byte3/galactic.wallet/helpers"
	"github.com/google/uuid"
)

func ExtractOidcToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		oidcToken := authHeader[1]
		ctx := context.WithValue(r.Context(), "oidc", oidcToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SignatureVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			msg := "Content-Type Header is not in the required format (json)"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}

		// request body
		// { amount: 20, signature: "adadadada" }
		jsonReq, err := helpers.ParseJsonBody(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// verify the signature
		keystring := "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAIdUgg1Ii8IpJzcG+i4oPGZKzYxVJ3HpWpQvtqsa6ByFYsh36u3oW8VcjssH9BJTtytB7NHsBS26GFLtq+2Q3+0ml1V4kGe7rswFbMymG+UO2mH6q0pI1Y9OAiMCcuV54ruZ+rQ6NCrSWntWoxGAmLXWIWxMNkFuVrkxG8pQs7hhAgMBAAECgYAxfyU5+oizcR3YsIKckzbUKxPW+eY9cZ4hiPoExwiSNe8VZ1bTwSKfouTPOY95jIj4F1qoxOx39xKici9p6o0bwOdSaI6ivRCujMtRtyKelxJhNYg/gyZtTi9VWORkkOZkm5Et9Q44eeZmoUvd5NiSF5xYXQCZrstHfPqCzCK+CQJBAPNehog4Y+rjs1c5vT4N5mN3LBfizDD+N/OcXirdNDLYG7U2aQRhKZcftse9m5wvflKbqL4uK45vYO3MU1ITfE8CQQCOWosO1bRx5Zu62JLZ6+LBVdYLWgdWeC1qtbQCDR9KCwOqsOzhv86gzvq6/97DK3nxB949jtwlGZSfplwd5uRPAkA3lFPXEkHHZ/8SIY6VeGkwOAwq3FHTsosmqIRc962vGumhBe8P3/y2lbiRfzzle3c2+HOeEz9BRTB2vl4c9XRDAkAj65L1PlWW++Is5qM/m/cO4/Lr0F7ToeWL7KescNU5YMgfFR/g4v5ns3KvJwt14g2WFW8tx1OjhO3szxSlcvKfAkA80L1uuVGhGvUAdP0aYsH0temKXeTAtz+p673disGys2Q+82kNgpVIsIvH5w5COwqWi87hi/rP9bUCYkXH+kUM"

		key, err := helpers.BytesToPrivateKey([]byte(keystring))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		decryptedBody, err := helpers.DecryptWithPrivateKey([]byte(jsonReq["signature"].(string)), key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Print(decryptedBody)

		// compare amount and decrypted body amount
		if jsonReq["amount"].(int) != decryptedAmount {
			http.Error(w, "A modification to the request is detected", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "amount", decryptedAmount)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func ExtractUserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		oidcToken := authHeader[1]

		client := http.Client{}
		req, err := http.NewRequest("GET", config.GetConfig().AuthService, nil)
		if err != nil {
			http.Error(w, "Could not reach galactic auth service", http.StatusInternalServerError)
			return
		}
		req.Header = http.Header{
			"Authorization": {"Bearer " + oidcToken},
		}
		res, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()
		jsonRes, err := helpers.ParseJsonBody(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userId := jsonRes["userId"].(uuid.UUID)

		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
