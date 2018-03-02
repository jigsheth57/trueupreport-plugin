# TrueupReport Plugin
This CF CLI Plugin to shows memory consumption and application & service instances (only part of service suite (RabbitMQ, Redis & MySQL)) for each org and space you have permission to access.

#Usage

For human readable output:

```
➜  trueupreport-plugin git:(master) ✗ cf trueup-report
Org Central is consuming 22650 MB of 25600 MB.
	Space development is consuming 200 MB memory (0%) of org quota.
		4 apps: 2 running 2 stopped
		6 app instances: 4 running, 2 stopped
		0 service instances of type Service Suite
	Space staging is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space production is consuming 0 MB memory (0%) of org quota.
		5 apps: 0 running 5 stopped
		5 app instances: 0 running, 5 stopped
		0 service instances of type Service Suite
	Space Workshop is consuming 0 MB memory (0%) of org quota.
		3 apps: 0 running 3 stopped
		3 app instances: 0 running, 3 stopped
		1 service instances of type Service Suite
	Space sales-tracker-dev is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space ross is consuming 0 MB memory (0%) of org quota.
		5 apps: 0 running 5 stopped
		5 app instances: 0 running, 5 stopped
		0 service instances of type Service Suite
	Space busch is consuming 0 MB memory (0%) of org quota.
		3 apps: 0 running 3 stopped
		4 app instances: 0 running, 4 stopped
		0 service instances of type Service Suite
	Space lock is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space ssahadevan is consuming 1024 MB memory (4%) of org quota.
		4 apps: 1 running 3 stopped
		4 app instances: 1 running, 3 stopped
		0 service instances of type Service Suite
	Space jigsheth is consuming 0 MB memory (0%) of org quota.
		1 apps: 0 running 1 stopped
		1 app instances: 0 running, 1 stopped
		0 service instances of type Service Suite
	Space erds is consuming 2048 MB memory (8%) of org quota.
		3 apps: 1 running 2 stopped
		4 app instances: 2 running, 2 stopped
		0 service instances of type Service Suite
	Space ford is consuming 0 MB memory (0%) of org quota.
		2 apps: 0 running 2 stopped
		2 app instances: 0 running, 2 stopped
		0 service instances of type Service Suite
	Space GM is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space pcf-demo is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space basler is consuming 256 MB memory (1%) of org quota.
		4 apps: 2 running 2 stopped
		4 app instances: 2 running, 2 stopped
		0 service instances of type Service Suite
	Space womack is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space sullivan is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space bbyers is consuming 1024 MB memory (4%) of org quota.
		5 apps: 1 running 4 stopped
		6 app instances: 2 running, 4 stopped
		1 service instances of type Service Suite
	Space PB-demo is consuming 0 MB memory (0%) of org quota.
		3 apps: 0 running 3 stopped
		3 app instances: 0 running, 3 stopped
		0 service instances of type Service Suite
	Space chjohnson is consuming 0 MB memory (0%) of org quota.
		12 apps: 0 running 12 stopped
		13 app instances: 0 running, 13 stopped
		2 service instances of type Service Suite
	Space ripka is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space sf is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space phopper is consuming 10240 MB memory (40%) of org quota.
		9 apps: 6 running 3 stopped
		13 app instances: 10 running, 3 stopped
		1 service instances of type Service Suite
	Space jimerson is consuming 0 MB memory (0%) of org quota.
		1 apps: 0 running 1 stopped
		1 app instances: 0 running, 1 stopped
		0 service instances of type Service Suite
	Space guna is consuming 3072 MB memory (12%) of org quota.
		13 apps: 3 running 10 stopped
		13 app instances: 3 running, 10 stopped
		1 service instances of type Service Suite
	Space nrudnik is consuming 128 MB memory (0%) of org quota.
		3 apps: 1 running 2 stopped
		4 app instances: 1 running, 3 stopped
		0 service instances of type Service Suite
	Space akadari is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space cjett is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space jmclaughlin is consuming 2048 MB memory (8%) of org quota.
		2 apps: 2 running 0 stopped
		2 app instances: 2 running, 0 stopped
		1 service instances of type Service Suite
	Space mmcnichol is consuming 50 MB memory (0%) of org quota.
		1 apps: 1 running 0 stopped
		1 app instances: 1 running, 0 stopped
		0 service instances of type Service Suite
Org S1Pdemo14 is consuming 7168 MB of 102400 MB.
	Space development is consuming 2048 MB memory (2%) of org quota.
		3 apps: 2 running 1 stopped
		3 app instances: 2 running, 1 stopped
		1 service instances of type Service Suite
	Space IoT-ConnectedCar-Emulator is consuming 1024 MB memory (1%) of org quota.
		1 apps: 1 running 0 stopped
		1 app instances: 1 running, 0 stopped
		1 service instances of type Service Suite
	Space sandbox is consuming 0 MB memory (0%) of org quota.
		1 apps: 0 running 1 stopped
		2 app instances: 0 running, 2 stopped
		0 service instances of type Service Suite
	Space auto-2 is consuming 0 MB memory (0%) of org quota.
		1 apps: 0 running 1 stopped
		3 app instances: 0 running, 3 stopped
		0 service instances of type Service Suite
	Space scdf-twitter-demo is consuming 2048 MB memory (2%) of org quota.
		2 apps: 2 running 0 stopped
		2 app instances: 2 running, 0 stopped
		1 service instances of type Service Suite
	Space scs-demo is consuming 0 MB memory (0%) of org quota.
		7 apps: 0 running 7 stopped
		8 app instances: 0 running, 8 stopped
		2 service instances of type Service Suite
	Space cfml is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space meetup-rsvp-demo is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space gt-scdf-twitter-demo is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
You have deployed 98 apps across 2 org(s), with a total of 113 app instances configured. You are currently running 25 apps with 33 app instances and using 12 service instances of type Service Suite.

```

