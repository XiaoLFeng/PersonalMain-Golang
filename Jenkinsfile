pipeline {
    environment {
      QODANA_TOKEN=credentials('qodana-token')
   }
    node(本地) {
        println "使用本地辅助节点"
        agent {
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
}