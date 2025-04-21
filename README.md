# WIP: graft
A lightweight distributed log protocol for server clusters. Written in [Go](https://go.dev).

## Use

## TL;DR

## Introduction

### Individual servers
A computer performing the function of a server responds to incoming requests. It uses an internal log/database to determine appropriate response to inputs. As such, the server is a deterministic state machine, generating a predictable output for a given input. 

### Server cluster
In cases where the server performs a crucial function, its availability and ability to correctly handle inputs will, ideally, meet minimum performance specs, which should be monitored. One common mitigation tactic is to run multiple servers as duplicates and, if one or more should fail, the duplicate servers will ensure the integrity of the system.

### Consensus algorithms

Servers operating as a cluster of duplicates need a set of rules to coordinate behavior and guarantee both uptime and data integrity.

The [Paxos](https://paxos.systems/how/) protocol is used ubiquitously; e.g., by [Google Spanner](https://cloud.google.com/spanner). It's built upon theoretical foundations that guarantee minimum performance specs according to cluster configuration alongside the protocol. As a theory-focused protocol, it is considered relatively complicated for non-experts to learn and implement.

In 2014, the [Raft](https://raft.github.io/raft.pdf) protocol was published as an alternative to Paxos. Its design goals were motivated both by performance and by user-friendliness. It closely resembles Paxos, but is (reportedly) easier to learn and understand. A team of non-experts would, presumably, be able to self-manage a Raft cluster with fewer interventions requiring a distributed systems expert. This is a useful practical advantage in production systems for smaller teams. 

The ease of understanding Raft is useful for a lightweight implementation of a consensus mechanism, hence raftlite.

## Distributed logs

Generally, the following considerations are relevant to any consensus protocol for server clusters:

0. accepting and responding to inputs
1. membership and hierarchy of the current core cluster
2. coordination of accurate log duplication among cluster members
3. individual server failures

## Raft's distinguishing features

The original [Raft publication](https://raft.github.io/raft.pdf) introduced the concept of separating the key responsibilities of the distributed log:
1. leader election
2. log replication
3. safety

Two other advances with Raft:
1. require increased coherence: this reduces the number of states that need to be tracked for the algorithm to work robustly as intended
2. cluster membership determination mechanism: overlapping majority (improved safety)

### Leader election

When the current leader server fails*, a new leader is elected. A typical cluster may have five servers. Each of those servers is in one of N states:

0. ok
1. error
2. standby

### Log replication

The leader must accept log entries from clients (input), and replicate those entries across the cluster, simultaneously forcing other cluster servers' logs to match the leader's logs.

### Safety

### Protocol outline

1. elect a distinguished leader (dl)
2. assign complete responsibility for managing the replicated log to the dl:
  i. accepts log entries from clients, 
  ii. replicates them on other servers,
  iii. tells servers when it is safe to apply log entries to their state machines

### Protocol rules

dl powers:
- decide where to place new entries in the log without consulting other servers
- uniquely able to initiate outward flow of information/logs to cluster servers

dl failure:
- new dl elected

log entries only flow from the leader to other servers

randomized timers to elect leaders
- adds minimal overhead as heartbeats already required for any consensus algorithm

joint consensus approach to determine servers in new cluster
- the majorities of two different configurations overlap during transitions, enabling uninterrupted cluster operation


## Paxos vs Raft

[Paxos vs Raft\: Have we reached consensus on distributed consensus\?](https://arxiv.org/abs/2004.05074 "[2004.05074] Paxos vs Raft: Have we reached consensus on distributed consensus?")

"We find that both Paxos and Raft take a very similar approach to distributed consensus, differing only in their approach to leader election. Most notably, Raft only allows servers with up-to-date logs to become leaders, whereas Paxos allows any server to be leader provided it then updates its log to ensure it is up-to-date. Raft’s approach is surprisingly efficient given its simplicity as, unlike Paxos, it does not require log entries to be exchanged during leader election."



