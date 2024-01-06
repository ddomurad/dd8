library IEEE;
use IEEE.std_logic_1164.all;
-- use IEEE.numeric_std.all;

entity ic_selector is
  port (
    clk        : in std_logic;
    rw         : in std_logic;
    a          : in std_logic_vector(15 downto 13);
    ce_rom_b   : out std_logic; 
    ce_ram_b   : out std_logic; 
    ce1_uart   : out std_logic; 
    ce2_uart_b : out std_logic; 
    we_b       : out std_logic;
    oe_b       : out std_logic
  );

end entity ic_selector;

architecture rtl of ic_selector is
  signal i_ce_uart : std_logic;
begin
  ce_ram_b <= a(15);
  ce_rom_b <= not(a(15) and a(14) and a(13));
  i_ce_uart <= a(15) and a(14) and not a(13);
  ce1_uart <= i_ce_uart; 
  ce2_uart_b <= not i_ce_uart;
  
  we_b <= not clk or rw;
  oe_b <= not (rw and clk);
end architecture rtl;
