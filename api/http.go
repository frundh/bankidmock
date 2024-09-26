package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

type order struct {
	CreatedAt int64 `json:"createdAt"`
	IsValid   bool  `json:"isValid"`
}

type Controller struct {
	db *bbolt.DB
}

func NewController(db *bbolt.DB) *Controller {
	return &Controller{
		db: db,
	}
}

func (ctrl *Controller) Auth(w http.ResponseWriter, r *http.Request) {
	log.Println("auth")
	if r.Method != http.MethodPost {
		log.Println("method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("unsupported media type")
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("error reading request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var authReqDTO AuthRequestDTO
	if err := json.Unmarshal(data, &authReqDTO); err != nil {
		log.Println("error unmarshalling request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderRef := uuid.NewString()

	if err := ctrl.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("orders"))
		t := time.Now().Add(time.Duration(3 * time.Second)).Unix()
		orderData, err := json.Marshal(order{CreatedAt: t, IsValid: false})
		if err != nil {
			return err
		}

		return bucket.Put([]byte(orderRef), orderData)
	}); err != nil {
		log.Println("error updating database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(AuthResponseDTO{
		OrderRef:       orderRef,
		AutoStartToken: uuid.NewString(),
		QRStartToken:   uuid.NewString(),
		QRStartSecret:  uuid.NewString(),
	}); err != nil {
		log.Println("error encoding response: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ctrl *Controller) Collect(w http.ResponseWriter, r *http.Request) {
	log.Println("collect")
	if r.Method != http.MethodPost {
		log.Println("method not allowed")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("unsupported media type")
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	data, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("error reading request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var collectReqDTO CollectRequestDTO
	if err := json.Unmarshal(data, &collectReqDTO); err != nil {
		log.Println("error unmarshalling request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := ctrl.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("orders"))
		data := bucket.Get([]byte(collectReqDTO.OrderRef))
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			return nil
		}
		var order order
		if err := json.Unmarshal(data, &order); err != nil {
			return err
		}

		isPassed := time.Now().Compare(time.Unix(order.CreatedAt, 0).Add(time.Duration(6*time.Second))) > 0
		if !isPassed {
			if err := json.NewEncoder(w).Encode(CollectResponseDTO{
				OrderRef: collectReqDTO.OrderRef,
				Status:   CollectStatusPending,
				HintCode: HintCodeStarted,
			}); err != nil {
				return err
			}
			w.Header().Set("Content-Type", "application/json")
			return nil
		}

		order.IsValid = true
		orderBytes, err := json.Marshal(order)
		if err != nil {
			return err
		}
		bucket.Put([]byte(collectReqDTO.OrderRef), orderBytes)

		if err := json.NewEncoder(w).Encode(CollectResponseDTO{
			OrderRef: collectReqDTO.OrderRef,
			Status:   CollectStatusComplete,
			HintCode: "",
			CompletionData: CompletionData{
				User: User{
					PersonalNumber: "199001011234",
					Name:           "John Doe",
					GivenName:      "John",
					Surname:        "Doe",
				},
				Device: Device{
					IPAddress:        "127.0.0.1",
					UniqueHardwareID: "1234567890",
				},
				BankIDIssueDate: time.Now(),
				Signature:       "signature",
				OCSPResponse:    "ocspResponse",
			},
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		log.Println("error updating database: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
