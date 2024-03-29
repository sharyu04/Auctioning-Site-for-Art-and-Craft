package api

// "bytes"
// "fmt"
// "net/http"
// "net/http/httptest"
// "testing"

// "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/app/user/mocks"
// "github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
// "github.com/stretchr/testify/mock"

// func TestCreateUserHandler(t *testing.T) {
// 	userSvc := mocks.NewService(t)
// 	CreateUserHandler := createUserHandler(userSvc)

// 	tests := []struct {
// 		name               string
// 		input              string
// 		setup              func(mock *mocks.Service)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name: "Success for create user",
// 			input: `{
// 						"firstName" : "user1",
// 						"lastName" : "Marwadi",
// 						"email": "sharayu6@gmail.com",
// 						"password": "Root"
//                     }`,
// 			setup: func(mockSvc *mocks.Service) {
// 				mockSvc.On("CreateUser", mock.Anything, mock.Anything).Return(dto.UserSignupResponse{}, nil).Once()
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 		// {
// 		//     name:               "Fail Invalid json",
// 		//     input:              "",
// 		//     setup:              func(mockSvc *mocks.Service) {},
// 		//     expectedStatusCode: http.StatusBadRequest,
// 		// },

// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(userSvc)

// 			req, err := http.NewRequest("POST", "/user/signup", bytes.NewBuffer([]byte(test.input)))
// 			if err != nil {
// 				t.Fatal(err)
// 				return
// 			}

// 			rr := httptest.NewRecorder()
// 			handler := http.HandlerFunc(CreateUserHandler)
// 			handler.ServeHTTP(rr, req)

// 			fmt.Println("Error")

// 			if rr.Result().StatusCode != test.expectedStatusCode {
// 				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
// 			}
// 		})
// 	}
// }

// func TestLoginHandler(t *testing.T) {
// 	userSvc := mocks.NewService(t)
// 	LoginHandler := loginHandler(userSvc)

// 	tests := []struct {
// 		name               string
// 		input              string
// 		setup              func(mock *mocks.Service)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name: "Success for user login",
// 			input: `{
//                         "email": "abc@gmail.com",
//                         "password": "root"
//                     }`,
// 			setup: func(mockSvc *mocks.Service) {
// 				mockSvc.On("LoginUser", mock.Anything).Return("token", nil).Once()
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 		// {
// 		//     name:               "Fail Invalid json",
// 		//     input:              "",
// 		//     setup:              func(mockSvc *mocks.Service) {},
// 		//     expectedStatusCode: http.StatusBadRequest,
// 		// },
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(userSvc)

// 			req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(test.input)))
// 			if err != nil {
// 				t.Fatal(err)
// 				return
// 			}

// 			rr := httptest.NewRecorder()
// 			handler := http.HandlerFunc(LoginHandler)
// 			handler.ServeHTTP(rr, req)

// 			fmt.Println("Error")

// 			if rr.Result().StatusCode != test.expectedStatusCode {
// 				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
// 			}
// 		})
// 	}
// }

// func TestGetAllUsersHandler(t *testing.T) {
// 	userSvc := mocks.NewService(t)
// 	GetAllUsersHandler := GetAllUsersHandler(userSvc)

// 	tests := []struct {
// 		name               string
// 		start              int
// 		count              int
// 		role               string
// 		setup              func(mock *mocks.Service)
// 		expectedStatusCode int
// 	}{
// 		{
// 			name:  "Success for get all users",
// 			start: 0,
// 			count: 4,
// 			setup: func(mockSvc *mocks.Service) {
// 				mockSvc.On("GetAllUsers", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(userSvc)

// 			req, err := http.NewRequest("GET", fmt.Sprintf("/users?start=%d&count=%d&role=%s", test.start, test.count, test.role), bytes.NewBuffer([]byte("")))
// 			if err != nil {
// 				t.Fatal(err)
// 				return
// 			}

// 			rr := httptest.NewRecorder()
// 			handler := http.HandlerFunc(GetAllUsersHandler)
// 			handler.ServeHTTP(rr, req)

// 			fmt.Println("Error")

// 			if rr.Result().StatusCode != test.expectedStatusCode {
// 				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
// 			}
// 		})
// 	}
// }
