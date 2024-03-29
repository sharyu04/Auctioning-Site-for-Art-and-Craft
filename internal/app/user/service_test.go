package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/pkg/dto"
	"github.com/sharyu04/Auctioning-Site-for-Art-and-Craft/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		input           dto.CreateUserRequest
		role            string
		uuid_id         string
		setup           func(userMock *mocks.UserStorer)
		isErrorExpected bool
	}{
		{
			name: "Success for create user",
			input: dto.CreateUserRequest{
				FirstName: "Sharyu",
				LastName:  "Marwadi",
				Email:     "sharyu@gmail.com",
				Password:  "Root$123",
			},
			role: "",
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("CheckEmailExists", mock.Anything).Return(nil).Once()
				userMock.On("GetRoleID", mock.Anything).Return(uuid.Nil, nil).Once()
				userMock.On("CreateUser", mock.Anything).Return(dto.UserSignupResponse{}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo)

			// test service
			_, err := service.CreateUser(test.input, test.role)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

// func TestLoginUser(t *testing.T) {
// 	userRepo := mocks.NewUserStorer(t)
// 	service := NewService(userRepo)

// 	tests := []struct {
// 		name            string
// 		input           dto.LoginRequest
// 		setup           func(userMock *mocks.UserStorer)
// 		isErrorExpected bool
// 	}{
// 		{
// 			name: "Success for User Login",
// 			input: dto.LoginRequest{
// 				Email:    "bidder1@gmail.com",
// 				Password: "root",
// 			},
// 			setup: func(userMock *mocks.UserStorer) {
// 				// hashedPassword, _ := helpers.HashPassword(password)
// 				Id, _ := uuid.Parse("010d9c50-646c-48af-8cb1-2db613102c4d")
// 				userMock.On("GetUserByEmail", mock.Anything).Return(dto.User{
// 					Id:        Id,
// 					FirstName: "bidder1",
// 					LastName:  "Marwadi",
// 					Email:     "bidder1@gmail.com",
// 					Password:  "$2a$14$9QkXuA/eLtP/B0tzK1/nBeH5FwJPZik2QPA5H8ys9sxglJkDKAZDW",
// 					Role:      "user",
// 				}, nil).Once()
// 			},
// 			isErrorExpected: false,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.setup(userRepo)

// 			// test service
// 			_, err := service.LoginUser(test.input)

// 			if (err != nil) != test.isErrorExpected {
// 				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
// 			}
// 		})
// 	}
// }

func TestGetAllUsers(t *testing.T) {
	userRepo := mocks.NewUserStorer(t)
	service := NewService(userRepo)

	tests := []struct {
		name            string
		start           int
		count           int
		role            string
		setup           func(userMock *mocks.UserStorer)
		isErrorExpected bool
	}{
		{
			name:  "Success for User Login",
			start: 0,
			count: 4,
			role:  "user",
			setup: func(userMock *mocks.UserStorer) {
				Id, _ := uuid.Parse("010d9c50-646c-48af-8cb1-2db613102c4d")
				userMock.On("GetAllUsersByRole", mock.Anything, mock.Anything, mock.Anything).Return([]dto.GetAllUserResponse{
					{
						ID:        Id,
						FirstName: "bidder1",
						LastName:  "Marwadi",
						Email:     "bidder1@gmail.com",
						RoleID:    "user",
					},
				}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(userRepo)

			// test service
			_, err := service.GetAllUsers(test.start, test.count, test.role)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
