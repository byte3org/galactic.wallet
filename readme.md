### galactic.wallet

Each citizen of the known universe has a wallet associated to their Galactic ID. This service is responsible for handling every function related to that wallet.

## How it works

# endpoints

* GET /api/v1/wallet/   - returns all the wallets (not implemented)
* POST /api/v1/wallet   - creates a new wallet (if not already exists) 
* GET /api/v1/wallet/user/balance - get Galatic Solaris balance of the user's wallet (not implemented) 
* Get /api/v1/wallet/transactions - get transaction history of the user's wallet (not implemented)
* POST /api/v1/wallet/user/deposit - make a Galactic Solaris deposit to user's wallet (not implemented)
* POST /api/v1/wallet/user/withdraw - withdraw Galactic Solaris from users wallet

Every HTTP request that is being sent to this service goes through number of middlewares. These include

* `ExtractUserId` - call /userinfo endpoint on OpenId auth server and extract the user id, which can then be used to get wallet balance and transaction history
* `SignatureVerify` - combined with above middleware to be used with `/withdraw` and `/deposit` endpoints to verify the amount that has been sent as a parameter to enable secure payments

### Run

go build cmd/payment/main.go
