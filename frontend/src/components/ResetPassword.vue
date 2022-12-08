<template>
    <div>
        <h1 class="text-4xl">Reset Your Password</h1>

        <p class="py-3">Enter and confirm your new password.</p>

        <form @submit.stop.prevent="attemptResetPassword()">
            <div class="flex flex-col my-2">
                <label for="reset-password">New password</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="password" id="reset-password" v-model="password" required />
            </div>

            <div class="flex flex-col my-2">
                <label for="reset-password-confirmation">Confirm new password</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="password" id="reset-password-confirmation" v-model="password_confirmation" required />
            </div>

            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Submit" id="submit-button" />
        </form>

        <div v-if="error" class="error-message">
            <p>Error resetting your password.</p>
        </div>

        <div class="flex flex-row flex-nowrap">
            <p>Remember your old password?</p>
            <router-link to="/login" class="font-semibold text-green-500 hover:underline px-1">
                <p>Log in</p>
            </router-link>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
//import VueRouter from "vue-router"

//let token = localStorage.getItem("token")

export default {
    name: 'ResetPassword',
    data: function () {
        return {
            password: "",
            password_confirmation: "",
            reset_token: "",
            error: ""
        }
    },
    created() {
        if (this.$route.params.reset_token == null) {
            this.$router.push("/login")
        } else {
            this.reset_token = this.$route.params.reset_token
        }
    },
    methods: {
        attemptResetPassword: function () {
            axios({
                method: "POST",
                url: "http://127.0.0.1:50009/api/v1/reset-password",
                headers: {
                    "content-type": "application/json"
                },
                data: {
                    reset_token: this.reset_token,
                    password: this.password,
                    password_confirmation: this.password_confirmation
                }
            })
                .then(response => {
                    console.log(response)
                    this.$router.push("/success")
                })
                .catch(error => {
                    this.error = error.response.status
                })
        }
    }
}
</script>