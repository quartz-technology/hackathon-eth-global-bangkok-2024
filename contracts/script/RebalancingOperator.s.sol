// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.19;

import {Script, console} from "forge-std/Script.sol";
import "evc/interfaces/IEthereumVaultConnector.sol";
import "solmate/tokens/ERC4626.sol";
import {RebalancingOperator} from "../src/RebalancingOperator.sol";

contract RebalancingOperatorScript is Script {
    address public constant BOT = 0xfB79b77fF3Fc4a12795e571Bb91Aac6F35b5C207;
    address public constant EVC = 0x0C9a3dd6b8F28529d72d7f9cE918D493519EE383;
    address public constant USDC = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;

    RebalancingOperator public rebalancingOperator;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        rebalancingOperator = new RebalancingOperator(
            BOT,
            IEVC(EVC),
            ERC20(USDC)
        );

        vm.stopBroadcast();
    }
}
