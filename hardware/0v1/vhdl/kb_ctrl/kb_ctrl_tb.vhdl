library IEEE;
use IEEE.std_logic_1164.all;


entity kb_ctrl_tb is
end entity kb_ctrl_tb;

architecture behav of kb_ctrl_tb is
  component kb_ctrl 
    port (
      nce, noe, nwe  : in  std_logic;
      a              : in  std_logic_vector(1 downto 0);
      r              : in  std_logic_vector(3 downto 0);
      c              : out std_logic_vector(3 downto 0);
      d_out          : out std_logic_vector(7 downto 0)
    );
  end component;
  for kb_ctrl_0: kb_ctrl use entity work.kb_ctrl;
  signal nce, noe, nwe : std_logic;
  signal d_out  : std_logic_vector(7 downto 0);
  signal a      : std_logic_vector(1 downto 0);
  signal c,r    : std_logic_vector(3 downto 0);
begin
  kb_ctrl_0: kb_ctrl 
    port map (
      nce     => nce,
      noe     => noe,
      nwe     => nwe,
      d_out   => d_out,
      a       => a,
      r       => r,
      c       => c
    );
  process 
    type pattern_type is record 
      nce, noe, nwe  : std_logic;
      a              : std_logic_vector(1 downto 0);
      d_out          : std_logic_vector(7 downto 0);
    end record;
    type pattern_array is array (natural range <>) of pattern_type;
    constant patterns : pattern_array := ( 
      ('0', '0', '0', "00", "ZZZZZZZZ"),
      ('1', '0', '0', "00", "ZZZZZZZZ"),
      ('1', '0', '1', "00", "ZZZZZZZZ"),
      ('0', '1', '0', "00", "ZZZZZZZZ"),
      ('0', '1', '1', "00", "ZZZZZZZZ"),
      ('1', '1', '1', "00", "ZZZZZZZZ"),

      ('0', '0', '1', "00", "00000001"),
      ('0', '0', '1', "01", "00000010"),
      ('0', '0', '1', "10", "00000100"),
      ('0', '0', '1', "11", "00001000")
    );
  begin
    for i in patterns'range loop 
      nce <= patterns(i).nce;
      noe <= patterns(i).noe;
      nwe <= patterns(i).nwe;
      a <= patterns(i).a;
      wait for 500 ps;
      r <= c; -- loop back column to the row
      wait for 500 ps;
      assert d_out = patterns(i).d_out 
        report "bad d_out value in set: " & integer'image(i) severity error;
    end loop;
    assert false report "end of test" severity note;
    wait;
  end process;
end architecture behav;
