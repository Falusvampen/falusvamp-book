<template>
  <div>
    <!-- <Sidebar /> -->
    <NavBar />
    <header>
      <h1>Welcome to FalusvampBook</h1>
      <h2>Group {{ $route.params.id }}</h2>
    </header>
    <main>
      <!-- <Post /> -->
      <!-- <div class="profile-tabs">
        <div
          class="profile-tab"
          :class="{ active: activeTab === 'allPosts' }"
          @click="activateTab('allPosts')"
        >
          All posts
        </div>
        <div
          class="profile-tab"
          :class="{ active: activeTab === 'group' }"
          @click="activateTab('group')"
        >
          Groups
        </div>
      </div>

      <div v-if="activeTab === 'allPosts'" class="profile-section">
        test
      </div> -->
      <div class="create">
        <CreatePost />
        <!-- <GroupsList /> -->
        <EventsList />

        <!-- <CreateGroup /> -->
        <CreateEvent />
      </div>
      <div class="feed">
        <div v-for="post in posts" :key="post.id">
          <Post
            :image="post.image"
            :title="post.title"
            :content="post.body"
            :user="post.user"
            :createdAt="post.createdAt"
            :comments="post.comments"
          />
        </div>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // Replace with data from server

      // A post should include the following data:
      // - Title
      // - Content
      // - User
      // - Created at
      // - Id (unique identifier)
      // - Image (optional)

      posts: [
        {
          id: 1,
          title: "Post 1",
          body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
          user: "John Doe",
          createdAt: "2023-07-06",
          // image: "path/to/image1.jpg",
          comments: [
            {
              id: 1,
              text: "Comment 1",
              author: "Jane Smith",
              date: "2023-07-07",
              // profileImage: "path/to/profileImage1.jpg",
              image: "path/to/commentImage1.jpg",
            },
            {
              id: 2,
              text: "Comment 2",
              author: "Jane Smith",
              date: "2023-07-07",
              // profileImage: "path/to/profileImage1.jpg",
              image: "path/to/commentImage1.jpg",
            },
            // Add more comments for the post
          ],
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
    };
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

header {
  /* background-color: #4267b2; */
  padding: 20px;
  /* color: white; */
  text-align: center;
  font-family: "Poppins", sans-serif;
}

.main {
  padding: 20px;
}

.feed {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  grid-gap: 20px;
}

.create {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.create > * {
  flex: 1;
  margin: 0 20px;
}
</style>
