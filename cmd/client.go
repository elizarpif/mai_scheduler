package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"mai_scheduler/config"
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

}

//   <div class="sc-table-row">
//   <a class="sc-table-col" href="#fac-3-Институт-№8" role="button" data-toggle="collapse" aria-expanded="false" aria-controls="fac-3-Институт-№8">Институт №8</a>
//   <div class="sc-table-col">
//   <div class="sc-table-col-body " id="fac-3-Институт-№8">
//  <div class="sc-groups">
// <div class="sc-program">Аспирантура<
// -  <a class="sc-group-item" href="detail.php?group=М8О-301А-17">М8О-301А-17</a>

func (s *MaiClient) GetGroupsByInstituteAndCourse(URL string) {
	resp, err := s.Client.Get(URL)

	if err != nil {
		s.Logger.WithError(err).Error("cannot get schedule")
		fmt.Println(err)
		return
	}

	pages, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println()
	page := string(pages[:])
	cutstring := strings.Index(page, `<div class="sc-container">`)
	rightcut := strings.Index(page, `<!-- block footer -->`)
	fmt.Println(page[cutstring:rightcut])

	defer func() {
		_ = resp.Body.Close()
	}()

	//_, _ = io.Copy(os.Stdout, resp.Body)
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
