/*
 * This file makes two assumptions about your Jenkins build environment:
 *
 * 1.  You have nodes set up with labels of the form "docker-${ARCH}" to
 *     support the various build architectures (currently 'x86_64',
 *     's390x', 'aarch64' (ARM), and 'ppc64le').
 * 2.  If you do not want to target the 'docker.io/openwhisk' registry
 *     (and unless you're the official maintainer, you shouldn't), then
 *     you've set up an OPENWHISK_TARGET_REGISTRY variable with the target
 *     registry you'll use.
 *
 * TODO:  Set up the build architectures as a parameter that will drive
 *        a scripted loop to build stages.
 */

def build_shell='''
env | grep JAVA || /bin/true
./gradlew \
  :actionloop:distDocker \
  :golang1.11:distDocker \
  -PdockerRegistry=${OPENWHISK_TARGET_REGISTRY:-docker.io} \
  -PdockerImagePrefix=${OPENWHISK_TARGET_PREFIX:-openwhisk} \
  -PdockerImageTag=latest-$(uname -m)
'''

def manifest_shell='''
registry=${OPENWHISK_TARGET_REGISTRY:-docker.io}
prefix=${OPENWHISK_TARGET_PREFIX:-openwhisk}
rm -rf ~/.docker/manifests
for i in actionloop-v2 actionloop-golang-v1.11; do
  docker manifest create ${registry}/${prefix}/$i:latest \
    ${registry}/${prefix}/$i:latest-x86_64 \
    ${registry}/${prefix}/$i:latest-s390x \
    ${registry}/${prefix}/$i:latest-aarch64 \
    ${registry}/${prefix}/$i:latest-ppc64le
  docker manifest annotate --os linux --arch amd64 \
    ${registry}/${prefix}/$i:latest \
    ${registry}/${prefix}/$i:latest-x86_64
  docker manifest annotate --os linux --arch s390x \
    ${registry}/${prefix}/$i:latest \
    ${registry}/${prefix}/$i:latest-s390x
  docker manifest annotate --os linux --arch arm64 \
    ${registry}/${prefix}/$i:latest \
    ${registry}/${prefix}/$i:latest-aarch64
  docker manifest annotate --os linux --arch ppc64le \
    ${registry}/${prefix}/$i:latest \
    ${registry}/${prefix}/$i:latest-ppc64le
  docker manifest push ${registry}/${prefix}/$i:latest
done
'''

pipeline {
  agent none
  stages {
    stage('Build') {
      parallel {
        stage("Build-x86_64") {
          agent {
            label "docker-x86_64"
          }
          steps {
            sh build_shell
          }
        }
        stage("Build-s390x") {
          agent {
            label "docker-s390x"
          }
          steps {
            sh build_shell
          }
        }
        stage("Build-aarch64") {
          agent {
            label "docker-aarch64"
          }
          steps {
            sh build_shell
          }
        }
        stage("Build-ppc64le") {
          agent {
            label "docker-ppc64le"
          }
          steps {
            sh build_shell
          }
        }
      }
    }
    stage("Manifest") {
      agent {
        // Could be anything capable of running 'docker manifest', but right
        // now only the x86_64 environment is set up for that.
       label "docker-x86_64"
      }
      steps {
        sh manifest_shell
      }
    }
  }
}
