# Deployments

Currently, the project uses IBM Cloud for cloudfoundry-style deployments. In
this case, that means that we have a `manifest.yml` file in the root of our
directory that specifies a couple attributes about the Cloudfoundry service that
we are deploying.

These attributes include things like:

* `applications`: a list of applications that we are deploying
* `name`: the name of the Cloudfoundry service we're creating
* `buildpack`: the tools needed to run our specific service in the Cloudfoundry
  environment
* `command`: the command used to start our service
* `memory`: how much RAM is available to our service
* `disk_quota`: how much storage is available on disk for our service and its
  dependencies
* `env`: a set of key-value pairs that set ENV variables in the process that our
  service will be running in

Raw Changes:

First things first:

CI = Continuous Integration

CD = Continuous Deployment

**Specs:**

For this project, we are using Travis CI to deploy the project once a new PR is merged into the master branch. CloudFoundry hosts our deployment [here](keep-up-graphql.mybluemix.net).

**Why continuously integrate new code?**

A common pattern I've noticed in what I think is "good code"(whatever that means exactly, I'm not really sure yet) is the prevention of problems before they occur and catching errors as soon as they occur. CI falls in the latter group.
By continuous deploying, and maybe having that deploy fail, it seems like it would be easier to figure out what exactly caused the build to fail. In CI, you are deploying small changes that were made in a single PR meaning that you've already narrowed down the problem to whatever new code was introduced with the PR. Obviously, that makes finding the problem easier.

**How we set up this build/deploy with Travis and CloudFoundry:**

We are using CloudFoundry as the provider and for that you need an account. If you don't have one go [here](http://console.ng.bluemix.net/). After that , create a resource, pick the language you are writing in and fill in the details of your project. Finally, create a manifest.yml file in your project. This will differ depending on what language you are using but for us, using Golang, it's this:

```
---
applications:
  - name: {project_repo)
    buildpack: https://github.com/cloudfoundry/go-buildpack.git
    command: pubapid
    memory: 128M
    disk_quota: 128M
    env:
      GOVERSION: "go1.10"
      GOPACKAGENAME: github.com/{your_github_username}/{project_repo}
      GO_INSTALL_PACKAGE_SPEC: ./cmd/pubapid
      DEPLOY_ENV: production
```

Once that's done, run the command `bx login` in your terminal which will prompt you to login with the credentials from your IBM cloud account and specify what region of the world where you are running the server. When you are logged in, run `bx target -cf` then choose the target space you want to deploy with then run `bx cf push`, this will run the build and deploy your project. This checks that the initial build works. it's always good to know that your project deploys before adding anything new (the basic concept of CD).

Now that your CloudFoundry resource is up and running, we can add the Travis configuration

* We added a `.travis.yml` file to the project that specified what language the project is written in and what version of that language. For us, that was this code at the top of the `.travis.yml` file:

```
language: go

go:
  - "1.10"
```

* added ci check to PR in the setting in integration and services on GitHub for this repo
* added build tag to readme.md file which we got from the travis website by clicking the build(unknown) tag:

#

![screen shot 2018-03-27 at 7 42 48 pm](https://user-images.githubusercontent.com/17281178/38002330-2ad3e798-31f7-11e8-9951-48d7b07bdb97.png)

#

* To `.travis.yml` file we also added the deploy specification. The end result of this for us was:

```
deploy:
  provider: cloudfoundry
  username: apikey
  password: $YOUR_CLOUD_API_KEY
  api: https://api.ng.bluemix.net
  organization: [your cloudfoundry org]
  space: [your space]
```

* for the deploy part of the `travis.yml` file, most of the configurations come from the dashboard application.

#

![screen shot 2018-03-27 at 8 57 38 pm](https://user-images.githubusercontent.com/17281178/38004365-902c9784-3201-11e8-95d6-aff6d057055e.png)

#

Cloud Foundry Space === [your space] and same goes for [your cloudfoundry org]. Username and the api are the same for everyone so you can use what's provided. How you make the password for the deployment is the trickier part. You don't want to put your own password in there for security reasons so here's what you do:

1.  Go to your ibm cloud dashboard for this project.

1.  Hover over manage in the upper right hand corner and hover opver security and the click and then click platform api keys where you will create a secure API key to replace your password in the `.travis.yml` file. Name it whatever you like.

* when you've created the API key, show it and copy it and then run the command `bx login --apikey {YOUR_API_KEY}` where YOUR_API_KEY is that long string of numbers and letters that you got from cloudfoundry.

If i remembered correctly, this should be everything for our version of CI build. Now when a pull request is merged, travis will automatically run the build to deploy the changes to the repository.

**Conclusion:**
Using Travis CI, CI and CD is simple to set up and is likely to save you headaches down the road. I've had projects that I deployed at the end and then have spent literally hours chasing down the problem that caused it, CI and CD protect against that.
