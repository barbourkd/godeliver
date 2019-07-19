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
      devices: [],
    }
  },
  methods: {
    async getDevices() {
      try {
        const response = await fetch('http://localhost:8081/devices')
        const data = await response.json()
        this.devices = data
      } catch (error) {
        console.error(error)
      }
    },
    addDevice(device) {
      const lastId =
        this.devices.length > 0
          ? this.devices[this.devices.length - 1 ].id
          : 0 ;
      const id = lastId + 1;
      const newDevice = { ...device, id };

      this.devices = [...this.devices, newDevice]
    }
  },
  mounted() {
    this.getDevices()
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