CSV output:

```
you can skip header using "-h skip" flag. helps if you are running across multiple foundation and would like to append the result!
➜  trueupreport-plugin git:(master) ✗ cf trueup-report -f csv -h skip

➜  trueupreport-plugin git:(master) ✗ cf trueup-report -f csv
OrgName, SpaceName, SpaceMemoryUsed, OrgMemoryQuota, AppsDeployed, AppsRunning, AppInstancesDeployed, AppInstancesRunning, TotalServiceInstancesDeployed, RabbitMQServiceInstanceDeployed, RedisServiceInstanceDeployed, MySQLServiceInstanceDeployed
Central, development, 200, 25600, 4, 2, 6, 4, 0, 0, 0, 0
Central, staging, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, production, 0, 25600, 5, 0, 5, 0, 0, 0, 0, 0
Central, Workshop, 0, 25600, 3, 0, 3, 0, 1, 0, 0, 1
Central, sales-tracker-dev, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, ross, 0, 25600, 5, 0, 5, 0, 0, 0, 0, 0
Central, busch, 0, 25600, 3, 0, 4, 0, 0, 0, 0, 0
Central, lock, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, ssahadevan, 1024, 25600, 4, 1, 4, 1, 0, 0, 0, 0
Central, jigsheth, 0, 25600, 1, 0, 1, 0, 0, 0, 0, 0
Central, erds, 2048, 25600, 3, 1, 4, 2, 0, 0, 0, 0
Central, ford, 0, 25600, 2, 0, 2, 0, 0, 0, 0, 0
Central, GM, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, pcf-demo, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, basler, 256, 25600, 4, 2, 4, 2, 0, 0, 0, 0
Central, womack, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, sullivan, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, bbyers, 1024, 25600, 5, 1, 6, 2, 1, 0, 0, 1
Central, PB-demo, 0, 25600, 3, 0, 3, 0, 0, 0, 0, 0
Central, chjohnson, 0, 25600, 12, 0, 13, 0, 2, 0, 1, 1
Central, ripka, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, sf, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, phopper, 10240, 25600, 9, 6, 13, 10, 1, 0, 1, 0
Central, jimerson, 0, 25600, 1, 0, 1, 0, 0, 0, 0, 0
Central, guna, 3072, 25600, 13, 3, 13, 3, 1, 0, 0, 1
Central, nrudnik, 128, 25600, 3, 1, 4, 1, 0, 0, 0, 0
Central, akadari, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, cjett, 0, 25600, 0, 0, 0, 0, 0, 0, 0, 0
Central, jmclaughlin, 2048, 25600, 2, 2, 2, 2, 1, 0, 1, 0
Central, mmcnichol, 50, 25600, 1, 1, 1, 1, 0, 0, 0, 0
S1Pdemo14, development, 2048, 102400, 3, 2, 3, 2, 1, 1, 0, 0
S1Pdemo14, IoT-ConnectedCar-Emulator, 1024, 102400, 1, 1, 1, 1, 1, 0, 1, 0
S1Pdemo14, sandbox, 0, 102400, 1, 0, 2, 0, 0, 0, 0, 0
S1Pdemo14, auto-2, 0, 102400, 1, 0, 3, 0, 0, 0, 0, 0
S1Pdemo14, scdf-twitter-demo, 2048, 102400, 2, 2, 2, 2, 1, 0, 1, 0
S1Pdemo14, scs-demo, 0, 102400, 7, 0, 8, 0, 2, 1, 1, 0
S1Pdemo14, cfml, 0, 102400, 0, 0, 0, 0, 0, 0, 0, 0
S1Pdemo14, meetup-rsvp-demo, 0, 102400, 0, 0, 0, 0, 0, 0, 0, 0
S1Pdemo14, gt-scdf-twitter-demo, 0, 102400, 0, 0, 0, 0, 0, 0, 0, 0
```

##Installation
```
For OSX
$ cf install-plugin https://github.com/jigsheth57/trueupreport-plugin/blob/master/bin/osx/trueupreport-plugin?raw=true -f

For Windows 32bit
$ cf install-plugin https://github.com/jigsheth57/trueupreport-plugin/blob/master/bin/win32/trueupreport-plugin.exe?raw=true -f

For Windows 64bit
$ cf install-plugin https://github.com/jigsheth57/trueupreport-plugin/blob/master/bin/win64/trueupreport-plugin.exe?raw=true -f

For Linux 64bit
$ cf install-plugin https://github.com/jigsheth57/trueupreport-plugin/blob/master/bin/linux64/trueupreport-plugin?raw=true -f

```
#####Install from Source (need to have [Go](http://golang.org/dl/) installed)
  ```
  $ go get github.com/cloudfoundry/cli
  $ go get github.com/jigsheth57/trueupreport-plugin
  $ cd $GOPATH/src/github.com/jigsheth57/trueupreport-plugin
  $ go build
  $ cf install-plugin trueupreport-plugin
  ```
