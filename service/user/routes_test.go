package user

import	 (
	"fmt"
	"bytes"
	"net/http/httptest"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"testing"
	"github.com/ndraoo/restapi-go/types"
)
func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}	
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is valid", func(t *testing.T) {
		payload :=  types.RegisterUserPay{ 
			FirstName: "John",
			LastName:  "Doe",
			Email:    "sasd" ,
			Password: "password",
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("failed to marshal payload: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))	
		
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest { 
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {

}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
