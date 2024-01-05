library IEEE;
use IEEE.std_logic_1164.all;

entity clk_ctrl_tb is
end entity clk_ctrl_tb;


architecture behav of clk_ctrl_tb is
  component clk_ctrl 
    port (
      master_clk : in std_logic; 
      strech_selector : in std_logic_vector (1 downto 0);
      clk : out std_logic
   );
  end component;


  signal master_clk : std_logic;
  signal strech_selector : std_logic_vector(1 downto 0);
  signal s_clk : std_logic;

  for dut: clk_ctrl use entity work.clk_ctrl;

  constant mclk_period : time := 1 ns;
begin
  dut: clk_ctrl port map(
    master_clk => master_clk,
    strech_selector => strech_selector,
    clk => s_clk);

  master_clk_process :process 
  begin 
    master_clk <= '0';
    wait for mclk_period/2;
    master_clk <= '1';
    wait for mclk_period/2;
  end process;

  test_process :process 
  begin 
    strech_selector <= "00";
    no_streach_test: for i in 0 to 10 loop 
      wait until master_clk = '0';
      wait for 1 ps;
      assert s_clk = '0' 
        report "bad s_clk no straching" severity error;
      wait until master_clk = '1';
      wait for 1 ps;
      assert s_clk = '1' 
        report "bad s_clk no straching" severity error;
    end loop;

    wait until master_clk = '0';
    strech_selector <= "01";
    wait for 5 ns;
    strech_selector <= "00";
    wait;

  end process;

end architecture behav;
