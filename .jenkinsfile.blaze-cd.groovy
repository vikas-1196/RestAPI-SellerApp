lock(label: 'hpse-build-pipeline', quantity: 1) {
    withCredentials([usernamePassword(credentialsId: 'ARTIFACTORY_PROD', passwordVariable: 'TF_VAR_ARTIFACTORY_KEY', usernameVariable: 'TF_VAR_ARTIFACTORY_USER'),
            usernamePassword(credentialsId: 'IBMCLOUD_API_KEY_STAGING', passwordVariable: 'TF_VAR_ibmcloud_api_key', usernameVariable: 'DOCKER_IMAGE_ENV')]) {
        withEnv(['TF_VAR_ARTIFACT_ID=hpse-pipeline-phase2-dev-build-vm']) {
            bzCD(
                pushLatestTag: true,
                verbose: false,
                architecture: 'x86_64',
                promoteImages: false,
                useAnalysisSonarQube: true,
                sonarQube: [
                        useCIO: true,
                        failOnQualityGate: true
                ],
                native: [
                    kind: 'make',
                    mavenResolver: [
                        [
                            repository: 'sys-zaas-team-hpse-dev-release-maven-local',
                            artifacts: [
                                'base-image.qcow2': 'com.ibm.zaas.cicd:hpse-pipeline-bootstrap-22-04-dev-build-vm:25.4.5@qcow2',
                                'hpse-cidata-jwt-dev-key.pem': 'com.ibm.zaas.zaas:hpse-cidata-jwt-dev-key-deb:25.3.0:privkey-build@pem'
                            ]
                        ]
                    ],
                    contacts: [
                        "sativar.sainath@ibm.com": [
                            moniker: "Sainath Sativar",
                            github: "Sativar-sainath",
                            slack: "sativar.sainath",
                            roles: ["notify", "owner"]
                        ],
                        "ankis021@in.ibm.com": [
                            moniker: "Ankitha S",
                            github: "ankis021",
                            slack: "Ankitha S ",
                            roles: ["notify", "owner"]
                        ]
                    ],
                    docker: false,
                    helm: false
                ])
        }
    }
}
