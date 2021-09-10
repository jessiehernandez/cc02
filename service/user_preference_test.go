package service_test

import (
	"context"
	"errors"

	"github.com/testing/example/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type userPreferenceRepositoryStub struct {
	ReturnError error
}

func (u *userPreferenceRepositoryStub) Save(ctx context.Context, userID string, preferences interface{}) (err error) {
	return u.ReturnError
}

var _ = Describe("User Preference Service", func() {
	ctx := context.Background()

	When("the preferences are successfully saved", func() {
		It("returns success", func() {
			srv := service.NewUserPreference(&userPreferenceRepositoryStub{})
			err := srv.Save(ctx, "foo", map[string]string{"foo": "bar"})

			Expect(err).To(BeNil())
		})
	})

	When("there is an error saving the preferences", func() {
		It("returns an error", func() {
			srv := service.NewUserPreference(&userPreferenceRepositoryStub{
				ReturnError: errors.New("mock error"),
			})
			err := srv.Save(ctx, "foo", map[string]string{"foo": "bar"})

			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("mock error"))
		})
	})
})
