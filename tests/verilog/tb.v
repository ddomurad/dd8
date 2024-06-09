`timescale 1ns / 1ns
`include "and_gate.v" 

module tb;
  reg a,b; 
  wire y;
  
  and_gate dut(.a(a), .b(b), .y(y));


  initial begin 
    $dumpfile("./tb.vcd");
    $dumpvars(0, tb);

    a = 0;
    b = 0;
    #20;


    a = 1;
    b = 0;
    #20;


    a = 0;
    b = 1;
    #20;


    a = 1;
    b = 1;
    #20;

    $display("test completed");
  end

endmodule
