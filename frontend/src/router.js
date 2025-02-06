import { createWebHistory, createRouter } from "vue-router";
import Login from './components/Login.vue'
import Signup from './components/Signup.vue'
import UpdateAccount from './components/UpdateAccount.vue'
import ResetPassword from './components/ResetPassword.vue'
import RequestPassword from './components/RequestPassword.vue'
import Success from './components/Success.vue'
import DeleteConfirmation from './components/DeleteConfirmation.vue'

const routes = [
    {
        path: "/",
        component: Login,
    },
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/sign-up",
        component: Signup,
    },
    {
        path: "/success",
        component: Success,
    },
    {
        path: "/update-account",
        component: UpdateAccount,
    },
    {
        path: "/request-password",
        component: RequestPassword,
    },
    {
        path: "/reset-password/:reset_token",
        component: ResetPassword
    },
    {
        path: "/delete-confirmation",
        component: DeleteConfirmation
    }
    // {
    //     path: "/verification/:verification_token",
    //     component: Verification
    // }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
