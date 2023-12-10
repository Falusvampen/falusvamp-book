<script setup>
  import { useAuthUser } from '/store/authUser';
  const authUser = useAuthUser();
// Equivalent of Vue 2 data 
const showNotifications = ref(false),
      showSidebar = ref(false),
      friendsList = reactive([]); // friendsList is an Array

async function logout() {
  await authUser.unsetAuthUser();  // Adjust this based on your project's auth plugin
}

function toggleNotifications() {
  showNotifications.value = !showNotifications.value;
  showSidebar.value = false; // Close sidebar when opening notifications
}

function toggleSidebar() {
  showNotifications.value = false; // Close notifications when opening sidebar
  showSidebar.value = !showSidebar.value;
}

function closeSidebar() {
  showSidebar.value = false;
}

// Equivalent of fetchFriendsList method
function fetchFriendsList() {
  // Logic for fetching the friends list
  friendsList.value = [
    { id: 1, name: "Friend 1" },
    { id: 2, name: "Friend 2" },
    { id: 3, name: "Friend 3" },
  ];
}

onMounted(fetchFriendsList)

</script>


<template>
  <div style="padding-bottom: 100px">
    <div class="navbar fixed-top">
      <div class="navbar">
        <div class="navbar-left">
          <router-link to="/" class="navbar-logo"> FalusvampBook </router-link>
        </div>
        <div class="navbar-right">
          <div class="navbar-icon" @click="toggleNotifications">
            <!-- <i class="far fa-bell"></i> -->
            <!-- Notification bell icon -->
            <iconsBell />
            <div class="dropdown-menu" v-if="showNotifications">
              <!-- Notification dropdown menu content -->
              <ul>
                <li>Notification 1</li>
                <li>Notification 2</li>
                <li>Notification 3</li>
              </ul>
            </div>
          </div>
          <div class="navbar-icon" @click="toggleSidebar">
            <iconsMessage />
          </div>
          <div class="navbar-icon" @click="logout">
          
              <!-- <i class="fas fa-sign-out-alt"></i> -->
              <!-- Logout icon -->
              <!-- <navigateTo path="/login" /> -->

              <iconsLogout />
            
          </div>
          <div class="navbar-icon">
            <!-- <i class="far fa-user"></i> -->
            <!-- Profile icon -->
            <profileImage />
          </div>
        </div>
      </div>
    </div>
    <div class="sidebar" v-if="showSidebar" @click="closeSidebar">
      <ul>
        <li v-for="friend in friendsList" :key="friend.id">
          {{ friend.name }}
        </li>
      </ul>
    </div>
  </div>
</template>


<!-- <script>
export default {
  data() {
    return {
      showNotifications: false,
      showSidebar: false,
      friendsList: [],
    };
  },
  methods: {
    async logout() {
      await authUser.unsetAuthUser();
    },
    toggleNotifications() {
      this.showNotifications = !this.showNotifications;
      this.showSidebar = false; // Close sidebar when opening notifications
    },
    toggleSidebar() {
      this.showNotifications = false; // Close notifications when opening sidebar
      this.showSidebar = !this.showSidebar;
    },
    closeSidebar() {
      this.showSidebar = false;
    },
    fetchFriendsList() {
      // Fetch friends list from JSON source (replace with your own API call)
      // Example code using axios:
      // axios.get('/api/friends').then(response => {
      //   this.friendsList = response.data;
      // }).catch(error => {
      //   console.error(error);
      // });

      // Simulated friends list data
      this.friendsList = [
        { id: 1, name: "Friend 1" },
        { id: 2, name: "Friend 2" },
        { id: 3, name: "Friend 3" },
      ];
    },
  },
  mounted() {
    this.fetchFriendsList();
  },
};
</script> -->

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

.navbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 20px;
  background-color: #4267b2;
  color: #fff;
  font-family: "Poppins", sans-serif;
}

.navbar-left {
  display: flex;
  align-items: center;
}

.navbar-logo {
  font-size: 24px;
  font-weight: bold;
  color: #fff;
  text-decoration: none;
}

.navbar-right {
  display: flex;
  align-items: center;
}

.navbar-icon {
  margin-left: 20px;
  font-size: 20px;
  cursor: pointer;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background-color: #a32222;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  min-width: 200px;
  padding: 10px;
  z-index: 100;
}

.dropdown-menu ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
}

.dropdown-menu li {
  padding: 10px;
  cursor: pointer;
}

.dropdown-menu li:hover {
  background-color: #f0f0f0;
}

.navbar.fixed-top {
  position: fixed;
  display: inline-block;
  top: 0;
  left: 0;
  right: 0;
  z-index: 9999;
}

.sidebar {
  position: fixed;
  top: 93px;
  right: 0;
  width: 200px;
  height: 100%;
  background-color: #ffffff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  z-index: 100;
  overflow-y: auto;
}

.sidebar ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
}

.sidebar li {
  padding: 10px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.sidebar li:hover {
  background-color: #f0f0f0;
}
</style>
