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
    uint8 public year;
    string public owneraddress;
    string public vehiclenumber;


    uint public nodrives;
    address[] public drives;



    /* this runs when the contract is executed */
    function car(string _model, string _brand, uint8 _year, string _owneraddress, string _vehiclenumber) public {
        model = _model;
        brand = _brand;
        year = _year;
        owneraddress = _owneraddress;
        vehiclenumber = _vehiclenumber;
        owner = msg.sender;
    }

    function addDrive(address driveaddress) public {
       if (msg.sender != owner) return;
        drives.push(driveaddress);
        nodrives = drives.length;
    }
}