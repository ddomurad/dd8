library IEEE;
use IEEE.std_logic_1164.all;
-- use IEEE.numeric_std.all;

entity board_ctrl is
  port (
    clk       : in  std_logic;
    rw        : in  std_logic;
    a         : in  std_logic_vector(15 downto 0);

    ce_rom_b  : out std_logic;          
    ce_ram_b  : out std_logic;          
    ce_uart_b : out std_logic;          
    we_b      : out std_logic;
    oe_b      : out std_logic;
    r0        : out std_logic;
    r1        : out std_logic;
    r2        : out std_logic;
    r3        : out std_logic;
    clk_aux   : out std_logic;
    clk2      : out std_logic
  );

end entity board_ctrl;

architecture rtl of board_ctrl is
  signal i_ce_uart : std_logic;
begin
    ce_rom_b <= '0'; -- we select only the rom
    ce_ram_b <= '1';
    ce_uart_b <= '1';
    we_b <= '1';  -- we never write in this test
    oe_b <= not (rw and clk);
    r0 <= '0';
    r1 <= '0';
    r2 <= '0';
    r3 <= '0';
    clk_aux <= clk;
    clk2 <= clk;
end architecture rtl;
