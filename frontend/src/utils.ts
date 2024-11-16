export const formatAddress = (address: string) => {
	return `${address.slice(0, 6)}...${address.slice(-4)}`;
};

export const formatNumber = (number: string | number, compact = false) => {
	return new Intl.NumberFormat("en-US", {
		notation: compact ? "compact" : "standard",
		maximumFractionDigits: 2,
		minimumFractionDigits: 0,
	}).format(number as number);
};

export const OPERATOR_ADDRESS = "0xd185b4846e5fd5419fd4d077dc636084befc51c0";

export const EVC_ADDRESS = "0x0C9a3dd6b8F28529d72d7f9cE918D493519EE383";
