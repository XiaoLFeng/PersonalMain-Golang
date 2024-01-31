pipeline {
    environment {
      QODANA_TOKEN=credentials('qodana-token')
    }
    agent {
        label '本地'
    }

    stages {
        stage('Qodana代码分析') {
            steps {
                script {
                    docker.image('jetbrains/qodana-go').run('''
                    -v "${WORKSPACE}":/data/project
                    -w /data/project
                    -entrypoint ""
                    ''') {
                        sh 'qodana'
                    }
                }
            }
        }
    }
}
