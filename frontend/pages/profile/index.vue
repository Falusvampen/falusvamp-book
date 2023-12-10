<template>
  <div>
    <NavBar />
    <div class="profile">
      <div class="profile-header">
        <img
          :src="user.image"
          class="profile-image"
          style="width: 180px; border-radius: 25px"
        />
        <h1>{{ user.username }}</h1>
        <p>{{ user.email }}</p>
      </div>

      <div class="profile-tabs">
        <div
          class="profile-tab"
          :class="{ active: activeTab === 'userInfo' }"
          @click="activateTab('userInfo')"
        >
          User Information
        </div>
        <div
          class="profile-tab"
          :class="{ active: activeTab === 'userActivity' }"
          @click="activateTab('userActivity')"
        >
          User Activity
        </div>
        <div
          class="profile-tab"
          :class="{ active: activeTab === 'followers' }"
          @click="activateTab('followers')"
        >
          Followers
        </div>
        <div
          class="profile-tab"
          :class="{ active: activeTab === 'following' }"
          @click="activateTab('following')"
        >
          Following
        </div>
      </div>

      <div v-if="activeTab === 'userInfo'" class="profile-section">
        <h2>User Information</h2>
        <div class="user-info">
          <!-- <div>
            <label>First Name:</label>
            <p>{{ user.firstName }}</p>
          </div>
          <div>
            <label>Last Name:</label>
            <p>{{ user.lastName }}</p>
          </div>
          <div>
            <label>Username:</label>
            <p>{{ user.username }}</p>
          </div>
          <div>
            <label>Email:</label>
            <p>{{ user.email }}</p>
          </div>
          <div>
            <label>Date of Birth:</label>
            <p>{{ user.dateOfBirth }}</p>
          </div> -->
          <div>
            <p>
              My name is {{ user.firstName }} {{ user.lastName }} and my date of
              birth is {{ user.dateOfBirth }}
            </p>
          </div>
          <div>
            <h3>About:</h3>
            <p>{{ user.about }}</p>
          </div>
        </div>
      </div>

      <div v-if="activeTab === 'userActivity'" class="profile-section">
        <h2>User Activity</h2>
        <div v-for="post in posts" :key="post.id">
          <Post
            :image="post.image"
            :title="post.title"
            :content="post.body"
            :user="post.user"
            :createdAt="post.createdAt"
          />
        </div>
      </div>

      <div v-if="activeTab === 'followers'" class="profile-sections"></div>

      <!-- <div v-else class="profile-section">
        <h2>{{ activeTab }}</h2>
        Display content for other tabs here
      </div> -->

      <div class="profile-section">
        <h2>Profile Privacy</h2>
        <p>Profile Visibility: {{ profileVisibility }}</p>
        <button @click="toggleProfileVisibility">{{ toggleButtonText }}</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      posts: [
        {
          id: 1,
          title: "Post 1",
          body: "This is the content of post 1.This is the content of post 1.This is the content of post 1.This is the content of post 1.This is the content of post 1.This is the content of post 1.This is the content of post 1.",
          user: "John Doe",
          createdAt: "2021-07-03",
          image:
            "https://t1.gstatic.com/licensed-image?q=tbn:ANd9GcRRv9ICxXjK-LVFv-lKRId6gB45BFoNCLsZ4dk7bZpYGblPLPG-9aYss0Z0wt2PmWDb",
        },
        {
          id: 2,
          title: "Post 2",
          body: "This is the content of post 2.",
          user: "John Doe",
          createdAt: "2021-07-03",
        },
        {
          id: 3,
          title: "Post 3",
          body: "This is the content of post 3.",
          user: "John Doe",
          createdAt: "2021-07-03",
        },
        {
          id: 4,
          title: "Post 4",
          body: "This is the content of post 4.",
          user: "John Doe",
          createdAt: "2021-07-03",
        },
        {
          id: 5,
          title: "Post 5",
          body: "This is the content of post 5.",
          user: "John Doe",
          createdAt: "2021-07-03",
        },
        {
          id: 6,
          title: "Post 5",
          body: "This is the content of post 5.",
          user: "John Doe",
          createdAt: "2021-07-03",
        },
        {
          id: 7,
          title: "Post 5",
          body: "This is the content of post 5.",
          user: "John Doe",
          createdAt: "2021-07-03",
        },
      ],
      user: {
        image: "https://i.pravatar.cc/300",
        name: "John Doe",
        firstName: "John",
        lastName: "Doe",
        username: "johndoe",
        email: "johndoe@example.com",
        dateOfBirth: "1990-01-01",
        about:
          "I am a web developer passionate about creating amazing user experiences.",
        // Other user information
      },
      profileVisibility: "public",
      activeTab: "userInfo",
    };
  },
  computed: {
    toggleButtonText() {
      return this.profileVisibility === "public"
        ? "Make Private"
        : "Make Public";
    },
  },
  methods: {
    toggleProfileVisibility() {
      this.profileVisibility =
        this.profileVisibility === "public" ? "private" : "public";
    },
    activateTab(tab) {
      this.activeTab = tab;
    },
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

.profile {
  font-family: "Poppins", sans-serif;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f1f1f1;
  border-radius: 8px;
}

.profile-header {
  text-align: center;
  margin-bottom: 20px;
}

.profile-tabs {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.profile-tab {
  padding: 10px 20px;
  margin-right: 10px;
  background-color: #fff;
  border: 1px solid #ccc;
  border-bottom: none;
  border-radius: 8px 8px 0 0;
  cursor: pointer;
}

.profile-tab.active {
  background-color: #007bff;
  color: #fff;
  border-color: #007bff;
}

.profile-section {
  margin-bottom: 30px;
}

.profile-section h2 {
  margin-bottom: 10px;
}

button {
  padding: 10px 20px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

button:focus {
  outline: none;
}
</style>
