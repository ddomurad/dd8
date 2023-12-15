library IEEE;
use IEEE.std_logic_1164.all;
-- use IEEE.numeric_std.all;

entity ic_selector is
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
end entity ic_selector;

architecture rtl of ic_selector is
  signal i_ce_ram : std_logic;
begin
  i_ce_ram <= (not a(15) and not a(14) and not a(13)); 

  nce_rom <= not (a(15) and a(14) and a(13));
  nce_ram <= not i_ce_ram;
  ce_ram <=  i_ce_ram;
  nce_key <= not (a(15) and a(14) and not a(13));
  e_lcd <= (a(15) and not a(14) and not a(13)) and clk;
  
  nwe <= not clk or rw;
  noe <= not (rw and clk);

end architecture rtl;
