library IEEE;
use IEEE.std_logic_1164.all;

entity kb_addr_decoder_tb is
end entity kb_addr_decoder_tb;

architecture behav of kb_addr_decoder_tb is
  component kb_addr_decoder 
    port (
     e    : in  std_logic;
     a    : in  std_logic_vector(1 downto 0);
     c    : out std_logic_vector(3 downto 0));
  end component;
  for kb_addr_decoder_0: kb_addr_decoder use entity work.kb_addr_decoder;
  signal e : std_logic;
  signal a : std_logic_vector(1 downto 0);
  signal c : std_logic_vector(3 downto 0);
begin
  kb_addr_decoder_0: kb_addr_decoder port map (e => e, a => a, c => c);
  process 
    type pattern_type is record
     e : std_logic;
     a : std_logic_vector(1 downto 0);
     c : std_logic_vector(3 downto 0);
    end record;
    type pattern_array is array (natural range <>) of pattern_type; 
    constant patterns : pattern_array := (
      ('1', "00", "0001"),
      ('1', "01", "0010"),
      ('1', "10", "0100"),
      ('1', "11", "1000"),
      ('1', "X1", "0000"),
      ('1', "1X", "0000"),
      ('1', "XX", "0000"),

      ('0', "00", "0000"),
      ('0', "01", "0000"),
      ('0', "10", "0000"),
      ('0', "11", "0000"),
      ('0', "XX", "0000")
    );
  begin
    for i in patterns'range loop 
      e <= patterns(i).e;
      a <= patterns(i).a;
      wait for 1 ns;
      assert c = patterns(i).c 
        report "bad c value" severity error;
    end loop;
    assert false report "end of test" severity note;
    wait;
  end process;
  
end architecture behav;
