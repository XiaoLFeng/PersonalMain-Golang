pipeline {
   environment {
      QODANA_TOKEN=credentials('qodana-token')
   }
   agent '本地'

   stages {
      stage('Qodana代码检查') {
         when {
            branch 'feature'
         }
         steps {
            script {
               docker.image('jetbrains/qodana-go').run('''
               -v "${WORKSPACE}":/data/project
               --entrypoint=""
               ''') {
                  sh 'qodana'
               }
            }
         }
      }
      stage('项目部署') {
          steps {
            sh 'echo "部署项目"'
            sh 'export GOARCH=amd64'
            sh 'export GOOS=linux'
            sh 'go build -o personalMain'
          }
      }
      stage('项目部署至服务器') {
          steps {
            sh 'echo "部署项目至服务器"'
            // 这里放置部署至服务器的代码
          }
      }
   }
}
