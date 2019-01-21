pragma solidity >=0.4.22<0.6.0;
pragma experimental ABIEncoderV2;

contract Qlearning {
  string state;
  bytes32 private learningRate;
  bytes32 private gamma;
  bytes32 private epsilon;
  bytes32 private epsilonDecay;
  uint private successratio;
  uint private prediction;

  mapping(uint => bytes32[2]) public qtable;
  mapping(uint => string) private stateMap;
  //mapping(address => Qlearning) public users;

  constructor(
    string memory _st,
    bytes32 _lr,
    bytes32 _gm,
    bytes32 _ep,
    bytes32 _epd,
    uint _sr,
    uint _pd
  ) public payable {
    state = _st;
    learningRate = _lr;
    gamma = _gm;
    epsilon = _ep;
    epsilonDecay = _epd;
    successratio = _sr;
    prediction = _pd;

    addState(_st, [bytes32(0), bytes32(0)]);
  }

  string[] private states;
  Qlearning private ql;

  /*struct State {
    uint coffeeCount;
    uint mateCount;
    uint waterCount;
    uint weekday;
    uint timeslot;
  }*/

  bytes32[2][] private actnvls;
  function getQtable() public returns (string[] memory, bytes32[2][] memory) {
    bytes32[2][] memory _actnvls = new bytes32[2][](states.length);

    actnvls = _actnvls;
    for (uint i = 0; i < states.length; i++) {
      actnvls.push(qtable[i]);
    }

    return (states, actnvls);
  }

  function setQValue(string memory _st, bytes32[2] memory qvalArr) public {
    int stateIndex = stateExists(_st, states);
    //require(stateIndex > -1);
    qtable[uint(stateIndex)] = qvalArr;
  }

  function addState(string memory st, bytes32[2] memory qvalArr) public {
    //require(stateExists(st, states) < 0);
    states.push(st);
    uint newIndex = states.length - 1;

    qtable[newIndex] = qvalArr;
    stateMap[newIndex] = st;
  }

  function stateExists(string memory st, string[] memory stateArr)
    private
    pure
    returns (int)
  {
    for (uint i = 0; i < stateArr.length; i++) {
      //if (equals(stateArr[i], st)) {
      if (compareStrings(stateArr[i], st)) {
        return int(i);
      }
    }
    return -1;
  }

  function compareStrings(string memory a, string memory b)
    internal
    pure
    returns (bool)
  {
    return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
  }

  /*function equals(State memory st1, State memory st2)
    internal
    pure
    returns (bool)
  {
    return (keccak256(
      abi.encodePacked(
        st1.coffeeCount,
        st1.mateCount,
        st1.waterCount,
        st1.weekday,
        st1.timeslot
      )
    ) == keccak256(
      abi.encodePacked(
        st2.coffeeCount,
        st2.mateCount,
        st2.waterCount,
        st2.weekday,
        st2.timeslot
      )
    ));
  }*/

}
