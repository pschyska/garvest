package harvest

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/config"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"time"
)

type Harvest struct {
	domain   string
	username string
	password string
}

func loadProjects(b []byte) (*[]Project, error) {
	var data = &[]Project{}
	err := json.Unmarshal(b, data)
	return data, err
}

type Project struct {
	Project ProjectDetail
}

type ProjectDetail struct {
	Id                        uint32
	Client_id                 uint32
	Name                      string
	Code                      string
	Active                    bool
	Notes                     string
	Billable                  bool
	BillBy                    string
	CostBudget                uint32
	CostBudgetIncludeExpenses bool
	HourlyRate                float32
	// budget float32
	BudgetBy string
	// notify_when_over_budget             bool
	// over_budget_notification_percentage float32
	// // over_budget_notified_at date time.Time
	// show_budget_to_all bool
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// estimate":null,"estimate_by":"none","hint_earliest_record_at":"2013-02-11","hint_latest_record_at":"2013-08-20"
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func New() Harvest {
	c, err := config.ReadDefault(getHome() + "/.garvest.cfg")
	if err != nil {
		log.Fatal(err)
	}
	domain, _ := c.String("api", "domain")
	username, _ := c.String("auth", "username")
	password, _ := c.String("auth", "password")
	return Harvest{domain: domain, username: username, password: password}
}

func (h Harvest) Connect() (string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://"+h.domain+"/projects", nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(h.username, h.password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "garvest")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	projects, err := loadProjects(body)
	if err != nil {
		return "", err
	}
	for _, v := range *projects {
		fmt.Println("Project! %v", v.Project)
	}

	return string(body), nil
}
