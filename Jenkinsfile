pipeline {
   environment {
      QODANA_TOKEN=credentials('qodana-token')
   }
   agent {
      label '本地'
      docker {
         args '''
         -v "${WORKSPACE}":/data/project
         --entrypoint=""
         '''
         image 'jetbrains/qodana-go'
      }
   }
   stages {
      stage('Qodana') {
         steps {
            sh '''qodana'''
         }
      }
   }
}