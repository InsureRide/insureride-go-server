contract mortal {
    /* Define variable owner of the type address*/
    address owner;

    /* this function is executed at initialization and sets the owner of the contract */
    function mortal() { owner = msg.sender; }

    /* Function to recover the funds on the contract */
    function kill() { if (msg.sender == owner) selfdestruct(owner); }
}

contract car is mortal {

    address owner;
    string public brand;
    string public model;
    uint16 public year;
    string public owneraddress;
    string public vehiclenumber;

    uint public nodrives;
    address[] public drives;


    uint16 public balance;


    /* this runs when the contract is executed */
    function car(string _model, string _brand, uint16 _year, string _owneraddress, string _vehiclenumber) public {
        model = _model;
        brand = _brand;
        year = _year;
        owneraddress = _owneraddress;
        vehiclenumber = _vehiclenumber;
        owner = msg.sender;

        balance = 49900;
    }

    function addDrive(address driveaddress) public {
       if (msg.sender != owner) return;
        drives.push(driveaddress);
        nodrives = drives.length;
    }

    function payInsurance(uint16 amount){
      if (msg.sender != owner) return;
        balance -= amount;
    }
}