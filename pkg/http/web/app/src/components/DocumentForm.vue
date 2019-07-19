<template>
    <div id="document-form">
        <form @submit.prevent="handleSubmit">
            <label>Document name</label>
            <input 
                ref="first"
                type="text"
                :class="{ 'has-error': submitting && invalidName }"
                v-model="device.name"
                @focus="clearStatus"
                @keypress="clearStatus"
            />
            <label>Document content</label>
            <input
                type="text"
                :class="{ 'has-error': submitting && invalidStatus }"
                v-model="device.status"
                @focus="clearStatus"
             />
             <p v-if="error && submitting" class="error-message">
                ❗Please fill out all required fields
             </p>
             <p v-if="success" class="success-message">
                 ✅ Device successfully added
            </p>
            <button>Create Device</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'device-form',
    data() {
        return {
            submitting: false,
            error: false,
            success: false,
            device: {
                name: '',
                status: '',
            },
        }
    },
    methods: {
        handleSubmit() {
            this.submitting = true
            this.clearStatus()

            if (this.invalidName || this.invalidStatus) {
                this.error = true
                return
            }

            this.$emit('add:device', this.device)
            this.$refs.first.focus()
            this.device = {
                name: '',
                status: '',
            }
            this.error = false
            this.success = true
            this.submitting = false
        },

        clearStatus() {
            this.success = false
            this.error = false
        }
    },
    computed: {
        invalidName() {
            return this.device.name === ''
        },

        invalidStatus() {
            return this.device.status === ''
        }
    },
}
</script>

<style scoped>
    form {
        margin-bottom: 2rem;
    }

    [class*='-message'] {
        font-weight: 500;
    }

    .error-message {
        color: #d33c40;
    }

    .success-message {
        color: #32a95d;
    }
</style>

