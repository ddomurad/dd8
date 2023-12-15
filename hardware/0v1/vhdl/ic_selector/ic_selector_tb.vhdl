library IEEE;
use IEEE.std_logic_1164.all;
-- use IEEE.numeric_std.all;

entity ic_selector_tb is
end entity ic_selector_tb;


architecture behav of ic_selector_tb is
  component ic_selector 
    port (
      clk    : in std_logic;
      rw     : in std_logic;
      a      : in std_logic_vector(15 downto 13);
      nce_rom : out std_logic; 
      nce_ram : out std_logic; 
      ce_ram : out std_logic; 
      nce_key : out std_logic;
      e_lcd  : out std_logic;
      nwe    : out std_logic;
      noe    : out std_logic
    );
  end component;

  signal clk     : std_logic;
  signal rw      : std_logic;
  signal a       : std_logic_vector(15 downto 13);
  signal nce_rom : std_logic; 
  signal nce_ram : std_logic; 
  signal ce_ram  : std_logic; 
  signal nce_key : std_logic;
  signal e_lcd   : std_logic;
  signal nwe     : std_logic;
  signal noe     : std_logic;

  for ic_selector_0: ic_selector use entity work.ic_selector;
begin
  ic_selector_0: ic_selector port map(
    clk => clk, rw => rw, a => a, 
    nce_rom => nce_rom, nce_ram => nce_ram, 
    ce_ram => ce_ram, nce_key => nce_key, e_lcd => e_lcd, 
    nwe => nwe, noe => noe);

  process 
    type pattern_type is record
       clk     : std_logic;
       rw      : std_logic;
       a       : std_logic_vector(15 downto 13);
       nce_rom : std_logic; 
       nce_ram : std_logic; 
       ce_ram  : std_logic; 
       nce_key : std_logic;
       e_lcd   : std_logic;
       nwe     : std_logic;
       noe     : std_logic;
    end record;
    type pattern_array is array (natural range <>) of pattern_type; 
    constant patterns : pattern_array := (
      -- clk low, write
      ('0', '0', "000", '1', '0', '1', '1', '0', '1', '1'),
      ('0', '0', "001", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '0', "010", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '0', "011", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '0', "100", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '0', "101", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '0', "110", '1', '1', '0', '0', '0', '1', '1'),
      ('0', '0', "111", '0', '1', '0', '1', '0', '1', '1'),

      -- clk low, read
      ('0', '1', "000", '1', '0', '1', '1', '0', '1', '1'),
      ('0', '1', "001", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '1', "010", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '1', "011", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '1', "100", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '1', "101", '1', '1', '0', '1', '0', '1', '1'),
      ('0', '1', "110", '1', '1', '0', '0', '0', '1', '1'),
      ('0', '1', "111", '0', '1', '0', '1', '0', '1', '1'),

      -- clk high, write 
      ('1', '0', "000", '1', '0', '1', '1', '0', '0', '1'),
      ('1', '0', "001", '1', '1', '0', '1', '0', '0', '1'),
      ('1', '0', "010", '1', '1', '0', '1', '0', '0', '1'),
      ('1', '0', "011", '1', '1', '0', '1', '0', '0', '1'),
      ('1', '0', "100", '1', '1', '0', '1', '1', '0', '1'),
      ('1', '0', "101", '1', '1', '0', '1', '0', '0', '1'),
      ('1', '0', "110", '1', '1', '0', '0', '0', '0', '1'),
      ('1', '0', "111", '0', '1', '0', '1', '0', '0', '1'),

      -- clk high, read 
      ('1', '1', "000", '1', '0', '1', '1', '0', '1', '0'),
      ('1', '1', "001", '1', '1', '0', '1', '0', '1', '0'),
      ('1', '1', "010", '1', '1', '0', '1', '0', '1', '0'),
      ('1', '1', "011", '1', '1', '0', '1', '0', '1', '0'),
      ('1', '1', "100", '1', '1', '0', '1', '1', '1', '0'),
      ('1', '1', "101", '1', '1', '0', '1', '0', '1', '0'),
      ('1', '1', "110", '1', '1', '0', '0', '0', '1', '0'),
      ('1', '1', "111", '0', '1', '0', '1', '0', '1', '0')
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

      assert nce_rom = patterns(i).nce_rom 
        report "bad nce_rom value: " & std_logic'image(nce_rom) severity error;
      assert nce_ram = patterns(i).nce_ram 
        report "bad nce_ram value" & std_logic'image(nce_ram) severity error;
      assert nce_key = patterns(i).nce_key 
        report "bad nce_key value" & std_logic'image(nce_key) severity error;
      assert e_lcd = patterns(i).e_lcd 
        report "bad e_lcd value" & std_logic'image(e_lcd) severity error;
      assert nwe = patterns(i).nwe 
        report "bad nwe value" & std_logic'image(nwe) severity error;
      assert noe = patterns(i).noe 
        report "bad noe value" & std_logic'image(noe) severity error;
    end loop;

    report "end of test" severity note;
    wait;
  end process;
end architecture behav;
