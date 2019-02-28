pragma solidity >=0.4.22<0.6.0;

contract BeverageList {
  struct DrinkData {
    bytes32 drink;
    bytes32 weekday;
  }

  event NewDrink(bytes32 time, bytes32 drink, bytes32 weekday);

  mapping(address => mapping(bytes32 => DrinkData)) private list;

  function setDrinkData(bytes32 time, bytes32 _drink, bytes32 _wd) external {
    list[msg.sender][time] = DrinkData(_drink, _wd);
    emit NewDrink(time, _drink, _wd);
  }
}
