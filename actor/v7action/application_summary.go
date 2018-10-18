package v7action

import (
	"reflect"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
)

//go:generate counterfeiter . RouteActor

type RouteActor interface {
	GetApplicationRoutes(appGUID string) (v2action.Routes, v2action.Warnings, error)
}

// ApplicationSummary represents an application with its processes and droplet.
type ApplicationSummary struct {
	Application
	CurrentDroplet   Droplet
	ProcessSummaries ProcessSummaries
	Routes           []v2action.Route
}

// GetApplicationSummaryByNameAndSpace is temporary until we merge to master.
// Delete after master merge and rename GetApplicationSummaryByNameAndSpaceNew.
func (actor Actor) GetApplicationSummaryByNameAndSpace(appName string, spaceGUID string, withObfuscatedValues bool) (ApplicationSummary, Warnings, error) {
	return actor.GetApplicationSummaryByNameAndSpaceNew(appName, spaceGUID, withObfuscatedValues, nil)
}

// GetApplicationSummaryByNameAndSpace returns an application with process and
// instance stats.
func (actor Actor) GetApplicationSummaryByNameAndSpaceNew(appName string, spaceGUID string, withObfuscatedValues bool, routeActor RouteActor) (ApplicationSummary, Warnings, error) {
	app, allWarnings, err := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if err != nil {
		return ApplicationSummary{}, allWarnings, err
	}

	processSummaries, processWarnings, err := actor.getProcessSummariesForApp(app.GUID, withObfuscatedValues)
	allWarnings = append(allWarnings, processWarnings...)
	if err != nil {
		return ApplicationSummary{}, allWarnings, err
	}

	droplet, warnings, err := actor.GetCurrentDropletByApplication(app.GUID)
	allWarnings = append(allWarnings, Warnings(warnings)...)
	if err != nil {
		if _, ok := err.(actionerror.DropletNotFoundError); !ok {
			return ApplicationSummary{}, allWarnings, err
		}
	}

	var appRoutes []v2action.Route
	if routeActor != nil || (reflect.ValueOf(routeActor).Kind() == reflect.Ptr && reflect.ValueOf(routeActor).IsNil()) {
		routes, warnings, err := routeActor.GetApplicationRoutes(app.GUID)
		allWarnings = append(allWarnings, Warnings(warnings)...)
		if err != nil {
			if _, ok := err.(ccerror.ResourceNotFoundError); !ok {
				return ApplicationSummary{}, allWarnings, err
			}
		}
		appRoutes = routes
	}

	summary := ApplicationSummary{
		Application:      app,
		ProcessSummaries: processSummaries,
		CurrentDroplet:   droplet,
		Routes:           appRoutes,
	}
	return summary, allWarnings, nil
}
