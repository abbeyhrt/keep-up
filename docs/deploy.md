# Deployments

Currently, the project uses IBM Cloud for cloudfoundry-style deployments. In
this case, that means that we have a `manifest.yml` file in the root of our
directory that specifies a couple attributes about the Cloudfoundry service that
we are deploying.

These attributes include things like:

- `applications`: a list of applications that we are deploying
- `name`: the name of the Cloudfoundry service we're creating
- `buildpack`: the tools needed to run our specific service in the Cloudfoundry
  environment
- `command`: the command used to start our service
- `memory`: how much RAM is available to our service
- `disk_quota`: how much storage is available on disk for our service and its
  dependencies
- `env`: a set of key-value pairs that set ENV variables in the process that our
  service will be running in
