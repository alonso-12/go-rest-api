package handlers

import (
	"matryer/internal"
	"matryer/internal/db"

	"matryer/internal/types"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type IUserHandler interface {
	HandleGetUsers() http.HandlerFunc
	HandleCreateUser() http.HandlerFunc
}

type UserHandler struct {
	logger    *logrus.Logger
	userStore db.UserStore
}

func NewUserHandler(logger *logrus.Logger, userStore db.UserStore) *UserHandler {
	return &UserHandler{
		logger:    logger,
		userStore: userStore,
	}
}

type userResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
	Age       int32  `json:"age"`
}

func (u *UserHandler) HandleGetUsers() http.HandlerFunc {
	type response struct {
		Status    string         `json:"status"`
		HTTPCode  int            `json:"http_code"`
		Datetime  string         `json:"datetime"`
		Timestamp int64          `json:"timestamp"`
		User      []userResponse `json:"user"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.userStore.GetUsers()
		if err != nil {
			u.logger.Printf("error get users: %v", err)
			internal.RenderResponse(w, err)
			return
		}

		var userResp []userResponse
		for _, user := range users {
			userResp = append(userResp, userResponse{
				ID:        user.ID,
				Name:      user.Name,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Phone:     user.Phone,
				Age:       user.Age,
			})
		}

		internal.Respond(w, response{
			Status:    "success",
			HTTPCode:  http.StatusOK,
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
			User:      userResp,
		}, http.StatusOK)
	}
}

func (u *UserHandler) HandleCreateUser() http.HandlerFunc {
	type (
		request struct {
			Name      string `json:"name"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			Age       int32  `json:"age"`
		}

		response struct {
			Status    string       `json:"status"`
			HTTPCode  int          `json:"http_code"`
			Datetime  string       `json:"datetime"`
			Timestamp int64        `json:"timestamp"`
			User      userResponse `json:"user"`
		}
	)
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req request
			err error
		)

		if err = internal.Decode(r, &req); err != nil {
			u.logger.Errorf("invalid user request: %v", err)
			internal.RenderResponse(w, err)
			return
		}

		userParams := &types.User{
			Name:      types.Name(req.Name),
			FirstName: types.FirstName(req.FirstName),
			LastName:  types.LastName(req.LastName),
			Email:     types.Email(req.Email),
			Phone:     types.Phone(req.Phone),
			Age:       types.Age(req.Age),
		}

		if err = userParams.Validate(); err != nil {
			u.logger.Errorf("invalid user params: %v", err)
			internal.RenderResponse(w, err)
			return
		}

		user, err := u.userStore.CreateUser(r.Context(), userParams)
		if err != nil {
			u.logger.Errorf("error created user: %v", err)
			internal.RenderResponse(w, err)
			return
		}

		internal.Respond(w, response{
			Status:    "success",
			HTTPCode:  http.StatusCreated,
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
			User: userResponse{
				ID:        user.ID,
				Name:      user.Name,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Phone:     user.Phone,
				Age:       user.Age,
			},
		}, http.StatusCreated)
	}
}
