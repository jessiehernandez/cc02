package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/testing/example/handler"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type userPreferenceServiceStub struct {
	ReturnError error
}

func (u *userPreferenceServiceStub) Save(ctx context.Context, userID string, preferences interface{}) (err error) {
	return u.ReturnError
}

var _ = Describe("User Preference Handler", func() {
	var response *httptest.ResponseRecorder

	ctx := context.Background()

	BeforeEach(func() {
		response = httptest.NewRecorder()
	})

	When("an unsupported HTTP method is used", func() {
		It("returns an error", func() {
			userPreferenceHandler := handler.NewUserPreference(&userPreferenceServiceStub{})
			request, err := http.NewRequestWithContext(ctx, "PATCH", "/preference", strings.NewReader(""))

			Expect(err).To(BeNil())

			userPreferenceHandler.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})

	When("the payload is in an invalid format", func() {
		It("returns an error", func() {
			userPreferenceHandler := handler.NewUserPreference(&userPreferenceServiceStub{})
			request, err := http.NewRequestWithContext(ctx, "POST", "/preference", strings.NewReader("whatever"))

			Expect(err).To(BeNil())

			userPreferenceHandler.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusBadRequest))
		})
	})

	When("the preferences are saved successfully", func() {
		It("returns an OK status", func() {
			userPreferenceHandler := handler.NewUserPreference(&userPreferenceServiceStub{})
			request, err := http.NewRequestWithContext(ctx, "POST", "/preference", strings.NewReader(`
				{
					"userID": "user1",
					"preferences": {
						"first": "value"
					}
				}
			`))

			Expect(err).To(BeNil())

			userPreferenceHandler.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})

	When("the user ID is not passed", func() {
		It("returns an error", func() {
			userPreferenceHandler := handler.NewUserPreference(&userPreferenceServiceStub{})
			request, err := http.NewRequestWithContext(ctx, "POST", "/preference", strings.NewReader("{}"))

			Expect(err).To(BeNil())

			userPreferenceHandler.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusBadRequest))
		})
	})

	When("there is an error saving the user preferences", func() {
		It("returns an error", func() {
			userPreferenceHandler := handler.NewUserPreference(&userPreferenceServiceStub{
				ReturnError: errors.New("mock error"),
			})
			request, err := http.NewRequestWithContext(ctx, "POST", "/preference", strings.NewReader(`
				{
					"userID": "user1",
					"preferences": {}
				}
			`))

			Expect(err).To(BeNil())

			userPreferenceHandler.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusInternalServerError))
			Expect(response.Body.String()).To(ContainSubstring("Could not save user preferences"))
		})
	})
})
