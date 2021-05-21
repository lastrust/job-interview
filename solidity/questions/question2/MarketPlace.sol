// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.7.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MyToken is ERC721 {
    uint256 private newTokenId;

    constructor() public ERC721("MyToken", "MTK") {}

    function mintToken(address player, string memory tokenURI)
        public
        returns (uint256)
    {
        newTokenId++;
        _mint(player, newTokenId);
        _setTokenURI(newTokenId, tokenURI);
        return newTokenId;
    }
}

contract MarketPlace is Ownable {
    MyToken myToken;
    address owner;
    mapping(address => uint256) revenues;

    constructor() public {
        owner = msg.sender();
        myToken = new MyToken();
    }

    function exhibit(uint256 tokenId) public {
        myToken.delegatecall(
            abi.encodeWithSignature(
                "approve(address,uint256)",
                address(this),
                tokenId
            )
        );
    }

    function buyToken(address seller, uint256 tokenId) payable {
        revenues[seller] += msg.value;
        myToken.transferFrom(seller, msg.sender, tokenId);
    }

    function withdraw(address seller) public onlyOwner {
        uint256 balance = revenues[seller];
        revenues[seller] -= balance;
        seller.transfer(balance);
    }
}
