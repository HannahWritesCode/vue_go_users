<template>
    <div class="min-w-full">
        <h1 class="text-4xl font-semibold">Update your account</h1>

        <form @submit.stop.prevent="attemptUpdate()">
            <div class="flex flex-col my-2">
                <label for="update-email">Email</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="email" id="update-email" v-model="email" />
            </div>

            <div class="flex flex-col my-2">
                <label for="first_name">First Name</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="string" v-model="first_name" />
            </div>

            <div class="flex flex-col my-2">
                <label for="last_name">Last Name</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="string" v-model="last_name" />
            </div>

            <div class="flex flex-col my-2">
                <label for="update-password">Password</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="password" id="update-password" v-model="password" />
            </div>

            <div class="flex flex-col my-2">
                <label for="update-password-confirmation">Confirm your password</label>
                <input
                    class="border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 p-0.5"
                    type="password" id="update-password-confirmation" v-model="password_confirmation" />
            </div>

            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Update my Account" id="submit-button" />
        </form>

        <div v-if="response == 200">
            <p class="font-bold py-3">Your account has been updated!</p>
        </div>

        <div v-else-if="response">
            <p class="font-bold py-3">Error updating your account. <br /> If you are updating your password, confirm it
                by
                typing it again in the "Confirm your password" field.</p>
        </div>

        <router-link to="/success" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Return to profile page</p>
        </router-link>

        <div>
            <button><a class="font-semibold py-3 text-green-500 hover:underline" @click="logout()">Log out</a></button>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'UpdateAccount',
    data: function () {
        return {
            email: "",
            first_name: "",
            last_name: "",
            password: "",
            password_confirmation: "",
            response: ""
        }
    },
    created() {
        let token = localStorage.getItem("token")
        if (token == null) {
            this.$router.push("/login")
        }
    },
    methods: {
        attemptUpdate: function () {
            let token = localStorage.getItem("token")
            axios({
                method: "PUT",
                url: "http://127.0.0.1:50009/api/v1/users",
                data: {
                    email: this.email,
                    first_name: this.first_name,
                    last_name: this.last_name,
                    password: this.password,
                    password_confirmation: this.password_confirmation
                },
                headers: {
                    "content-type": "application/json",
                    "authorization": `Bearer ${token}`
                }
            })
                .then(response => {
                    this.response = response.status
                })
                .catch(error => {
                    this.response = error.response.status
                })
        },
        logout: function () {
            localStorage.removeItem('token')
            this.$router.push("/login")
        }
    }
}
</script>