package repositories

import (
	"testing"

	"github.com/LastSprint/InfrastructureSlackApp/models"
)

func getRep(t *testing.T) UserDBRepository {
	db, err := NewTestDB()

	if err != nil {
		t.FailNow()
	}

	return UserDBRepository{DB: db}
}

func TestReadAllDevelopers(t *testing.T) {

	// Arrange

	rep := getRep(t)

	// Act

	users, err := rep.ReadAllDevelopers()

	// Assert

	if err != nil {
		t.Error(err)
		return
	}

	if len(users) == 0 {
		t.Error("Noone user found")
	}

	for _, item := range users {
		if item.Member.Role != models.Developer && item.Member.Role != models.Lead {
			t.Error("Only lead and developers but found" + item.Member.ToString())
		}

		if item.Member.Department != models.IOS && item.Member.Department != models.Android {
			t.Error("Only iOS and Android" + item.Member.ToString())
		}
	}
}

func TestReadDevelopers(t *testing.T) {

	// Arrange

	rep := getRep(t)

	// Act

	users, err := rep.ReadDevelopers()

	// Assert

	if err != nil {
		t.Error(err)
		return
	}

	if len(users) == 0 {
		t.Error("Noone user found")
	}

	for _, item := range users {
		if item.Member.Role != models.Developer {
			t.Error("Only developers but found" + item.Member.ToString())
		}

		if item.Member.Department != models.IOS && item.Member.Department != models.Android {
			t.Error("Only iOS and Android" + item.Member.ToString())
		}
	}
}

func TestReadLeadDevelopers(t *testing.T) {

	// Arrange

	rep := getRep(t)

	// Act

	users, err := rep.ReadLeadDevelopers()

	// Assert

	if err != nil {
		t.Error(err)
		return
	}

	if len(users) == 0 {
		t.Error("Noone user found")
	}

	for _, item := range users {
		if item.Member.Role != models.Lead {
			t.Error("Only leads but found" + item.Member.ToString())
		}

		if item.Member.Department != models.IOS && item.Member.Department != models.Android {
			t.Error("Only iOS and Android" + item.Member.ToString())
		}
	}
}

func TestReadManagers(t *testing.T) {

	// Arrange

	rep := getRep(t)

	// Act

	users, err := rep.ReadManagers()

	// Assert

	if err != nil {
		t.Error(err)
		return
	}

	if len(users) == 0 {
		t.Error("Noone user found")
	}

	for _, item := range users {
		if item.Member.Role != models.Manager {
			t.Error("Only managers but found" + item.Member.ToString())
		}

		if item.Member.Department != models.Managers {
			t.Error("Only iOS and Android" + item.Member.ToString())
		}
	}
}
