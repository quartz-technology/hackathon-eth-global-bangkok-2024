import { log } from "@graphprotocol/graph-ts";
import { Rebalance as RebalanceEvent } from "../generated/RebalancingOperator/RebalancingOperator";
import { Rebalance as RebalanceEntity } from "../generated/schema";

export function handleRebalance(rebalanceEvent: RebalanceEvent): void {
	let rebalanceEntity = new RebalanceEntity(
		rebalanceEvent.transaction.hash.concatI32(rebalanceEvent.logIndex.toI32()),
	);

	rebalanceEntity.from = rebalanceEvent.params.from;
	rebalanceEntity.to = rebalanceEvent.params.to;
	rebalanceEntity.amount = rebalanceEvent.params.amount;
	rebalanceEntity.onBehalfOfAccount = rebalanceEvent.params.onBehalfOfAccount;
	rebalanceEntity.blockNumber = rebalanceEvent.block.number;
	rebalanceEntity.blockTimestamp = rebalanceEvent.block.timestamp;
	rebalanceEntity.transactionHash = rebalanceEvent.transaction.hash;

	rebalanceEntity.save();

	log.info("Rebalance event: {}", [rebalanceEvent.transaction.hash.toHex()]);
}
