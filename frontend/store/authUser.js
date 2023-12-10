import { defineStore } from 'pinia';

export const useAuthUser = defineStore('authUser', {
    state: () => ({
        isLoggedin: false,
        uuid: "",
    }),
    actions: {
        setAuthUser(isLoggedin, uuid) {
            this.isLoggedin = isLoggedin;
            this.uuid = uuid;
        },

        unsetAuthUser() {
            this.isLoggedin = false;
            this.uuid = "";
            document.cookie = "user_uuid=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/";
            navigateTo('/login');
        }
    },
    getters: {
        getAuthUser() {
            return {
                isLoggedin: this.isLoggedin,
                uuid: this.uuid,
            }
        }
    },
    persist: true,
})