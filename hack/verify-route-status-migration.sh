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

# This script counts total conditions in RouteConfigurations and compares
# with the total number of keys in target resources' status.routeConfigurations maps.

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
echo "Route Status Migration Verification"
echo "=========================================="
echo ""

# Count RouteConfigurations with conditions
echo -e "${BLUE}Counting RouteConfigurations with conditions...${NC}"
rtcfgs=$(kubectl get routeconfigurations.networking.liqo.io -A -o json)
total_rtcfg_conditions=$(echo "$rtcfgs" | jq '[.items[].status.conditions // [] | length] | add // 0')

# Store RouteConfiguration details for later comparison
# Key format: namespace/name to match status keys
declare -A rtcfg_map
total_rtcfg_with_conditions=0
while read -r rtcfg; do
    [ -z "$rtcfg" ] && continue
    rtcfg_name=$(echo "$rtcfg" | jq -r '.metadata.name')
    rtcfg_namespace=$(echo "$rtcfg" | jq -r '.metadata.namespace')
    rtcfg_conditions=$(echo "$rtcfg" | jq '.status.conditions // [] | length')
    if [ "$rtcfg_conditions" -gt 0 ]; then
        rtcfg_key="$rtcfg_namespace/$rtcfg_name"
        rtcfg_map["$rtcfg_key"]=$rtcfg_conditions
        total_rtcfg_with_conditions=$((total_rtcfg_with_conditions + 1))
    fi
done < <(echo "$rtcfgs" | jq -c '.items[]')

echo "  RouteConfigurations with conditions: $total_rtcfg_with_conditions"
echo "  Total conditions across all RouteConfigurations: $total_rtcfg_conditions"
echo ""

# Count RouteConfiguration keys in InternalNode status.routeConfigurations maps
echo -e "${BLUE}Counting InternalNode status.routeConfigurations keys...${NC}"
internalnodes=$(kubectl get internalnodes.networking.liqo.io -o json 2>/dev/null || echo '{"items":[]}')
declare -A found_rtcfgs
declare -A internalnode_rtcfgs
internalnode_total=0
while read -r node; do
    [ -z "$node" ] && continue
    node_name=$(echo "$node" | jq -r '.metadata.name')
    keys=$(echo "$node" | jq '.status.routeConfigurations // {} | keys | length')
    if [ "$keys" -gt 0 ]; then
        echo "  InternalNode $node_name: $keys keys"
        internalnode_total=$((internalnode_total + keys))
        # Mark these RouteConfigurations as found
        while read -r rtcfg_key; do
            [ -z "$rtcfg_key" ] && continue
            found_rtcfgs["$rtcfg_key"]=1
            internalnode_rtcfgs["$rtcfg_key"]=1
        done < <(echo "$node" | jq -r '.status.routeConfigurations // {} | keys[]')
    fi
