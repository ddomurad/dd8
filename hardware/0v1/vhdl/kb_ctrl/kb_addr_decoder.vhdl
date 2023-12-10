library IEEE;
use IEEE.std_logic_1164.all;

entity kb_addr_decoder is
  port (
   e    : in  std_logic;
   a    : in  std_logic_vector(1 downto 0);
   c    : out std_logic_vector(3 downto 0));
end entity kb_addr_decoder;

architecture rtl of kb_addr_decoder is
begin
  process (e, a)
  begin
    if e = '1' then 
      case a is 
        when "00" => c <= "0001";
        when "01" => c <= "0010";
        when "10" => c <= "0100";
        when "11" => c <= "1000";
        when others => c <= "0000";
      end case;
    else
      c <= "0000";
    end if;
  end process;
end architecture rtl;
