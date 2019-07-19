<template>
  <div id="app" class="small-container">
    <h1>Devices</h1>
    <device-form @add:device="addDevice" />

    <device-table :devices="devices"/>
  </div>
</template>

<script>
import DeviceTable from './components/DeviceTable.vue'
import DeviceForm from './components/DeviceForm.vue'

export default {
  name: 'app',
  components: {
    DeviceTable,
    DeviceForm
  },
  data() {
    return {
      devices: [
        {
          id: 1,
          name: "Email Device",
          status: "Active"
        },
        {
          id: 2,
          name: "Print Device",
          status: "Inactive"
        },
      ]
    }
  },
  methods: {
    addDevice(device) {
      const lastId =
        this.devices.length > 0
          ? this.devices[this.devices.length - 1 ].id
          : 0 ;
      const id = lastId + 1;
      const newDevice = { ...device, id };

      this.devices = [...this.devices, newDevice]
    }
  }
}
</script>

<style>
  button {
    background: #009435;
    border: 1px solid #009435;
  }

  .small-container {
    max-width: 680px;
  }
</style>
