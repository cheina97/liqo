#!/usr/bin/env bash
# Copyright 2019-2026 The Liqo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script counts total conditions in FirewallConfigurations and compares
# with the total number of keys in target resources' status.firewallConfigurations maps.

set -e

# Check for bash 4+ (required for associative arrays)
if [ "${BASH_VERSINFO[0]}" -lt 4 ]; then
    echo "Error: This script requires bash 4.0 or higher (current: $BASH_VERSION)"
    echo "On macOS, you may need to install bash via Homebrew: brew install bash"
    exit 1
fi

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo "=========================================="
echo "Firewall Status Migration Verification"
echo "=========================================="
echo ""

# Count FirewallConfigurations with conditions
echo -e "${BLUE}Counting FirewallConfigurations with conditions...${NC}"
fwcfgs=$(kubectl get firewallconfigurations.networking.liqo.io -A -o json)
total_fwcfg_conditions=$(echo "$fwcfgs" | jq '[.items[].status.conditions // [] | length] | add // 0')

# Store FirewallConfiguration details for later comparison
# Key format: namespace/name to match status keys
declare -A fwcfg_map
total_fwcfg_with_conditions=0
while read -r fwcfg; do
    [ -z "$fwcfg" ] && continue
    fwcfg_name=$(echo "$fwcfg" | jq -r '.metadata.name')
    fwcfg_namespace=$(echo "$fwcfg" | jq -r '.metadata.namespace')
    fwcfg_conditions=$(echo "$fwcfg" | jq '.status.conditions // [] | length')
    if [ "$fwcfg_conditions" -gt 0 ]; then
        fwcfg_key="$fwcfg_namespace/$fwcfg_name"
        fwcfg_map["$fwcfg_key"]=$fwcfg_conditions
        total_fwcfg_with_conditions=$((total_fwcfg_with_conditions + 1))
    fi
done < <(echo "$fwcfgs" | jq -c '.items[]')

echo "  FirewallConfigurations with conditions: $total_fwcfg_with_conditions"
echo "  Total conditions across all FirewallConfigurations: $total_fwcfg_conditions"
echo ""

# Count FirewallConfiguration keys in InternalNode status.firewallConfigurations maps
echo -e "${BLUE}Counting InternalNode status.firewallConfigurations keys...${NC}"
internalnodes=$(kubectl get internalnodes.networking.liqo.io -o json 2>/dev/null || echo '{"items":[]}')
declare -A found_fwcfgs
declare -A internalnode_fwcfgs
internalnode_total=0
while read -r node; do
    [ -z "$node" ] && continue
    node_name=$(echo "$node" | jq -r '.metadata.name')
    keys=$(echo "$node" | jq '.status.firewallConfigurations // {} | keys | length')
    if [ "$keys" -gt 0 ]; then
        echo "  InternalNode $node_name: $keys keys"
        internalnode_total=$((internalnode_total + keys))
        # Mark these FirewallConfigurations as found
        while read -r fwcfg_key; do
            [ -z "$fwcfg_key" ] && continue
            found_fwcfgs["$fwcfg_key"]=1
            internalnode_fwcfgs["$fwcfg_key"]=1
        done < <(echo "$node" | jq -r '.status.firewallConfigurations // {} | keys[]')
    fi
