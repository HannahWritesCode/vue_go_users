<template>
    <div class="min-w-full">
        <h1 class="text-4xl font-semibold pb-3">You made it!</h1>

        <form @submit.stop.prevent="attemptVerification()">
            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Verify your account" id="submit-button" />
        </form>

        <div v-if="error" class="error-message">
            <p class="font-bold py-3">Error verifying your account.</p>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Verification',
    data: function () {
        return {
            error: "",
            verification_token: ""
        }
    },
    created() {
        if (this.$route.params.verification_token == null) {
            this.$router.push("/login")
        } else {
            this.verification_token = this.$route.params.verification_token
        }
    },
    methods: {
        attemptVerification: function () {
            axios({
                method: "PUT",
                url: "http://127.0.0.1:50009/api/v1/users",
                data: {
                    email: this.email,
                },
                headers: {
                    "content-type": "application/json",
                    "authorization": `Bearer ${this.token}`
                }
            })
                .then(() => {
                    localStorage.removeItem('token')
                    this.$router.push("/login")
                })
                .catch(error => {
                    console.log(error)
                    this.error = error.response.status
                })
        }
    }
}
</script>