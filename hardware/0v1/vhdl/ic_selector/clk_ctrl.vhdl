library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.numeric_std.all;


entity clk_ctrl is
  port (
    master_clk : in std_logic;
    strech_selector : in std_logic_vector(1 downto 0);
    clk : out std_logic
  );
end entity clk_ctrl;

architecture rtl of clk_ctrl is
  signal s0_enabled : std_logic;
  signal s0_cnt : unsigned(3 downto 0);
begin
  s0_enabled <= strech_selector(0);
  sa :process(master_clk)
  begin
    if s0_enabled = '0' then
      clk <= master_clk;
      s0_cnt <= "0000";
    else 
      clk <= s0_cnt(2);
      s0_cnt <= s0_cnt + 1;
    end if;
  end process;
end architecture rtl;
