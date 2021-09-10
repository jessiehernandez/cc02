package repository_test

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/testing/example/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Preference Repository", func() {
	var (
		db   *sql.DB
		err  error
		repo *repository.UserPreference
	)

	ctx := context.Background()

	BeforeEach(func() {
		db, err = sql.Open("pgx", "postgres://jessie.hernandez@localhost:5432/testing?sslmode=disable")

		Expect(err).To(BeNil())
		Expect(db).ToNot(BeNil())

		_, _ = db.ExecContext(ctx, `TRUNCATE TABLE example.user_preference`)

		repo = repository.NewUserPreference(db)
		Expect(repo).ToNot(BeNil())
	})

	AfterEach(func() {
		_, _ = db.ExecContext(ctx, `ALTER TABLE IF EXISTS example.user_preference_bak RENAME TO user_preference`)
	})

	Context("Saving a user's preferences", func() {
		When("the preferences are in an invalid format", func() {
			It("returns an error", func() {
				err = repo.Save(ctx, "user1", func() {})

				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("preferences passed in an invalid format"))
			})
		})

		When("the preferences are saved successfully", func() {
			It("stores the preferences in the database", func() {
				err = repo.Save(ctx, "user1", map[string]string{"foo": "bar"})

				Expect(err).To(BeNil())
			})
		})

		When("there is an error saving the preferences", func() {
			It("returns an error", func() {
				db.ExecContext(ctx, `ALTER TABLE example.user_preference RENAME TO user_preference_bak`)
				err = repo.Save(ctx, "user1", nil)

				Expect(err).ToNot(BeNil())
			})
		})
	})
})
