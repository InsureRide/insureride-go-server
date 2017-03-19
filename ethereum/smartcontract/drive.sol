contract mortal {
    /* Define variable owner of the type address*/
    address owner;

    /* this function is executed at initialization and sets the owner of the contract */
    function mortal() { owner = msg.sender; }

    /* Function to recover the funds on the contract */
    function kill() { if (msg.sender == owner) selfdestruct(owner); }
}

contract drive is mortal {

    string public kilometers;
    string public avgspeed;
    string public avgaccel;
    uint32 public starttime;
    uint32 public endtime;

    uint16 public price;


    /* this runs when the contract is executed */
    function drive(string _kilometers, string _avgspeed, string _avgaccel, uint32 _starttime, uint32 _endtime, uint16 _price) public {
        kilometers = _kilometers;
        avgspeed = _avgspeed;
        avgaccel = _avgaccel;
        starttime = _starttime;
        endtime = _endtime;
        price = _price;
    }
}