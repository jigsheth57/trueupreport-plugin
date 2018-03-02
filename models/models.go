package models

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Org struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      Spaces
}

type Space struct {
	Name string
	Apps Apps
	Services Services
}

//App representation
type App struct {
	Actual	int
	Desire	int
	RAM     int
}

//Service representation
type Service struct {
	Label    	string
	ServicePlan string
}

type SpaceStats struct {
	Name string
	DeployedAppsCount int
	RunningAppsCount int
	StoppedAppsCount int
	DeployedAppInstancesCount int
	RunningAppInstancesCount int
	StoppedAppInstancesCount int
	ServicesCount int
	ConsumedMemory int
}

type OrgStats struct {
	Name        string
	MemoryQuota int
	MemoryUsage int
	Spaces      Spaces
	DeployedAppsCount	int
	RunningAppsCount	int
	StoppedAppsCount	int
	DeployedAppInstancesCount int
	RunningAppInstancesCount	int
	StoppedAppInstancesCount	int
	ServicesCount	int
}

type Orgs []Org
type Spaces []Space
type Apps []App
type Services []Service

type Report struct {
	Orgs Orgs
}

func (org *Org) InstancesCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.InstancesCount()
	}
	return instancesCount
}

func (org *Org) RunningAppsCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.RunningAppsCount()
	}
	return instancesCount
}

func (org *Org) RunningInstancesCount() int {
	instancesCount := 0
	for _, space := range org.Spaces {
		instancesCount += space.RunningInstancesCount()
	}
	return instancesCount
}

func (org *Org) AppsCount() int {
	appsCount := 0
	for _, space := range org.Spaces {
		appsCount += len(space.Apps)
	}
	return appsCount
}

func (org *Org) ServicesCount() int {
	servicesCount := 0
	for _, space := range org.Spaces {
		servicesCount += len(space.Services)
	}
	return servicesCount
}

func (space *Space) ConsumedMemory() int {
	consumed := 0
	for _, app := range space.Apps {
		consumed += int(app.Actual * app.RAM)
	}
	return consumed
}

func (space *Space) RunningAppsCount() int {
	runningAppsCount := 0
	for _, app := range space.Apps {
		if (app.Actual > 0) {
			runningAppsCount++
		}
	}
	return runningAppsCount
}

func (space *Space) InstancesCount() int {
	instancesCount := 0
	for _, app := range space.Apps {
		instancesCount += int(app.Desire)
	}
	return instancesCount
}

func (space *Space) RunningInstancesCount() int {
	runningInstancesCount := 0
	for _, app := range space.Apps {
		runningInstancesCount += int(app.Actual)
	}
	return runningInstancesCount
}

func (space *Space) ServicesCount() int {
	servicesCount := len(space.Services)
	return servicesCount
}

func (space *Space) ServiceInstancesCount(serviceType string) int {
	boundedServiceInstancesCount := 0
	for _, service := range space.Services {
		if(strings.Contains(service.Label,serviceType)) {
			boundedServiceInstancesCount++
		}
	}
	return boundedServiceInstancesCount
}

func (spaces Spaces) Stats (c chan SpaceStats, skipSIcount bool) {
	for _, space := range spaces {
		lApps := len(space.Apps)
		rApps := space.RunningAppsCount()
		sApps := lApps-rApps
		lAIs := space.InstancesCount()
		rAIs := space.RunningInstancesCount()
		sAIs := lAIs-rAIs
		siCount := space.ServicesCount()
		if(skipSIcount) {
			siCount = 0
		}
		c <- SpaceStats{
			Name: space.Name,
			DeployedAppsCount: lApps,
			RunningAppsCount: rApps,
			StoppedAppsCount: sApps,
			DeployedAppInstancesCount: lAIs,
			RunningAppInstancesCount: rAIs,
			StoppedAppInstancesCount: sAIs,
			ServicesCount: siCount,
			ConsumedMemory: space.ConsumedMemory(),
		}
	}
	close(c)
}

func (orgs Orgs) Stats (c chan OrgStats) {
	for _, org := range orgs {
		lApps := org.AppsCount()
		rApps := org.RunningAppsCount()
		sApps := lApps-rApps
		lAIs := org.InstancesCount()
		rAIs := org.RunningInstancesCount()
		sAIs := lAIs-rAIs
		c <- OrgStats{
			Name: org.Name,
			MemoryQuota: org.MemoryQuota,
			MemoryUsage: org.MemoryUsage,
			Spaces: org.Spaces,
			DeployedAppsCount: lApps,
			RunningAppsCount: rApps,
			StoppedAppsCount: sApps,
			DeployedAppInstancesCount: lAIs,
			RunningAppInstancesCount: rAIs,
			StoppedAppInstancesCount: sAIs,
			ServicesCount: org.ServicesCount(),
		}
	}
	close(c)
}