done < <(echo "$internalnodes" | jq -c '.items[]')
internalnode_unique=${#internalnode_fwcfgs[@]}
echo "  Total FirewallConfiguration entries in InternalNodes: $internalnode_total (unique: $internalnode_unique)"
echo ""

# Count FirewallConfiguration keys in GatewayClient status.firewallConfigurations maps
echo -e "${BLUE}Counting GatewayClient status.firewallConfigurations keys...${NC}"
gatewayclients=$(kubectl get gatewayclients.networking.liqo.io -A -o json 2>/dev/null || echo '{"items":[]}')
declare -A gatewayclient_fwcfgs
gatewayclient_total=0
while read -r gw; do
    [ -z "$gw" ] && continue
    gw_name=$(echo "$gw" | jq -r '.metadata.name')
    gw_namespace=$(echo "$gw" | jq -r '.metadata.namespace')
    keys=$(echo "$gw" | jq '.status.firewallConfigurations // {} | keys | length')
    if [ "$keys" -gt 0 ]; then
        echo "  GatewayClient $gw_namespace/$gw_name: $keys keys"
        gatewayclient_total=$((gatewayclient_total + keys))
        # Mark these FirewallConfigurations as found
        while read -r fwcfg_key; do
            [ -z "$fwcfg_key" ] && continue
            found_fwcfgs["$fwcfg_key"]=1
            gatewayclient_fwcfgs["$fwcfg_key"]=1
        done < <(echo "$gw" | jq -r '.status.firewallConfigurations // {} | keys[]')
    fi
done < <(echo "$gatewayclients" | jq -c '.items[]')
gatewayclient_unique=${#gatewayclient_fwcfgs[@]}
echo "  Total FirewallConfiguration entries in GatewayClients: $gatewayclient_total (unique: $gatewayclient_unique)"
echo ""

# Count FirewallConfiguration keys in GatewayServer status.firewallConfigurations maps
echo -e "${BLUE}Counting GatewayServer status.firewallConfigurations keys...${NC}"
gatewayservers=$(kubectl get gatewayservers.networking.liqo.io -A -o json 2>/dev/null || echo '{"items":[]}')
declare -A gatewayserver_fwcfgs
gatewayserver_total=0
while read -r gw; do
    [ -z "$gw" ] && continue
    gw_name=$(echo "$gw" | jq -r '.metadata.name')
    gw_namespace=$(echo "$gw" | jq -r '.metadata.namespace')
    keys=$(echo "$gw" | jq '.status.firewallConfigurations // {} | keys | length')
    if [ "$keys" -gt 0 ]; then
        echo "  GatewayServer $gw_namespace/$gw_name: $keys keys"
        gatewayserver_total=$((gatewayserver_total + keys))
        # Mark these FirewallConfigurations as found
        while read -r fwcfg_key; do
            [ -z "$fwcfg_key" ] && continue
            found_fwcfgs["$fwcfg_key"]=1
            gatewayserver_fwcfgs["$fwcfg_key"]=1
        done < <(echo "$gw" | jq -r '.status.firewallConfigurations // {} | keys[]')
    fi
done < <(echo "$gatewayservers" | jq -c '.items[]')
gatewayserver_unique=${#gatewayserver_fwcfgs[@]}
echo "  Total FirewallConfiguration entries in GatewayServers: $gatewayserver_total (unique: $gatewayserver_unique)"
echo ""

# Calculate total FirewallConfiguration entries in target resources
total_in_targets=$((internalnode_total + gatewayclient_total + gatewayserver_total))
total_unique_in_targets=${#found_fwcfgs[@]}

# Print summary
echo "=========================================="
echo "Summary"
echo "=========================================="
echo "FirewallConfigurations with conditions: $total_fwcfg_with_conditions"
echo "  (Note: Total conditions across all: $total_fwcfg_conditions)"
echo ""
echo "FirewallConfiguration entries in target resources:"
echo "  - InternalNode:   $internalnode_total (unique: $internalnode_unique)"
echo "  - GatewayClient:  $gatewayclient_total (unique: $gatewayclient_unique)"
echo "  - GatewayServer:  $gatewayserver_total (unique: $gatewayserver_unique)"
echo -e "  - ${GREEN}TOTAL:          $total_in_targets${NC} (unique: $total_unique_in_targets)"
echo ""

if [ "$total_fwcfg_conditions" -eq "$total_in_targets" ]; then
    echo -e "${GREEN}✓ Perfect match! All FirewallConfiguration conditions are represented in target resources.${NC}"
    exit 0
elif [ "$total_fwcfg_conditions" -eq 0 ]; then
    echo -e "${YELLOW}⚠ No FirewallConfiguration conditions found to verify${NC}"
    exit 0
else
    difference=$((total_fwcfg_conditions - total_in_targets))
    if [ $difference -gt 0 ]; then
        echo -e "${RED}✗ Mismatch: $difference condition(s) not found in target resources${NC}"
    else
        echo -e "${RED}✗ Mismatch: $((difference * -1)) extra entry(ies) in target resources${NC}"
    fi
    echo ""

    # Show detailed information about missing FirewallConfigurations
    echo -e "${YELLOW}Detailed Mismatch Information:${NC}"
    echo ""

    # Find FirewallConfigurations with conditions that are not in target resources
    missing_count=0
    for fwcfg_key in "${!fwcfg_map[@]}"; do
        if [ -z "${found_fwcfgs[$fwcfg_key]}" ]; then
            conditions_count=${fwcfg_map[$fwcfg_key]}
            echo -e "  ${RED}✗${NC} FirewallConfiguration: $fwcfg_key"
            echo "    - Has $conditions_count condition(s) but NOT found in any target resource"
            missing_count=$((missing_count + 1))
        fi
    done

    if [ $missing_count -eq 0 ]; then
        echo -e "  ${YELLOW}Note: All FirewallConfigurations are represented in target resources.${NC}"
        echo -e "  ${YELLOW}The mismatch is in the total count of conditions vs entries.${NC}"
    else
        echo ""
        echo -e "${YELLOW}Total FirewallConfigurations not found in target resources: $missing_count${NC}"
    fi

    exit 1
fi
