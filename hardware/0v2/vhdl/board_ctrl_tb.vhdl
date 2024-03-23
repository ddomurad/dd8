library IEEE;
use IEEE.std_logic_1164.all;
-- use IEEE.numeric_std.all;

entity board_ctrl_tb is
end entity board_ctrl_tb;


architecture behav of board_ctrl_tb is
  component board_ctrl 
    port (
    clk      : in  std_logic;
    rw       : in  std_logic;
    a        : in  std_logic_vector(15 downto 0);
    ce_rom_b : out std_logic;          
    ce_ram_b : out std_logic;          
    ce_uart_b : out std_logic;          
    we_b     : out std_logic;
    oe_b     : out std_logic;
    r0       : out std_logic;
    r1       : out std_logic;
    r2       : out std_logic;
    r3       : out std_logic;
    clk_aux  : out std_logic;
    clk2     : out std_logic
    );
  end component;

  signal clk      : std_logic;
  signal rw       : std_logic;
  signal a        : std_logic_vector(15 downto 0);
  signal ce_rom_b : std_logic;          
  signal ce_ram_b : std_logic;          
  signal ce_uart_b : std_logic;          
  signal we_b     : std_logic;
  signal oe_b     : std_logic;
  signal r0       : std_logic;
  signal r1       : std_logic;
  signal r2       : std_logic;
  signal r3       : std_logic;
  signal clk_aux  : std_logic;
  signal clk2     : std_logic;

  for board_ctrl_0: board_ctrl use entity work.board_ctrl;
