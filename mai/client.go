package mai

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"

	"mai_scheduler/config"
	"mai_scheduler/utils"
)

type MaiClient struct {
	Client  http.Client
	Logger  logrus.Logger
	Address string
	Config  *config.Config
}

func (s *MaiClient) InitClient() {
	s.Client = http.Client{}
	s.Logger = logrus.Logger{}
	if (s.Config != nil) {
		s.Address = s.Config.Address
	}

}

type ScheduleDay struct {
	Day   string
	ScDay string
	Rows  []*ScheduleRow
}

type ScheduleRow struct {
	StartAt    string
	EndAt      string
	Lecturer   string
	LessonType string
	Title      string
	Location   string
}

func (c *MaiClient) GetScheduleByGroup(group string) ([]*ScheduleDay, error) {
	urlParse, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.WithError(err).Error("cannot parse address")
		return nil, err
	}

	urlParse.Path = path.Join(urlParse.Path, "detail.php")

	query := urlParse.Query()
	query.Add("group", group)

	urlParse.RawQuery = query.Encode()

	return c.getSchedule(urlParse.String())
}

func (s *MaiClient) getSchedule(URL string) ([]*ScheduleDay, error) {
	resp, err := s.Client.Get(URL)

	if err != nil {
		s.Logger.WithError(err).Error("cannot get schedule")
		return nil, err
	}

	pages, _ := ioutil.ReadAll(resp.Body)
	page := string(pages[:])

	st, err := utils.ParseHtml(page)
	if st == nil {
		s.Logger.WithError(err).Error("cannot get schedule")
		return nil, NewErrNotFound()
	}

	sts := st.FindAllSt("sc-container")

	schedules := []*ScheduleDay{}

	for _, s := range sts {
		sctableRow := s.FindSt("sc-table-row")

		days := sctableRow.FindAllValues("sc-table-col sc-day-header sc-blue")
		if len(days) == 0 {
			days = sctableRow.FindAllValues("sc-table-col sc-day-header sc-gray")
		}
		scday := sctableRow.FindAllValues("sc-day")

		sctableRows := sctableRow.FindAllChildSt("sc-table-row")
		schRows := make([]*ScheduleRow, 0, len(sctableRows))

		for _, scRow := range sctableRows {
			timeDur := scRow.FindAllValues("sc-table-col sc-item-time")
			timeDur = strings.Split(timeDur[0], " &ndash; ")

			lecType := scRow.FindAllValues("sc-table-col sc-item-type")
			lecTitle := scRow.FindAllValues("sc-title")
			lecturer := scRow.FindAllValues("sc-lecturer")
			location := scRow.FindAllValues("sc-table-col sc-item-location")

			sch := &ScheduleRow{
				StartAt:    timeDur[0],
				EndAt:      timeDur[1],
				Lecturer:   strings.Join(lecturer, " "),
				LessonType: strings.Join(lecType, " "),
				Title:      strings.Join(lecTitle, " "),
				Location:   strings.Join(location, " "),
			}
			schRows = append(schRows, sch)

		}

		schedule := &ScheduleDay{
			Day:   strings.Join(days, " "),
			ScDay: strings.Join(scday, " "),
			Rows:  schRows,
		}
		spew.Dump(schedule)

		schedules = append(schedules, schedule)
	}
	return schedules, nil

}

type Info struct {
	Groups    []string
	Institute string
	Course    string
}

func (c *MaiClient) GetGroups(dept, course int) (*Info, error) {
	urlParse, err := url.Parse(c.Address)
	if err != nil {
		c.Logger.WithError(err).Error("cannot parse address")
		return nil, err
	}
	query := urlParse.Query()

	query.Add("department", strconv.Itoa(c.getDepartment(dept)))
	query.Add("course", strconv.Itoa(course))

	urlParse.RawQuery = query.Encode()

	return c.getGroupsByInstituteAndCourse(urlParse.String())
}

func (s *MaiClient) getGroupsByInstituteAndCourse(URL string) (*Info, error) {
	resp, err := s.Client.Get(URL)

	if err != nil {
		s.Logger.WithError(err).Error("cannot get schedule")
		return nil, err
	}

	pages, _ := ioutil.ReadAll(resp.Body)
	page := string(pages[:])

	st, err := utils.ParseHtml(page)
	if st == nil {
		return nil, NewErrNotFound()
	}
	st = st.FindSt("sc-container")

	courses := st.FindAllValues("sc-container-header sc-gray")
	if len(courses) == 0 {
		return nil, NewErrNotFound()
	}

	inst := st.FindAllValues("sc-table-col")
	if len(inst) == 0 {
		return nil, NewErrNotFound()
	}

	groups := st.FindAllValues("sc-group-item")
	if len(groups) == 0 {
		return nil, NewErrNotFound()
	}
	info := &Info{
		Course:    courses[0],
		Institute: inst[0],
		Groups:    groups,
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	return info, nil

}
func (s *MaiClient) getDepartment(dep int) (int) {
	deps := s.Config.Departments
	switch (dep) {
	case 1:
		return deps.Inst1
	case 2:
		return deps.Inst2
	case 3:
		return deps.Inst3
	case 4:
		return deps.Inst4
	case 5:
		return deps.Inst5
	case 6:
		return deps.Inst6
	case 7:
		return deps.Inst7
	case 8:
		return deps.Inst8
	case 9:
		return deps.Inst9
	case 10:
		return deps.Inst10
	case 11:
		return deps.Inst11
	case 12:
		return deps.Inst12

	default:
		return 0
	}
}