done < <(echo "$internalnodes" | jq -c '.items[]')
internalnode_unique=${#internalnode_rtcfgs[@]}
echo "  Total RouteConfiguration entries in InternalNodes: $internalnode_total (unique: $internalnode_unique)"
echo ""

# Count RouteConfiguration keys in GatewayClient status.routeConfigurations maps
echo -e "${BLUE}Counting GatewayClient status.routeConfigurations keys...${NC}"
gatewayclients=$(kubectl get gatewayclients.networking.liqo.io -A -o json 2>/dev/null || echo '{"items":[]}')
declare -A gatewayclient_rtcfgs
gatewayclient_total=0
while read -r gw; do
    [ -z "$gw" ] && continue
    gw_name=$(echo "$gw" | jq -r '.metadata.name')
    gw_namespace=$(echo "$gw" | jq -r '.metadata.namespace')
    keys=$(echo "$gw" | jq '.status.routeConfigurations // {} | keys | length')
    if [ "$keys" -gt 0 ]; then
        echo "  GatewayClient $gw_namespace/$gw_name: $keys keys"
        gatewayclient_total=$((gatewayclient_total + keys))
        # Mark these RouteConfigurations as found
        while read -r rtcfg_key; do
            [ -z "$rtcfg_key" ] && continue
            found_rtcfgs["$rtcfg_key"]=1
            gatewayclient_rtcfgs["$rtcfg_key"]=1
        done < <(echo "$gw" | jq -r '.status.routeConfigurations // {} | keys[]')
    fi
done < <(echo "$gatewayclients" | jq -c '.items[]')
gatewayclient_unique=${#gatewayclient_rtcfgs[@]}
echo "  Total RouteConfiguration entries in GatewayClients: $gatewayclient_total (unique: $gatewayclient_unique)"
echo ""

# Count RouteConfiguration keys in GatewayServer status.routeConfigurations maps
echo -e "${BLUE}Counting GatewayServer status.routeConfigurations keys...${NC}"
gatewayservers=$(kubectl get gatewayservers.networking.liqo.io -A -o json 2>/dev/null || echo '{"items":[]}')
declare -A gatewayserver_rtcfgs
gatewayserver_total=0
while read -r gw; do
    [ -z "$gw" ] && continue
    gw_name=$(echo "$gw" | jq -r '.metadata.name')
    gw_namespace=$(echo "$gw" | jq -r '.metadata.namespace')
    keys=$(echo "$gw" | jq '.status.routeConfigurations // {} | keys | length')
    if [ "$keys" -gt 0 ]; then
        echo "  GatewayServer $gw_namespace/$gw_name: $keys keys"
        gatewayserver_total=$((gatewayserver_total + keys))
        # Mark these RouteConfigurations as found
        while read -r rtcfg_key; do
            [ -z "$rtcfg_key" ] && continue
            found_rtcfgs["$rtcfg_key"]=1
            gatewayserver_rtcfgs["$rtcfg_key"]=1
        done < <(echo "$gw" | jq -r '.status.routeConfigurations // {} | keys[]')
    fi
done < <(echo "$gatewayservers" | jq -c '.items[]')
gatewayserver_unique=${#gatewayserver_rtcfgs[@]}
echo "  Total RouteConfiguration entries in GatewayServers: $gatewayserver_total (unique: $gatewayserver_unique)"
echo ""

# Calculate total RouteConfiguration entries in target resources
total_in_targets=$((internalnode_total + gatewayclient_total + gatewayserver_total))
total_unique_in_targets=${#found_rtcfgs[@]}

# Print summary
echo "=========================================="
echo "Summary"
echo "=========================================="
echo "RouteConfigurations with conditions: $total_rtcfg_with_conditions"
echo "  (Note: Total conditions across all: $total_rtcfg_conditions)"
echo ""
echo "RouteConfiguration entries in target resources:"
echo "  - InternalNode:   $internalnode_total (unique: $internalnode_unique)"
echo "  - GatewayClient:  $gatewayclient_total (unique: $gatewayclient_unique)"
echo "  - GatewayServer:  $gatewayserver_total (unique: $gatewayserver_unique)"
echo -e "  - ${GREEN}TOTAL:          $total_in_targets${NC} (unique: $total_unique_in_targets)"
echo ""

if [ "$total_rtcfg_conditions" -eq "$total_in_targets" ]; then
    echo -e "${GREEN}✓ Perfect match! All RouteConfiguration conditions are represented in target resources.${NC}"
    exit 0
elif [ "$total_rtcfg_conditions" -eq 0 ]; then
    echo -e "${YELLOW}⚠ No RouteConfiguration conditions found to verify${NC}"
    exit 0
else
    difference=$((total_rtcfg_conditions - total_in_targets))
    if [ $difference -gt 0 ]; then
        echo -e "${RED}✗ Mismatch: $difference condition(s) not found in target resources${NC}"
    else
        echo -e "${RED}✗ Mismatch: $((difference * -1)) extra entry(ies) in target resources${NC}"
    fi
    echo ""

    # Show detailed information about missing RouteConfigurations
    echo -e "${YELLOW}Detailed Mismatch Information:${NC}"
    echo ""

    # Find RouteConfigurations with conditions that are not in target resources
    missing_count=0
    for rtcfg_key in "${!rtcfg_map[@]}"; do
        if [ -z "${found_rtcfgs[$rtcfg_key]}" ]; then
            conditions_count=${rtcfg_map[$rtcfg_key]}
            echo -e "  ${RED}✗${NC} RouteConfiguration: $rtcfg_key"
            echo "    - Has $conditions_count condition(s) but NOT found in any target resource"
            missing_count=$((missing_count + 1))
        fi
    done

    if [ $missing_count -eq 0 ]; then
        echo -e "  ${YELLOW}Note: All RouteConfigurations are represented in target resources.${NC}"
        echo -e "  ${YELLOW}The mismatch is in the total count of conditions vs entries.${NC}"
    else
        echo ""
        echo -e "${YELLOW}Total RouteConfigurations not found in target resources: $missing_count${NC}"
    fi

    exit 1
fi
