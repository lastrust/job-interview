import React, { Component } from "react";
import FizzBuzzContract from "./contracts/FizzBuzz.json";
import getWeb3 from "./getWeb3";

import "./App.css";

class App extends Component {
  state = { result: "", web3: null, accounts: null, contract: null };

  componentDidMount = async () => {
    try {
      const web3 = await getWeb3();
      const accounts = await web3.eth.getAccounts();
      const networkId = await web3.eth.net.getId();
      const deployedNetwork = FizzBuzzContract.networks[networkId];
      const instance = new web3.eth.Contract(
        FizzBuzzContract.abi,
        deployedNetwork && deployedNetwork.address,
      );

      this.setState({ web3, accounts, contract: instance }, this.runExample);
    } catch (error) {
      // Catch any errors for any of the above operations.
      alert(
        `Failed to load web3, accounts, or contract. Check console for details.`,
      );
      console.error(error);
    }
  };

  callIncrement = async () => {
    const { accounts, contract } = this.state;

    await contract.methods.increment().send({ from: accounts[0] });
  };

  getFizzBuzz = async () => {
    const { contract } = this.state;
    const response = await contract.methods.getFizzBuzz().call();
    this.setState({ result: response });
  };

  render() {
    if (!this.state.web3) {
      return <div>Loading Web3, accounts, and contract...</div>;
    }
    return (
      <div className="App">
        <h1>FizzBuzz Contract</h1>
        <button onClick={this.callIncrement}>increment</button>
        <button onClick={this.getFizzBuzz}>getFizzBuzz</button>
        <p>{this.state.result}</p>
      </div>
    );
  }
}

export default App;
