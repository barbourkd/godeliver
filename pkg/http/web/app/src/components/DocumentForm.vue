<template>
    <div id="document-form">
        <form @submit.prevent="handleSubmit">
            <label>Document name</label>
            <input 
                ref="first"
                type="text"
                :class="{ 'has-error': submitting && invalidName }"
                v-model="document.name"
                @focus="clearStatus"
                @keypress="clearStatus"
            />
            <label>Document content</label>
            <input
                type="text"
                :class="{ 'has-error': submitting && invalidStatus }"
                v-model="document.content"
                @focus="clearStatus"
             />
             <p v-if="error && submitting" class="error-message">
                ❗Please fill out all required fields
             </p>
             <p v-if="success" class="success-message">
                 ✅ Document submitted to queue
            </p>
            <button>Submit Document</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'document-form',
    data() {
        return {
            submitting: false,
            error: false,
            success: false,
            document: {
                name: '',
                content: '',
            },
        }
    },
    methods: {
        handleSubmit() {
            this.submitting = true
            this.clearStatus()

            if (this.invalidName) {
                this.error = true
                return
            }

            this.$emit('submit:document', this.document)
            this.$refs.first.focus()
            this.document = {
                name: '',
                content: '',
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
            return this.document.name === ''
        },
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

