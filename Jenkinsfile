#!groovy
@Library("ace") _

properties([disableConcurrentBuilds()])

def isMaster = "${env.BRANCH_NAME}" == 'master'
def isPR = "${env.CHANGE_URL}".contains('/pull/')

opts = [
    buildAgent: 'jenkins-docker-3',
]

// Hack to prevent ace function from being overloaded during test-deploy
def ace = ace

ace(opts) {
  def goVer = "1.13.1"

  def args = [
    "-v ${pwd()}:/src",
    "-w /src",
    "-e GOCACHE=/tmp/.GOCACHE",
    "-e GOPATH=/src/.GO",
    "-e CI=1",
  ]

  stage("setup") {
    docker.image("golang:${goVer}").inside(args.join(' ')) {
      sh "make setup"
    }
  }

  stage("test") {
    // docker.image("postgres:11").withRun("-e POSTGRES_PASSWORD=secret -e POSTGRES_DB=mydb") { c ->
    //   def testargs = args + [
    //     "--link ${c.id}:db",
    //     "-e PG_PASSWORD=secret",
    //     "-e PG_DB=mydb",
    //     "-e PG_HOST=db",
    //   ]
    def testargs = []

    docker.image("golang:${goVer}").inside(testargs.join(' ')) {
        sh """
        make test
        make cover
        """
    }
    // }
  }

  stage('Build') {
    docker.image("golang:${goVer}").inside(args.join(' ')) {
      sh """
        make
      """
    }

    dockerBuild()
    dockerBuild()
  }

  stage('Push') {
    dockerPush()
  }

  stage('deploy test') {
    deploy('test', [dryrun: isMaster == false])

    if (isMaster) {
      chat.notifyDeploy('test')
    } else {
      chat.notifySuccessful()
    }
  }
}

if (isMaster) {
  waitForInput("Deploy to prod?")

  ace(opts) {
    stage('Deploy to prod') {
      deploy("prod")
      slack.notifyDeploy('prod')
    }
  }
}

