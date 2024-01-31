pipeline {
   environment {
      QODANA_TOKEN=credentials('qodana-token')
   }
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
      stage('Qodana代码检查') {
         when {
            branch 'feature'
         }
         steps {
            sh '''qodana'''
         }
      }
      stage('项目部署') {
          steps {
            sh '''echo "部署项目"'''
            sh '''set GOARCH=amd64'''
            sh '''set GOOS=linux'''
            sh '''go build -o personalMain'''
          }
      }
      stage('项目部署至服务器') {
          steps {
            sh '''echo "部署项目至服务器"'''
            sshPublisher(publishers: [sshPublisherDesc(configName: 'XiaoLFengBlogServer', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: '''cd ./blog-main
            nohup ./personalMain > logger.log 2>&1 &''', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: 'blog-main', remoteDirectorySDF: false, removePrefix: '', sourceFiles: './personalMain')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: true)])
      }
   }
}