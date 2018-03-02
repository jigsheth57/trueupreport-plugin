package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/jigsheth57/trueupreport-plugin/apihelper"
	"github.com/jigsheth57/trueupreport-plugin/models"
)

//UsageReportCmd the plugin
type UsageReportCmd struct {
	apiHelper apihelper.CFAPIHelper
}

// contains CLI flag values
type flagVal struct {
	OrgName string
	Format  string
}

func ParseFlags(args []string) flagVal {
	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)

	// Create flags
	orgName := flagSet.String("o", "", "-o orgName")
	format := flagSet.String("f", "format", "-f <csv>")

	err := flagSet.Parse(args[1:])
	if err != nil {

	}

	return flagVal{
		OrgName: string(*orgName),
		Format:  string(*format),
	}
}

//GetMetadata returns metatada
func (cmd *UsageReportCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "trueup-report",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 0,
			Build: 6,
		},
		Commands: []plugin.Command{
			{
				Name:     "trueup-report",
				HelpText: "Report AIs, SIs and memory usage for orgs and spaces",
				UsageDetails: plugin.Usage{
					Usage: "cf trueup-report [-o orgName] [-f <csv>]",
					Options: map[string]string{
						"o": "organization",
						"f": "format",
					},
				},
			},
		},
	}
}

//UsageReportCommand doer
func (cmd *UsageReportCmd) UsageReportCommand(args []string) {
	flagVals := ParseFlags(args)

	var orgs models.Orgs
	var err error
	var report models.Report

	if flagVals.OrgName != "" {
		org, err := cmd.getOrg(flagVals.OrgName)
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
		orgs = append(orgs, org)
	} else {
		orgs, err = cmd.getOrgs()
		if nil != err {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	report.Orgs = orgs

	if flagVals.Format == "csv" {
		fmt.Println(report.CSV())
	} else {
		fmt.Println(report.String())
	}
}

func (cmd *UsageReportCmd) getOrgs() ([]models.Org, error) {
	rawOrgs, err := cmd.apiHelper.GetOrgs()
	if nil != err {
		return nil, err
	}

	var orgs = []models.Org{}

	for _, o := range rawOrgs {
		orgDetails, err := cmd.getOrgDetails(o)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, orgDetails)
	}
	return orgs, nil
}

func (cmd *UsageReportCmd) getOrg(name string) (models.Org, error) {
	rawOrg, err := cmd.apiHelper.GetOrg(name)
	if nil != err {
		return models.Org{}, err
	}

	return cmd.getOrgDetails(rawOrg)
}

func (cmd *UsageReportCmd) getOrgDetails(o apihelper.Organization) (models.Org, error) {
	usage, err := cmd.apiHelper.GetOrgMemoryUsage(o)
	if nil != err {
		return models.Org{}, err
	}
	quota, err := cmd.apiHelper.GetQuotaMemoryLimit(o.QuotaURL)
	if nil != err {
		return models.Org{}, err
	}
	spaces, err := cmd.getSpaces(o.SpacesURL)
	if nil != err {
		return models.Org{}, err
	}

	return models.Org{
		Name:        o.Name,
		MemoryQuota: int(quota),
		MemoryUsage: int(usage),
		Spaces:      spaces,
	}, nil
}

func (cmd *UsageReportCmd) getSpaces(spaceURL string) ([]models.Space, error) {
	rawSpaces, err := cmd.apiHelper.GetOrgSpaces(spaceURL)
	if nil != err {
		return nil, err
	}
	var spaces = []models.Space{}
	for _, s := range rawSpaces {
		apps, services, err := cmd.getAppsAndServices(s.SummaryURL)
		if nil != err {
			return nil, err
		}
		spaces = append(spaces,
			models.Space{
				Name: s.Name,
				Apps: apps,
				Services: services,
			},
		)
	}
	return spaces, nil
}

func (cmd *UsageReportCmd) getAppsAndServices(summaryURL string) ([]models.App, []models.Service, error) {
	rawApps, rawServices, err := cmd.apiHelper.GetSpaceAppsAndServices(summaryURL)
	if nil != err {
		return nil, nil, err
	}
	var apps = []models.App{}
	var services = []models.Service{}
	for _, a := range rawApps {
		apps = append(apps, models.App{
			Actual: int(a.Actual),
			Desire: int(a.Desire),
			RAM:	int(a.RAM),
		})
	}
	for _, s := range rawServices {
		services = append(services, models.Service{
			Label: string(s.Label),
			ServicePlan: string(s.ServicePlan),
		})
	}
	return apps, services, nil
}

//Run runs the plugin
func (cmd *UsageReportCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "trueup-report" {
		cmd.apiHelper = apihelper.New(cli)
		cmd.UsageReportCommand(args)
	}
}

func main() {
	plugin.Start(new(UsageReportCmd))
}
