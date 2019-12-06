package repositories

import (
	"testing"
)

func getFigmaRep(t *testing.T) FigmaFileDBRepository {
	db, err := NewTestDB()

	if err != nil {
		t.FailNow()
	}

	return FigmaFileDBRepository{DB: db}
}

func TestReadAllFiles(t *testing.T) {

	// Arrange

	rep := getFigmaRep(t)

	// Act

	files, err := rep.ReadAllFiles()

	// Assert

	if err != nil {
		t.Error(err)
		return
	}

	if len(files) == 0 {
		t.Error("No one file found")
	}
}

func TestUpdateFile(t *testing.T) {

	// Arrange

	rep := getFigmaRep(t)

	// Act

	file, _ := rep.ReadAllFiles()
	file[0].FileVersion.Label = "gfhsvhjs"

	err := rep.UpdateFile(file[0])

	// Assert

	if err != nil {
		t.Error(err)
		return
	}
}
