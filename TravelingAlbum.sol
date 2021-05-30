pragma solidity ^0.8.0;
//SPDX-License-Identifier: MIT

import "hardhat/console.sol";
//import "@openzeppelin/contracts/access/Ownable.sol"; //https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol

contract TravelingAlbum {
  constructor() {

    address public prevHolder;
    uint private gasPrice;
    mapping (address => uint) pendingWithdrawls;

    modifier onlyBy(address _account) {
      require(msg.sender == _account,
        "Sender not authorized."
      );
      _;
    }

    constructor() {
      prevHolder = msg.sender;
      gasPrice = msg.value;
      pendingWithdrawls[prevHolder] += 1;
    }
  }

  function setPrevHolder(string memory newHolder) public {
    prevHolder = newOwner;
    console.log(msg.sender,"set prevHolder to",prevHolder);
    emit SetPrevHolder(msg.sender, prevHolder);
  }
}

//contract YourContract {
//
//  event SetPurpose(address sender, string purpose);
//
//  string public purpose = "Building Unstoppable Apps";
//
//  constructor() {
//    // what should we do on deploy?
//  }
//
//  function setPurpose(string memory newPurpose) public {
//    purpose = newPurpose;
//    console.log(msg.sender,"set purpose to",purpose);
//    emit SetPurpose(msg.sender, purpose);
//  }
//
//}
