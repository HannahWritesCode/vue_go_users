<template>
    <div class="min-w-full">
        <h1 class="text-4xl font-semibold">Reset Your Password</h1>

        <p class="py-3">Enter the email associated with your account, and we'll help you reset your password.</p>

        <form @submit.stop.prevent="attemptRequestPassword()">
            <div class="flex flex-col my-2">
                <label for="request-password-email">Email</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="email" id="request-password-email" v-model="email" required />
            </div>

            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Email me a reset link" id="submit-button" />
        </form>

        <div v-if="response == 200">
            <p class="font-bold py-3">Check the email you submitted for a link to reset your password.</p>
        </div>

        <div v-else-if="response">
            <p class="font-bold py-3">Error requesting password reset.</p>
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
    name: 'RequestPassword',
    data: function () {
        return {
            email: "",
            response: ""
        }
    },
    methods: {
        attemptRequestPassword: function () {
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
                .then(response => {
                    this.response = response.status;
                })
                .catch(error => {
                    this.response = error.response.status
                })
        }
    }
}
</script>