package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTask_Validate(t *testing.T) {
	// empty title
	t1 := Task{Title: "", Description: "d"}
	assert.ErrorIs(t, t1.Validate(), ErrInvalidTaskTitle)

	// empty description
	t2 := Task{Title: "t", Description: "  \t  "}
	assert.ErrorIs(t, t2.Validate(), ErrInvalidTaskDescription)

	// valid
	t3 := Task{UserID: 1, Title: "t", Description: "d", DueDate: time.Now(), Status: "open"}
	assert.NoError(t, t3.Validate())
}

func TestUser_IsAdmin(t *testing.T) {
	u1 := User{Role: "Admin"}
	u2 := User{Role: "user"}
	assert.True(t, u1.IsAdmin())
	assert.False(t, u2.IsAdmin())
}