func (report *Report) String() string {
	var response bytes.Buffer

	totalApps := 0
	totalInstances := 0
	totalRunningApps := 0
	totalRunningInstances := 0
	totalServiceInstances := 0

	chOrgStats := make(chan OrgStats, len(report.Orgs))

	go report.Orgs.Stats(chOrgStats)
	for orgStats := range chOrgStats {
		response.WriteString(fmt.Sprintf("Org %s is consuming %d MB of %d MB.\n",
			orgStats.Name, orgStats.MemoryUsage, orgStats.MemoryQuota))
		chSpaceStats := make(chan SpaceStats, len(orgStats.Spaces))
		go orgStats.Spaces.Stats(chSpaceStats,orgStats.Name == "p-spring-cloud-services")
		for spaceState := range chSpaceStats {
			response.WriteString(
				fmt.Sprintf("\tSpace %s is consuming %d MB memory (%d%%) of org quota.\n",
					spaceState.Name, spaceState.ConsumedMemory, (100 * spaceState.ConsumedMemory / orgStats.MemoryQuota)))
			response.WriteString(
				fmt.Sprintf("\t\t%d apps: %d running %d stopped\n", spaceState.DeployedAppsCount,
					spaceState.RunningAppsCount, spaceState.StoppedAppsCount))
			response.WriteString(
				fmt.Sprintf("\t\t%d app instances: %d running, %d stopped\n", spaceState.DeployedAppInstancesCount,
					spaceState.RunningAppInstancesCount, spaceState.StoppedAppInstancesCount))
			response.WriteString(
				fmt.Sprintf("\t\t%d service instances of type Service Suite\n", spaceState.ServicesCount))
		}
		totalApps += orgStats.DeployedAppsCount
		totalInstances += orgStats.DeployedAppInstancesCount
		totalRunningApps += orgStats.RunningAppsCount
		totalRunningInstances += orgStats.RunningAppInstancesCount
		totalServiceInstances += orgStats.ServicesCount
	}
	response.WriteString(
		fmt.Sprintf("You have deployed %d apps across %d org(s), with a total of %d app instances configured. You are currently running %d apps with %d app instances and using %d service instances of type Service Suite.\n",
			totalApps, len(report.Orgs), totalInstances, totalRunningApps, totalRunningInstances, totalServiceInstances))

	return response.String()
}

func (report *Report) CSV(skipHeaders bool) string {
	var rows = [][]string{}
	var csv bytes.Buffer

	var headers = []string{"OrgName", "SpaceName", "SpaceMemoryUsed", "OrgMemoryQuota", "AppsDeployed", "AppsRunning", "AppInstancesConfigured", "AppInstancesRunning", "TotalServiceInstancesDeployed", "RabbitMQServiceInstanceDeployed", "RedisServiceInstanceDeployed", "MySQLServiceInstanceDeployed"}

	if !skipHeaders {
		rows = append(rows, headers)
	}

	for _, org := range report.Orgs {
		for _, space := range org.Spaces {
			appsDeployed := len(space.Apps)

			col1 := org.Name
			col2 := space.Name
			col3 := strconv.Itoa(space.ConsumedMemory())
			col4 := strconv.Itoa(org.MemoryQuota)
			col5 := strconv.Itoa(appsDeployed)
			col6 := strconv.Itoa(space.RunningAppsCount())
			col7 := strconv.Itoa(space.InstancesCount())
			col8 := strconv.Itoa(space.RunningInstancesCount())
			col9 := strconv.Itoa(space.ServicesCount())
			col10 := strconv.Itoa(space.ServiceInstancesCount("rabbit"))
			col11 := strconv.Itoa(space.ServiceInstancesCount("redis"))
			col12 := strconv.Itoa(space.ServiceInstancesCount("mysql"))
			if (org.Name == "p-spring-cloud-services") {
				col9 = "0"
				col10 = "0"
				col11 = "0"
				col12 = "0"
			}

			spaceResult := []string{
				col1,
				col2,
				col3,
				col4,
				col5,
				col6,
				col7,
				col8,
				col9,
				col10,
				col11,
				col12,
			}
			rows = append(rows, spaceResult)
		}
	}

	for i := range rows {
		csv.WriteString(strings.Join(rows[i], ", "))
		csv.WriteString("\n")
	}

	return csv.String()
}
