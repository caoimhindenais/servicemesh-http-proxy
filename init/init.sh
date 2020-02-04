#!/bin/bash
# Forward all inbound TCP traffic for 8080 on to 8000
# The proxy itself forwards the traffic on
iptables -t nat -A PREROUTING -p tcp -i eth0 --dport 8080 -j REDIRECT --to-port 8000

# List all iptables rules.
#iptables -t nat --list
iptables -t nat -L --line-numbers -n