# intigriti-cicd-plugin
Tool that can block your CI/CD pipeline depending on outstanding (open) intigriti issues.

## Setup
1. Download [the latest icp release](https://github.com/hazcod/intigriti-cicd-plugin/releases).
2. Retrieve your [intigriti API token](https://intigriti.com/) and pass your (external) IP address for whitelisting.
3. Create your configuration file:
```yaml
# your intigriti API credentials
intigriti_client_id: "XXXXXXXXXXX"
intigriti_client_secret: "XXXXXXXXXXX"

# what maximum amount of findings you tolerate per severity
tresholds:
  # we allow no criticals
  critical: 0
  # we allow no highs
  high: 0
  # we allow 1 medium
  medium: 1
  # we allow arbitrary amount of lows
  low: 100000
```
5. Run `icp` in your CI/CD pipeline with arguments:
```shell
./icp -conf=my-conf.yml
```
3. `icp` will return an error code whenever your defined tresholds are set, stopping your pipeline.

## Building
This requires `make` and `go` to be installed.
Just run `make`.