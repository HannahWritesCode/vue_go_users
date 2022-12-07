<template>
    <div>
        <h1 class="text-4xl">Reset Your Password</h1>

        <p class="py-3">Enter the email associated with your account, and we'll help you reset your password.</p>

        <form @submit.stop.prevent="attemptResetPassword()">
            <div class="flex flex-col my-2">
                <label for="reset-password-email">Email</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="email" id="reset-password-email" v-model="email" required />
            </div>

            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Submit" id="submit-button" />
        </form>

        <div v-if="email">
            <p>Check the email you submitted for a link to reset your password.</p>
        </div>

        <div v-else-if="error" class="error-message">
            <p>Error requesting password reset.</p>
        </div>

        <div class="flex flex-row flex-nowrap">
            <p>Remember your password?</p>
            <router-link to="/login" class="font-semibold text-green-500 hover:underline px-1">
                <p>Log in</p>
            </router-link>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'ResetPassword',
    data: function () {
        return {
            email: "",
            error: ""
        }
    },
    methods: {
        attemptResetPassword: function () {
            axios({
                method: "POST",
                url: "http://127.0.0.1:50009/api/v1/request-password",
                headers: {
                    "content-type": "application/json"
                },
                data: {
                    email: this.email
                }
            })
                .then(result => {
                    console.log(result)
                })
                .catch(error => {
                    this.error = error.response.status
                })
        }
    }
}
</script>