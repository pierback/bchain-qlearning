pragma solidity >=0.4.22<0.6.0;

contract BeverageList {
  struct DrinkData {
    bytes32 drink;
    bytes32 weekday;
  }

  event NewDrink(bytes32 time, bytes32 drink, bytes32 weekday);

  mapping(address => mapping(bytes32 => DrinkData)) private list;
  mapping(address => bytes32[]) timeStamps;

  function setDrinkData(bytes32 time, bytes32 _drink, bytes32 _wd) external {
    list[msg.sender][time] = DrinkData(_drink, _wd);
    timeStamps[msg.sender].push(time);
    emit NewDrink(time, _drink, _wd);
  }

  function getDrinkData(bytes32 time) public view returns (bytes32, bytes32) {
    DrinkData memory data = list[msg.sender][time];
    return (data.drink, data.weekday);
  }

  function lastDrink() public view returns (bytes32, bytes32) {
    bytes32 t = timeStamps[msg.sender][timeStamps[msg.sender].length - 1];
    DrinkData memory data = list[msg.sender][t];
    return (data.drink, data.weekday);
  }
}
