var FizzBuzz = artifacts.require("./FizzBuzz.sol");

module.exports = function(deployer) {
  deployer.deploy(FizzBuzz);
};
