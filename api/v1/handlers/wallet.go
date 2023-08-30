package handlers

mport (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/byte3/galactic.wallet/internal/database"
	"github.com/byte3/galactic.wallet/internal/models"
	"github.com/google/uuid"
)

type WalletHandler struct{}

func (wh WalletHandler) GetAllWallets(w http.ResponseWriter, r *http.Request) {
	return
}

func (wh WalletHandler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type Header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wallet models.WalletModel
	err = json.Unmarshal(body, &wallet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if wallet exist for the user
	_, err = database.SelectWalletByUserId(wallet.ID)
	if err == nil {
		http.Error(w, "User already has a wallet", http.StatusConflict)
	}

	rows, err := database.CreateWallet(&wallet)
	if err != nil || rows != 1 {
		http.Error(w, "Failed to create entry", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (wh WalletHandler) GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}

func (wh WalletHandler) GetWalletTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (wh WalletHandler) DepositToWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (wh WalletHandler) WithdrawFromWallet(w http.ResponseWriter, r *http.Request) {

	userId := r.Context().Value("user_id").(uuid.UUID)
	amount, err := strconv.Atoi(r.Context().Value("amount").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wallet, err := database.SelectWalletByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if wallet.Balance > amount {
		wallet.Balance -= amount
		// insert into database
		rows, err := database.UpdateWalletBalance(&wallet)
		if err != nil || rows != 1 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payment := models.PaymentModel{
			Amount:    amount,
			IsSuccess: true,
			WalletID:  wallet.ID,
			Wallet:    wallet,
		}
		rows, err = database.CreatePayment(&payment)
		if err != nil || rows != 1 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		msg := map[string]interface{}{
			"id":  payment.ID,
			"msg": "payment is successful",
		}

		j, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(j))
	} else {
		payment := models.PaymentModel{
			Amount:    amount,
			IsSuccess: false,
			WalletID:  wallet.ID,
			Wallet:    wallet,
		}
		rows, err := database.CreatePayment(&payment)
		if err != nil || rows != 1 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "Transaction failed. Not Enough Cash", 402)
		return
	}
}
