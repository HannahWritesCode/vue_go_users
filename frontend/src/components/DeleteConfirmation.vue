<template>
    <div class="min-w-full">
        <h1 class="text-4xl font-semibold pb-3">You sure?</h1>

        <form @submit.stop.prevent="attemptDelete()">
            <input class="bg-green-500 hover:bg-green-700 text-white font-bold my-2 py-2 px-4 rounded" type="submit"
                value="Yes I'm sure, delete my account" id="submit-button" />
        </form>

        <div v-if="error" class="error-message">
            <p class="font-bold py-3">Error deleting your account.</p>
        </div>

        <router-link to="/success" class="font-semibold py-3 text-green-500 hover:underline">
            <p>Never mind, go back</p>
        </router-link>

    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'DeleteConfirmation',
    data: function () {
        return {
            error: "",
            token: ""
        }
    },
    created() {
        let token = localStorage.getItem("token")
        if (token == null) {
            this.$router.push("/login")
        }
        else {
            this.token = token
        }
    },
    methods: {
        attemptDelete: function () {
            axios({
                method: "DELETE",
                url: "http://127.0.0.1:50009/api/v1/users",
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