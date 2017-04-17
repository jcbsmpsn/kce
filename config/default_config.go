package config

const DefaultConfigFileText = `# This is the TOML config file for kce (kubectl-expanded).

#[alias]
#ls = "get"
#ssh = """bash:
#    pod=$(kubectl get pods \
#                  -l run=$1 \
#                  -o custom-columns=:metadata.name | tail -1)
#    if [[ -z $pod ]]; then
#        echo "Could not find a pod for the deployment '$1'"
#        exit 1
#    fi
#
#    exec kubectl exec \
#        -it $pod \
#        bash
#"""
`
