pipeline {
  agent any
  stages {
    stage('error') {
      parallel {
        stage('error') {
          steps {
            sh 'echo test'
          }
        }

        stage('build') {
          steps {
            sh 'docker build -t test .'
          }
        }

      }
    }

  }
}