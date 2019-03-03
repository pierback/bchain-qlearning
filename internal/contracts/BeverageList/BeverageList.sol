pragma solidity >=0.4.22<0.6.0;

contract BeverageList {
  struct DrinkData {
    bytes32 drink;
    bytes32 weekday;
  }

  struct DrinkRecord {
    address Address;
    DrinkData dd;
    bytes32 time;
  }

  struct User {
    bool exists;
    bytes32[] timeStamps;
  }

  event NewDrink(address Address, bytes32 time, bytes32 drink, bytes32 weekday);

  mapping(address => mapping(bytes32 => DrinkData)) private usrBvrgHist;

  mapping(address => User) private userList;
  address[] private users;
  DrinkRecord[] private history;

  function setDrinkData(bytes32 time, bytes32 _drink, bytes32 _wd) external {
    if (!userList[msg.sender].exists) {
      userList[msg.sender].exists = true;
      users.push(msg.sender);
    }

    DrinkData memory dd = DrinkData(_drink, _wd);
    usrBvrgHist[msg.sender][time] = dd;

    userList[msg.sender].timeStamps.push(time);
    history.push(DrinkRecord(msg.sender, dd, time));

    emit NewDrink(msg.sender, time, _drink, _wd);
  }

  function getDrinkData(bytes32 time) public view returns (bytes32, bytes32) {
    DrinkData memory data = usrBvrgHist[msg.sender][time];
    return (data.drink, data.weekday);
  }

  function lastUserDrink() public view returns (bytes32, bytes32) {
    bytes32 t = userList[msg.sender].timeStamps[userList[msg.sender].timeStamps.length - 1];
    DrinkData memory data = usrBvrgHist[msg.sender][t];
    return (data.drink, data.weekday);
  }

  function lastDrink() public view returns (address, bytes32, bytes32) {
    DrinkRecord memory dr = history[history.length - 1];
    return (msg.sender, dr.time, dr.dd.drink);
  }

  function userCount() public view returns (uint256) {
    return users.length;
  }

  function getUsers() public view returns (address[] memory) {
    return users;
  }
}
