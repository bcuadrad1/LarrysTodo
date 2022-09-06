package actions

import (
	"database/sql"
	"net/http"

	"github.com/bcuadrad1/LarrysTodo/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

func TaskList(c buffalo.Context) error {
	tasks := models.Tasks{}

	tx := c.Value("tx").(*pop.Connection)
	if err := tx.All(&tasks); err != nil {
		if errors.Cause(err) != sql.ErrNoRows {
			return errors.WithStack(errors.Wrap(err, "Error in the db"))
		}
	}

	return c.Render(http.StatusOK, r.JSON(tasks))
}

func TaskPendingList(c buffalo.Context) error {
	tasks := models.Tasks{}

	tx := c.Value("tx").(*pop.Connection)
	if err := tx.Where("is_done = ?", false).All(&tasks); err != nil {
		if errors.Cause(err) != sql.ErrNoRows {
			return errors.WithStack(errors.Wrap(err, "Error in the db"))
		}
	}

	return c.Render(http.StatusOK, r.JSON(tasks))
}

func TaskCreate(c buffalo.Context) error {
	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "Error binding element"))
	}

	tx := c.Value("tx").(*pop.Connection)
	if err := tx.Create(&task); err != nil {
		return errors.WithStack(errors.Wrap(err, "Error when creating the task"))
	}

	return c.Render(http.StatusOK, r.JSON(task))
}
