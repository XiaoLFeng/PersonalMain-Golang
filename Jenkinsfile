pipeline {
   environment {
      QODANA_TOKEN=credentials('qodana-token')
   }
   agent any

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
      stage('项目构建') {
          steps {
            sh 'echo "部署项目"'
            sh 'export GOARCH=amd64'
            sh 'export GOOS=linux'
            sh '/usr/local/go/bin/go build -o personalMain'
          }
      }
      stage('项目部署至服务器') {
          steps {
            sh 'echo "部署项目至服务器"'
            sshPublisher(publishers: [sshPublisherDesc(configName: 'XiaoLFengBlogServer', sshCredentials: [encryptedPassphrase: '{AQAAABAAAAAQrXydf4kWME+M5r1FJeIC7y3P7e+14xuo5oNTQ6wxhrw=}', key: '', keyPath: '/var/jenkins_home/general-key', username: 'ecs-user'], transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '''cd ./blog-main
            chmod +x ./personalMain
            nohup ./personalMain > logger.log 2>&1 &''', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: 'blog-main', remoteDirectorySDF: false, removePrefix: '', sourceFiles: './personalMain')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: true)])
          }
      }
   }
}
