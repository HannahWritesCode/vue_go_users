<template>
    <div>
        <h1 class="text-4xl">Sign Up</h1>

        <form @submit.stop.prevent="attemptSignup()">
            <div class="flex flex-col my-2">
                <label for="signup-email">Email</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="email" id="sign-up" v-model="email" required />
            </div>

            <div class="flex flex-col my-2">
                <label for="signup-first-name">First Name</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="string" id="signup-first-name" v-model="first_name" required />
            </div>

            <div class="flex flex-col my-2">
                <label for="signup-last_name">Last Name</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="string" id="signup-first-name" v-model="last_name" required />
            </div>

            <div class="flex flex-col my-2">
                <label for="signup-password">Password</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="password" id="signup-password" v-model="password" required />
            </div>

            <div class="flex flex-col my-2">
                <label for="signup-password-confirmation">Confirm your password</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="password" id="signup-password-confirmation" v-model="password_confirmation" required />
            </div>

            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Create Account" id="submit-button" />
        </form>

        <div v-if="error" class="error-message">
            <p>Error creating your account.</p>
        </div>

        <div class="flex flex-row flex-wrap">
            <p class="pr-1">Already have an account?</p>
            <router-link to="/login" class="font-semibold text-green-500 hover:underline">
                <p>Log in</p>
            </router-link>
        </div>

    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Signup',
    data: function () {
        return {
            email: "",
            first_name: "",
            last_name: "",
            password: "",
            password_confirmation: "",
            error: ""
        }
    },
    methods: {
        attemptSignup: function () {
            axios({
                method: "POST",
                url: "http://127.0.0.1:50009/api/v1/users",
                data: {
                    email: this.email,
                    first_name: this.first_name,
                    last_name: this.last_name,
                    password: this.password,
                    password_confirmation: this.password_confirmation
                },
                headers: {
                    "content-type": "application/json"
                }
            })
                .then(response => {
                    console.log(response.data.token)
                    localStorage.setItem("token", response.data.token)
                    this.$router.push("/success")

                })
                .catch(error => {
                    this.error = error.response.status
                })
        }
    }
}
</script>