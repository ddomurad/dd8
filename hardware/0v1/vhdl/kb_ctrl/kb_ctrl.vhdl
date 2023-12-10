library IEEE;
use IEEE.std_logic_1164.all;
-- use IEEE.numeric_std.all;

entity kb_ctrl is 
  port (
    nce   : in  std_logic;
    noe   : in  std_logic;
    nwe   : in  std_logic;
    a     : in  std_logic_vector(1 downto 0);
    r     : in  std_logic_vector(3 downto 0);
    c     : out std_logic_vector(3 downto 0);
    d_out : out std_logic_vector(7 downto 0)
  );
end kb_ctrl;

architecture rtl of kb_ctrl is
  component kb_addr_decoder 
    port (
      e   : in  std_logic;
      a   : in  std_logic_vector(1 downto 0);
      c   : out std_logic_vector(3 downto 0));
  end component;
  signal intenral_enabled : std_logic;
begin
  intenral_enabled <= NOT (nce OR noe OR (NOT nwe)); 
  addr_decoder_unit: kb_addr_decoder port map (e => intenral_enabled, a => a, c => c);
  process (intenral_enabled, r)
  begin 
    if intenral_enabled = '1' then 
      d_out <= "0000" & r;
    else 
      d_out <= "ZZZZZZZZ";
    end if;
  end process;
end architecture rtl;

