// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.19;

import "solmate/utils/SafeTransferLib.sol";
import "evc/interfaces/IEthereumVaultConnector.sol";
import "solmate/tokens/ERC4626.sol";
import "openzeppelin/access/Ownable.sol";

/// @title RebalancingOperator
/// @notice This contract allows anyone, in exchange for a tip, to pull liquidity out
/// of a heavily utilised vault on behalf of someone else. Thanks to this operator,
/// a user can delegate the monitoring of their vault to someone else and go on with their life.
contract RebalancingOperator is Ownable {
    using SafeTransferLib for ERC20;

    IEVC public immutable evc;
    ERC20 public immutable asset;

    constructor(address owner, IEVC _evc, ERC20 _asset) Ownable(owner) {
        evc = _evc;
        asset = _asset;
    }

    /// @notice Allows anyone to withdraw on behalf of a onBehalfOfAccount.
    /// @dev Assumes that the onBehalfOfAccount owner had authorized the operator to withdraw on their behalf.
    /// @param from The address of the position to rebalance.
    /// @param to The new target address.
    /// @param onBehalfOfAccount The address of the account on behalf of which the operation is being executed.
    function rebalanceOnBehalf(address from, address to, address onBehalfOfAccount) external onlyOwner {
        uint256 maxWithdraw = ERC4626(from).maxWithdraw(onBehalfOfAccount);

        if (maxWithdraw == 0) return;

        // if there's anything to withdraw, withdraw it to this contract
        evc.call(
            from,
            onBehalfOfAccount,
            0,
            abi.encodeWithSelector(ERC4626.withdraw.selector, maxWithdraw, address(this), onBehalfOfAccount)
        );

        ERC20(asset).approve(to, maxWithdraw);
        ERC4626(to).deposit(maxWithdraw, onBehalfOfAccount);
    }
}
