contract mortal {
    /* Define variable owner of the type address*/
    address owner;

    /* this function is executed at initialization and sets the owner of the contract */
    function mortal() { owner = msg.sender; }

    /* Function to recover the funds on the contract */
    function kill() { if (msg.sender == owner) selfdestruct(owner); }
}

contract drive is mortal {

    uint64 public kilometers;
    uint16 public avgspeed;
    uint32 public starttime;
    uint32 public endtime;



    /* this runs when the contract is executed */
    function drive(uint64 _kilometers, uint16 _avgspeed, uint32 _starttime, uint32 _endtime) public {
        kilometers = _kilometers;
        avgspeed = _avgspeed;
        starttime = _starttime;
        endtime - _endtime;
    }
}