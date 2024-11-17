// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.19;

import "solmate/utils/SafeTransferLib.sol";
import "evc/interfaces/IEthereumVaultConnector.sol";
import "solmate/tokens/ERC4626.sol";
import "openzeppelin/access/Ownable.sol";

/// @title RebalancingOperator
/// @notice This contract allows anyone, to set an operator that can rebalance the positions to track better yields.
contract RebalancingOperator is Ownable {
    using SafeTransferLib for ERC20;

    event Rebalance(address indexed from, address indexed to, address indexed onBehalfOfAccount, uint256 amount);

    IEVC public immutable evc;
    ERC20 public immutable asset;

    constructor(address owner, IEVC _evc, ERC20 _asset) Ownable(owner) {
        evc = _evc;
        asset = _asset;
    }

    /// @notice Allows operator to rebalance a position from one vault to another.
    /// @param from The vault address of the position to rebalance.
    /// @param to The new vault address.
    /// @param onBehalfOfAccount The address of the account on behalf of which the operation is being executed.
    function rebalanceOnBehalf(address from, address to, address onBehalfOfAccount) external onlyOwner {
        uint256 maxWithdraw = ERC4626(from).maxWithdraw(onBehalfOfAccount);

        if (maxWithdraw == 0) return;

        evc.call(
            from,
            onBehalfOfAccount,
            0,
            abi.encodeWithSelector(ERC4626.withdraw.selector, maxWithdraw, address(this), onBehalfOfAccount)
        );

        ERC20(asset).approve(to, type(uint256).max);
        ERC4626(to).deposit(maxWithdraw, onBehalfOfAccount);

        emit Rebalance(from, to, onBehalfOfAccount, maxWithdraw);
    }
}
