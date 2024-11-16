// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.19;

import "../src/RebalancingOperator.sol";

import "solmate/tokens/ERC4626.sol";

import "forge-std/console.sol";
import "forge-std/Test.sol";

contract RebalancingMainnetTest is Test {
    function setUp() public {
        // fork
        vm.createSelectFork(vm.envString("ETH_RPC_URL"));
    }

    address constant USDC_WHALE = 0x37305B1cD40574E4C5Ce33f8e8306Be057fD7341;
    address constant USDC = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;

    address constant EVC = 0x0C9a3dd6b8F28529d72d7f9cE918D493519EE383;
    address constant USDC_VAULT_1 = 0x797DD80692c3b2dAdabCe8e30C07fDE5307D48a9;
    address constant USDC_VAULT_2 = 0xce45EF0414dE3516cAF1BCf937bF7F2Cf67873De;

    function testRebalancingOperator() public {
        address admin = vm.addr(1);
        address testUser = vm.addr(2);

        vm.label(address(this), "TEST");
        vm.label(EVC, "RebalancingOperator");
        vm.label(USDC_VAULT_1, "USDC_VAULT_1");
        vm.label(USDC_VAULT_2, "USDC_VAULT_2");
        vm.label(testUser, "test user");

        IEVC evc = IEVC(EVC);

        // deploy RebalancingOperator
        RebalancingOperator rebalancingOperator = new RebalancingOperator(admin, evc, ERC20(USDC));

        vm.prank(USDC_WHALE);
        ERC20(USDC).transfer(testUser, 1000e6);

        vm.startPrank(testUser);
        IEVC(evc).setAccountOperator(testUser, address(rebalancingOperator), true);
        ERC20(USDC).approve(address(rebalancingOperator), type(uint256).max);

        console.log("RebalancingOperator deployed at: ", address(rebalancingOperator));

        ERC20(USDC).approve(USDC_VAULT_1, 100e6);
        ERC4626(USDC_VAULT_1).deposit(100e6, testUser);

        vm.assertApproxEqAbs(ERC4626(USDC_VAULT_1).convertToAssets(ERC20(USDC_VAULT_1).balanceOf(testUser)), 100e6, 10);

        // rebalance
        vm.stopPrank();
        vm.prank(admin);
        rebalancingOperator.rebalanceOnBehalf(USDC_VAULT_1, USDC_VAULT_2, testUser);

        vm.assertApproxEqAbs(ERC4626(USDC_VAULT_2).convertToAssets(ERC20(USDC_VAULT_2).balanceOf(testUser)), 100e6, 10);
        vm.assertEq(ERC20(USDC_VAULT_1).balanceOf(testUser), 0);
    }
}