begin
  board_ctrl_0: board_ctrl port map(
    clk => clk, rw => rw, a => a, 
    ce_rom_b => ce_rom_b, ce_ram_b => ce_ram_b, 
    we_b => we_b, oe_b => oe_b,
    ce_uart_b => ce_uart_b, r0 => r0, r1 => r1, r2 => r2, r3 => r3,
    clk_aux => clk_aux, clk2 => clk2);

  process 
    type pattern_type is record
      clk      : std_logic;
      rw       : std_logic;
      a        : std_logic_vector(15 downto 0);

      ce_rom_b : std_logic;          
      ce_ram_b : std_logic;          
      ce_uart_b : std_logic;          

      we_b     : std_logic;
      oe_b     : std_logic;
      clk_aux  : std_logic;
      clk2     : std_logic;
    end record;
    type pattern_array is array (natural range <>) of pattern_type; 
    constant patterns : pattern_array := (
      -- ROM
      ('0', '1', "1100000000000000", '0', '1', '1', '1', '1', '0', '0'),
      ('1', '1', "1100000000000000", '0', '1', '1', '1', '0', '1', '1'),
      ('0', '0', "1100000000000000", '0', '1', '1', '1', '1', '0', '0'),
      ('1', '0', "1100000000000000", '0', '1', '1', '0', '1', '1', '1'),

      -- UART
      ('0', '1', "1000000000000000", '1', '1', '0', '1', '1', '0', '0'),
      ('1', '1', "1000000000000000", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '0', "1000000000000000", '1', '1', '0', '1', '1', '0', '0'),
      ('1', '0', "1000000000000000", '1', '1', '0', '0', '1', '1', '1'),

      -- RAM
      ('0', '1', "0000000000000000", '1', '0', '1', '1', '1', '0', '0'),
      ('1', '1', "0000000000000000", '1', '0', '1', '1', '0', '1', '1'),
      ('0', '0', "0000000000000000", '1', '0', '1', '1', '1', '0', '0'),
      ('1', '0', "0000000000000000", '1', '0', '1', '0', '1', '1', '1'),

      ('0', '1', "0100000000000000", '1', '0', '1', '1', '1', '0', '0'),
      ('1', '1', "0100000000000000", '1', '0', '1', '1', '0', '1', '1'),
      ('0', '0', "0100000000000000", '1', '0', '1', '1', '1', '0', '0'),
      ('1', '0', "0100000000000000", '1', '0', '1', '0', '1', '1', '1')

      -- -- clk low, write
      -- ('0', '0', "000", '1', '0', '0', '1', '1', '1'),
      -- ('0', '0', "001", '1', '0', '0', '1', '1', '1'),
      -- ('0', '0', "010", '1', '0', '0', '1', '1', '1'),
      -- ('0', '0', "011", '1', '0', '0', '1', '1', '1'),
      -- ('0', '0', "100", '1', '1', '0', '1', '1', '1'),
      -- ('0', '0', "101", '1', '1', '0', '1', '1', '1'),
      -- ('0', '0', "110", '1', '1', '1', '0', '1', '1'),
      -- ('0', '0', "111", '0', '1', '0', '1', '1', '1'),

      -- -- clk low, read
      -- ('0', '1', "000", '1', '0', '0', '1', '1', '1'),
      -- ('0', '1', "001", '1', '0', '0', '1', '1', '1'),
      -- ('0', '1', "010", '1', '0', '0', '1', '1', '1'),
      -- ('0', '1', "011", '1', '0', '0', '1', '1', '1'),
      -- ('0', '1', "100", '1', '1', '0', '1', '1', '1'),
      -- ('0', '1', "101", '1', '1', '0', '1', '1', '1'),
      -- ('0', '1', "110", '1', '1', '1', '0', '1', '1'),
      -- ('0', '1', "111", '0', '1', '0', '1', '1', '1'),

      -- -- clk high, read
      -- ('1', '1', "000", '1', '0', '0', '1', '1', '0'),
      -- ('1', '1', "001", '1', '0', '0', '1', '1', '0'),
      -- ('1', '1', "010", '1', '0', '0', '1', '1', '0'),
      -- ('1', '1', "011", '1', '0', '0', '1', '1', '0'),
      -- ('1', '1', "100", '1', '1', '0', '1', '1', '0'),
      -- ('1', '1', "101", '1', '1', '0', '1', '1', '0'),
      -- ('1', '1', "110", '1', '1', '1', '0', '1', '0'),
      -- ('1', '1', "111", '0', '1', '0', '1', '1', '0'),

      -- -- clk high, write
      -- ('1', '0', "000", '1', '0', '0', '1', '0', '1'),
      -- ('1', '0', "001", '1', '0', '0', '1', '0', '1'),
      -- ('1', '0', "010", '1', '0', '0', '1', '0', '1'),
      -- ('1', '0', "011", '1', '0', '0', '1', '0', '1'),
      -- ('1', '0', "100", '1', '1', '0', '1', '0', '1'),
      -- ('1', '0', "101", '1', '1', '0', '1', '0', '1'),
      -- ('1', '0', "110", '1', '1', '1', '0', '0', '1'),
      -- ('1', '0', "111", '0', '1', '0', '1', '0', '1')


    );
  begin
    for i in patterns'range loop
      -- report "test:" 
      --   & " rw: " & std_logic'image(patterns(i).rw)
      --   & " clk: " & std_logic'image(patterns(i).clk)
      --   & " a: " & std_logic'image(patterns(i).a(15))
      --     & std_logic'image(patterns(i).a(14))
      --     & std_logic'image(patterns(i).a(13));

      clk <= patterns(i).clk;
      rw <= patterns(i).rw;
      a <= patterns(i).a;
      wait for 1 ns;

      assert ce_rom_b = patterns(i).ce_rom_b
        report "bad ce_rom_b value: " & std_logic'image(ce_rom_b) severity error;
      assert ce_ram_b = patterns(i).ce_ram_b
        report "bad ce_ram_b value" & std_logic'image(ce_ram_b) severity error;
      assert ce_uart_b = patterns(i).ce_uart_b
        report "bad ce_uart_b value" & std_logic'image(ce_uart_b) severity error;
      assert we_b = patterns(i).we_b 
        report "bad we_b value" & std_logic'image(we_b) severity error;
      assert oe_b = patterns(i).oe_b
        report "bad oe_b value" & std_logic'image(oe_b) severity error;
      assert clk_aux = patterns(i).clk_aux
        report "bad clk_aux value" & std_logic'image(clk_aux) severity error;
      assert clk2 = patterns(i).clk2
        report "bad clk2 value" & std_logic'image(clk2) severity error;
    end loop;

    report "end of test" severity note;
    wait;
  end process;
end architecture behav;
