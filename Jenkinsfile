pipeline {
   environment {
      QODANA_TOKEN=credentials('qodana-token')
      SSH_KEY=credentials('ssh-key')
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
            echo '部署项目'
            sh 'export GOARCH=amd64'
            sh 'export GOOS=linux'
            sh '/usr/local/go/bin/go build -o personalMain'
          }
      }
      stage('项目部署至服务器') {
          steps {
            echo '部署项目至服务器'
            sshPublisher(publishers: [sshPublisherDesc(configName: 'JslServer', sshCredentials: [encryptedPassphrase: '{AQAAABAAAAAQSo3d/A9JT3b4GyF3ko9CLF5tctSczhkSZNxEzdL8mEw=}', key: SSH_KEY, keyPath: '', username: 'ecs-user'], transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '''cd ./blog-main
            chmod +x ./personalMain
            nohup ./personalMain > logger.log 2>&1 &''', execTimeout: 120000, flatten: false, keepFilePermissions: true, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: 'blog-main', remoteDirectorySDF: false, removePrefix: '', sourceFiles: './personalMain')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
      }
      stage('项目归档') {
           steps {
              echo '项目归档'
              archiveArtifacts artifacts: './personalMain', fingerprint: true, followSymlinks: false, onlyIfSuccessful: true
          }
      }
   }
}
