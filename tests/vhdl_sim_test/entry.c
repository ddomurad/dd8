#include "vpi_user.h"
#include <stdio.h>
#include <string.h>
#include <math.h>

vpiHandle clk_net, a_net, rom_net, ram_net, d_net, ram_data_net;
int sim_iters = 5;

static s_vpi_time double_to_vpi_time(double t, double time_resolution)
{
	s_vpi_time ts;
	uint64_t simtime = (uint64_t)(t * time_resolution);

	ts.type = vpiSimTime;
	ts.low = (uint32_t)(simtime & 0xffffffffUL);
	ts.high = (uint32_t)(simtime >> 32);

	return ts;
}

static void register_cb_after(PLI_INT32 (*cb_rtn)(struct t_cb_data *), double delay) {
	s_cb_data cb;
	vpiHandle callback_handle;
	double time_resolution = pow(10, -vpi_get(vpiTimePrecision, NULL));
	s_vpi_time time = double_to_vpi_time(delay, time_resolution);

	cb.reason = cbAfterDelay;
	cb.cb_rtn = cb_rtn;
	cb.obj = NULL;
	cb.time = &time;
	cb.value = NULL;
	cb.user_data = NULL;

	callback_handle = vpi_register_cb(&cb);
	if (!callback_handle) {
		vpi_printf("\e[31mERROR: Cannot register cbAfterDelay\e[0m\n");
	}
}


static PLI_INT32 iter_sim(p_cb_data cb_data) {
  vpi_printf("sim iteration\n");

	s_vpi_value clk_val;
	clk_val.format = vpiIntVal;
	clk_val.value.integer = sim_iters%2;
  vpi_put_value(clk_net, &clk_val, NULL, vpiNoDelay);

  s_vpi_value d_val;
  d_val.format = vpiIntVal;
  d_val.value.integer = 0xaa;
  vpi_put_value(d_net, &d_val, NULL, vpiNoDelay);

  s_vpi_value a_val;
  a_val.format = vpiIntVal;
  a_val.value.integer = 0xffff;
  vpi_put_value(a_net, &a_val, NULL, vpiNoDelay);

  s_vpi_value ram_val;
  ram_val.format = vpiIntVal;

  s_vpi_value rom_val;
  rom_val.format = vpiIntVal;

  // s_vpi_arrayvalue ram_data_val;
  s_vpi_value ram_data_val;
  ram_data_val.format = vpiIntVal;

  vpi_get_value(clk_net, &clk_val);
  vpi_get_value(rom_net, &rom_val);
  vpi_get_value(ram_net, &ram_val);
  vpi_get_value(a_net, &a_val);
  vpi_get_value(d_net, &d_val);
  // vpi_get_value_array(ram_data_net, &ram_data_val, &idx, 100);
  vpi_get_value(ram_data_net, &ram_data_val);

  vpi_printf("clk_val: %d\n", clk_val.value.integer);
  vpi_printf("rom_val: %d\n", rom_val.value.integer);
  vpi_printf("ram_val: %d\n", ram_val.value.integer);
  vpi_printf("a_val: %d\n", a_val.value.integer);
  vpi_printf("d_val: %d\n", d_val.value.integer);
  // vpi_printf("ram[0xffff]: %d\n", ram_data_val.value.vector[10]);
  // vpi_printf("ram[0xffff]: %d\n", ram_data_val.value.integer);

  // vpi_printf("ram_data_val: %d\n", ram_data_val.value.integer);

  if (++sim_iters > 10) {
    return 0;
  }
  register_cb_after(iter_sim, 1); 
  return 0;
}

void iterate_stuff(vpiHandle parrentMod) {
  vpiHandle mod_iter, iter, mod, net;
  int net_w, net_d;

  mod_iter = vpi_iterate(vpiModule, parrentMod);
  if (!iter) {
    vpi_printf("ERR: no iter!\n");
    return;
  }

  const char *name;
  while ((mod = vpi_scan(mod_iter))) {
    name = vpi_get_str(vpiName, mod);
    vpi_printf("mod name: %s\n", name);

    iter = vpi_iterate(vpiNet, mod);
    if (!iter) {
      vpi_printf("ERR: no net iter for mod: %s!\n", name);
      return;
    }

    while (1) {
      net = vpi_scan(iter);
      if (!net) {
        break;
      }
      name = vpi_get_str(vpiName, net);
      net_w = vpi_get(vpiSize, net);
      net_d = vpi_get(vpiDirection, net);
      vpi_printf("\tnet name: %s, %d, %d\n", name, net_w, net_d);
      if(strcmp("clk", name) == 0) {
        printf("found clk\n");
        clk_net = net;
      }
      if(strcmp("a", name) == 0) {
        printf("found a\n");
        a_net = net;
      }
      if(strcmp("ce_rom_b", name) == 0) {
        printf("found ce_rom_b\n");
        rom_net = net;
      }
      if(strcmp("ce_ram_b", name) == 0) {
        printf("found ce_ram_b\n");
        ram_net = net;
      }
      if(strcmp("d", name) == 0) {
        printf("found d\n");
        d_net = net;

        ram_net = vpi_handle_by_name("ram", mod);
        if (ram_net == NULL) {
          printf("just null\n");
        }
      }
      // break;
      // vpi_free_object(net);
    }

    // vpi_free_object(iter);
    // vpi_free_object(mod);
    vpi_printf("lil\n");
    iterate_stuff(mod);
  }
}

void iterate_range(vpiHandle hdl) {
  const char *hdl_name = vpi_get_str(vpiName, hdl);
  printf("it_range name: %s\n", hdl_name);

  vpiHandle iter = vpi_iterate(vpiRange,  hdl);
  if(iter == NULL) {
    printf("iter is null\n");
    return;
  }

  vpiHandle rngH;
  while((rngH = vpi_scan(iter)) != NULL) {

  }
}

PLI_INT32 cb_sim_start() {
  printf("Hello!\n");
  iterate_stuff(NULL);
  register_cb_after(iter_sim, 1); 
  return 0;
}

static PLI_INT32 end_of_sim_cb(p_cb_data cb_data)
{
	vpi_printf("Simulation ended.\n");
	return 0;
}

void entry_point_cb() {
  s_cb_data cb;

  cb.reason = cbStartOfSimulation;
  cb.cb_rtn = &cb_sim_start;
  cb.user_data = NULL;

  if (vpi_register_cb(&cb) == NULL) {
    vpi_printf("cannot register cbStartOfSimulation call back\n");
  }

	/* Register end of simulation callback */
	cb.reason = cbEndOfSimulation;
	cb.cb_rtn = &end_of_sim_cb;
	cb.user_data = NULL;

	 vpi_register_cb(&cb);
}

// List of entry points called when the plugin is loaded
void (*vlog_startup_routines[])() = {entry_point_cb, 0};
