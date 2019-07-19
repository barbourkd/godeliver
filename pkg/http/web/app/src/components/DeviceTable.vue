<template>
    <div id="device-table">
        <table>
            <thead>
                <tr>
                    <th>Device Name</th>
                    <th>Device Status</th>
                    <th>Enable/Disable</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="device in devices" :key="device.id">
                    <td>{{ device.name }}</td>
                    <td>{{ device.status }}</td>
                    <td>
                        <button v-if="device.status == 'Inactive'" class="enable-button">Enable</button>
                        <button v-if="device.status == 'Active'" class="disable-button">Disable</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
export default {
    name: 'device-table',
    props: {
        devices: Array,
    }
}
</script>

<style scoped>
    .disable-button {
        background-color: red;
        border-color: darkred;
    }
</style>

