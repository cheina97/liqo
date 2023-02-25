#!/bin/bash

# This scripts expects the following variables to be set:
# CLUSTER_NUMBER        -> the number of liqo clusters
# K8S_VERSION           -> the Kubernetes version
# CNI                   -> the CNI plugin used
# TMPDIR                -> the directory where the test-related files are stored
# BINDIR                -> the directory where the test-related binaries are stored
# TEMPLATE_DIR          -> the directory where to read the cluster templates
# NAMESPACE             -> the namespace where liqo is running
# KUBECONFIGDIR         -> the directory where the kubeconfigs are stored
# LIQO_VERSION          -> the liqo version to test
# INFRA                 -> the Kubernetes provider for the infrastructure
# LIQOCTL               -> the path where liqoctl is stored
# POD_CIDR_OVERLAPPING  -> the pod CIDR of the clusters is overlapping
# CLUSTER_TEMPLATE_FILE -> the file where the cluster template is stored

set -e           # Fail in case of error
set -o nounset   # Fail if undefined variables are used
set -o pipefail  # Fail if one of the piped commands fails

error() {
   local sourcefile=$1
   local lineno=$2
   echo "An error occurred at $sourcefile:$lineno."
}
trap 'error "${BASH_SOURCE}" "${LINENO}"' ERR

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
# shellcheck disable=SC1091
# shellcheck source=./helm-utils.sh
source "${SCRIPT_DIR}/helm-utils.sh"

download_helm

function get_cluster_labels() {
  case $1 in
  1)
  echo "provider=Azure,region=A"
  ;;
  2)
  echo "provider=AWS,region=B"
  ;;
  3)
  echo "provider=GKE,region=C"
  ;;
  4)
  echo "provider=GKE,region=D"
  ;;
  esac
}

LIQO_VERSION="${LIQO_VERSION:-$(git rev-parse HEAD)}"

export SERVICE_CIDR=10.100.0.0/16
export POD_CIDR=10.200.0.0/16
export POD_CIDR_OVERLAPPING=${POD_CIDR_OVERLAPPING:-"false"}

for i in $(seq 1 "${CLUSTER_NUMBER}");
do
  export KUBECONFIG="${TMPDIR}/kubeconfigs/liqo_kubeconf_${i}"
  CLUSTER_LABELS="$(get_cluster_labels "${i}")"
  if [[ ${POD_CIDR_OVERLAPPING} != "true" ]]; then
		# this should avoid the ipam to reserve a pod CIDR of another cluster as local external CIDR causing remapping
		export POD_CIDR="10.$((i * 10)).0.0/16"
	fi
  COMMON_ARGS=(--cluster-name "liqo-${i}" --local-chart-path ./deployments/liqo
    --version "${LIQO_VERSION}" --set controllerManager.config.enableResourceEnforcement=true)
  if [[ "${CLUSTER_LABELS}" != "" ]]; then
    COMMON_ARGS=("${COMMON_ARGS[@]}" --cluster-labels "${CLUSTER_LABELS}" --pod-cidr "${POD_CIDR}" --service-cidr "${SERVICE_CIDR}")
  fi

  curl --fail -LS "https://github.com/liqotech/liqo/releases/download/v0.7.1/liqoctl-linux-amd64.tar.gz" | tar
  chmod 777 ./liqoctl
  mv ./liqoctl "${LIQOCTL}"

  if [ "${i}" == "1" ]; then
    # Install Liqo with Helm, to check that values generation works correctly.
    echo "Using $LIQOCTL to install Liqo on cluster ${i}"
    echo "Installing Liqo with Helm ..."
    echo "Generating values with liqoctl ..."
    "${LIQOCTL}" install "${INFRA}" "${COMMON_ARGS[@]}" --verbose --only-output-values --dump-values-path "${TMPDIR}/values.yaml"
    echo "Applying values with Helm ..."
    "${HELM}" install -n "${NAMESPACE}" --create-namespace liqo ./deployments/liqo -f "${TMPDIR}/values.yaml"
    echo "Applied values with Helm"
  else
    echo "Using $LIQOCTL to install Liqo on cluster ${i}"
    "${LIQOCTL}" install "${INFRA}" "${COMMON_ARGS[@]}" --verbose
  fi
  
done;
