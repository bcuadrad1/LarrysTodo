package actions

import (
	"fmt"
	"net/http"
	"net/url"
	"task/models"
	"time"

	"github.com/gobuffalo/uuid"
)

func (as *ActionSuite) Test_Get_Tasks_List_Empty() {
	res := as.HTML("/tasks").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "")
}

func (as *ActionSuite) Test_Get_Tasks_List_1_row() {
	task := models.Tasks{
		models.Task{
			ID:             uuid.FromStringOrNil("06607471-cbe1-4014-a615-e399d850fbb3"),
			Description:    "description",
			IsDone:         false,
			CompletionDate: time.Now().String(),
			RequestedBy:    "Larry",
			ExecutedBy:     "Bryan",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	as.NoError(as.DB.Create(&task))

	res := as.HTML("/tasks").Get()
	fmt.Println(res.Body.String())
	as.Equal(http.StatusOK, res.Code)

	for _, t := range task {
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.Description))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.IsDone))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.CompletionDate))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.RequestedBy))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.ExecutedBy))
	}
}

func (as *ActionSuite) Test_Get_Tasks_List_Pendin() {
	task := models.Tasks{
		models.Task{
			ID:             uuid.FromStringOrNil("06607471-cbe1-4014-a615-e399d850fbb3"),
			Description:    "description",
			IsDone:         false,
			CompletionDate: time.Now().String(),
			RequestedBy:    "Larry",
			ExecutedBy:     "Waldo",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		models.Task{
			ID:             uuid.FromStringOrNil("f9142486-611b-4a19-9129-afa83e36c983"),
			Description:    "description 2",
			IsDone:         false,
			CompletionDate: time.Now().String(),
			RequestedBy:    "Danier",
			ExecutedBy:     "Sra Ana",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
		models.Task{
			ID:             uuid.FromStringOrNil("81114510-0ad4-4a90-a4f7-55b4083b34b2"),
			Description:    "description 3",
			IsDone:         true,
			CompletionDate: time.Now().String(),
			RequestedBy:    "Antonio",
			ExecutedBy:     "Gonzalo",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	}

	as.NoError(as.DB.Create(&task))

	res := as.HTML("/pending_tasks").Get()
	as.Equal(http.StatusOK, res.Code)

	for _, t := range task {
		if t.IsDone {
			as.NotContains(res.Body.String(), fmt.Sprintf("%v", t.Description))
			as.NotContains(res.Body.String(), fmt.Sprintf("%v", t.IsDone))
			as.NotContains(res.Body.String(), fmt.Sprintf("%v", t.CompletionDate))
			as.NotContains(res.Body.String(), fmt.Sprintf("%v", t.RequestedBy))
			as.NotContains(res.Body.String(), fmt.Sprintf("%v", t.ExecutedBy))
			continue
		}
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.Description))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.IsDone))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.CompletionDate))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.RequestedBy))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.ExecutedBy))
	}
}

func (as *ActionSuite) Test_Create_Task() {
	tasks := models.Tasks{
		models.Task{
			Description:    "description",
			IsDone:         false,
			CompletionDate: time.Now().String(),
			RequestedBy:    "Bryan C",
			ExecutedBy:     "Tony jeje",
		},
		models.Task{
			Description:    "other description",
			IsDone:         false,
			CompletionDate: time.Now().String(),
			RequestedBy:    "Rodolfo",
			ExecutedBy:     "Larry",
		},
	}

	for _, t := range tasks {
		url := url.Values{
			"Description":    []string{t.Description},
			"IsDone":         []string{fmt.Sprintf("%v", t.IsDone)},
			"CompletionDate": []string{t.CompletionDate},
			"RequestedBy":    []string{t.RequestedBy},
			"ExecutedBy":     []string{t.ExecutedBy},
		}

		res := as.HTML("/tasks").Post(url)

		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.Description))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.IsDone))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.CompletionDate))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.RequestedBy))
		as.Contains(res.Body.String(), fmt.Sprintf("%v", t.ExecutedBy))
	}
}
